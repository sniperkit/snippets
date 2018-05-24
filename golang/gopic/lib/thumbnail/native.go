// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build !vips

package thumbnail

import (
	"image/jpeg"
	"github.com/nfnt/resize"
)

// Create generates a thumbnail from src filename and creates the directories to dst and writes
func Create(dst, src string, maxWidth int) error {
	in, err := createSrc(src)
	if err != nil {
		return nil
	}

	img, err := jpeg.Decode(in)
	if err != nil {
		in.Close()
		return err
	}
	in.Close()

	// Resize to ratio with a max width of 1000 pixels
	m := resize.Resize(uint(maxWidth), 0, img, resize.Lanczos3)

	out, err := createDst(dst)
	if err != nil {
		return err
	}

	jpeg.Encode(out, m, nil)
	out.Close()

	return nil
}
