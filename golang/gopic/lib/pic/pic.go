// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pic

import (
	"encoding/hex"
	"github.com/xor-gate/goexif2/exif"
	"github.com/xor-gate/goexif2/mknote"
	"golang.org/x/crypto/blake2s"
	"strconv"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"image"
	"time"
	"path/filepath"
	"github.com/xor-gate/gopic/lib/thumbnail"
)

// File holds the (meta-)information of a single Picture
type File struct {
	ID        ID `storm:"id"`
	Filename  string
	Filepath  string `storm:"index"`
	Mime      string
	Width     int   // Width measured in pixels
	Height    int   // Height measured in pixels
	Time      time.Time
	f         *os.File
}

// NewFile opens and creates a File instance
// Close must be called!
func NewFile(filename string) *File {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}

	p := &File{}
	p.f = f
	p.Filepath = filename
	p.Filename = filepath.Base(p.Filepath)
	p.mimeLoad()

	// Only allow real pictures
	if p.Mime != "image/jpeg" {
		p.f.Close()
		return nil
	}

	return p
}

// Close closes the filedescriptor
func (p *File) Close() {
	p.f.Close()
}

// ThumbnailPath calculates the filepath to the thumbnail with maxWidth
func (p *File) ThumbnailPath(maxWidth int) string {
	// TODO: thumbnail path should be configurable
	return "thumbnails/"+strconv.FormatUint(uint64(maxWidth), 10)+"px/"+p.ID.String()+".jpg"
}

// Thumbnail generates a jpeg File thumbnail of maxWidth
func (p *File) Thumbnail(maxWidth int) error {
	if p.ID.String() == "" {
		p.hashLoad()
	}

	return thumbnail.Create(p.ThumbnailPath(maxWidth), p.Filepath, maxWidth)
}

// GetHash generates the blake2s hash id of the filecontent
func (p *File) GetHash() string {
	if p.ID.String() == "" {
		p.hashLoad()
	}
	return p.ID.String()
}

// LoadExif reads the exif metadata
func (p *File) LoadExif() error {
	exif.RegisterParsers(mknote.All...)

	p.f.Seek(0, 0)
	x, err := exif.Decode(p.f)
	if err != nil {
		return err
	}

	// Width + Height
	w, err := x.Get(exif.PixelXDimension)
	if err == nil {
		p.Width, _ = w.Int(0)
	}
	h, err := x.Get(exif.PixelYDimension)
	if err == nil {
		p.Height, _ = h.Int(0)
	}
	if p.Height == 0 || p.Width == 0 {
		p.f.Seek(0, 0)
		img, _, _ := image.DecodeConfig(p.f)

		p.Width = img.Width
		p.Height = img.Height
	}

	// DateTime, we allow error during datetime
	datetime, err := x.DateTime()
	if err == nil {
		p.Time = datetime
	}

	return nil
}

// mimeLoad first based on file extension or else using mime-magic of first 512bytes of the file
func (p *File) mimeLoad() error {
	p.Mime = mime.TypeByExtension(filepath.Ext(p.Filename))
	if p.Mime == "" {
		// Only the first 512 bytes are used to sniff the content type.
		buffer := make([]byte, 512)
		p.f.Seek(0, 0)
		_, err := p.f.Read(buffer)
		if err != nil {
			return err
		}

		p.Mime = http.DetectContentType(buffer)
	}
	return nil
}

// hashLoad reads the whole file memory and calculates the blake2s checksum (256bits) as hex encoded string
func (p *File) hashLoad() error {
	p.f.Seek(0, 0)
	bytes, err := ioutil.ReadAll(p.f)
	if err != nil {
		return err
	}

	hash := blake2s.Sum256(bytes)
	copy(p.ID[:], hex.EncodeToString(hash[:]))

	return nil
}
