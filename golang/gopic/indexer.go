// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/xor-gate/gopic/lib/pic"
	"encoding/hex"
	"golang.org/x/crypto/blake2s"
	"os"
	"path/filepath"
)

// Indexer walks recursive through searchPath and adds jpeg pictures to the database
func Indexer(db *storm.DB, searchPath string, stats *IndexerStat) {
	// Make sure we scan valid folders and nothing else
	if finfo, err := os.Stat(searchPath); err != nil || !finfo.IsDir() {
		fmt.Println("warn: directory \"", searchPath, "\" is not valid")
		return
	}

	_ = filepath.Walk(searchPath, func(path string, f os.FileInfo, err error) error {
		// Skip directory entries
		if f.IsDir() {
			return nil
		}

		// Check if PicFile is already known by filepath
		pfile := &pic.File{}
		err = db.One("Filepath", path, pfile)
		if err == nil {
			stats.AddFile(f.Size())
			return nil
		}

		// Create new PicFile
		pfile = pic.NewFile(path)
		if pfile == nil {
			return nil
		}
		defer pfile.Close()

		// Check if PicFile Hash already exists
		err = db.One("Hash", pfile.GetHash(), &pic.File{})
		if err == nil {
			stats.AddDuplicateFile(f.Size())
			return nil
		}

		pfile.LoadExif()

		// Store the PicFile
		err = db.Save(pfile)
		if err != nil {
			fmt.Println("Error on save",err,"on", pfile.Filename,"(hash:",pfile.GetHash(),")")
			return nil
		}

		stats.AddNewFile(f.Size())

		// PicFilePath
		basePath := filepath.Dir(pfile.Filepath)
		pfilePath := &pic.Path{}
		err = db.One("Path", basePath, pfilePath)
		if err != nil {
			pfilePath.Path = basePath
			hash := blake2s.Sum256([]byte(pfilePath.Path))
			pfilePath.Hash = hex.EncodeToString(hash[:])
			pfilePath.HashList = append(pfilePath.HashList, pfile.GetHash())
			err = db.Save(pfilePath)
		} else {
			pfilePath.HashList = append(pfilePath.HashList, pfile.GetHash())
			err = db.Save(pfilePath)
		}

		return nil
	})
}
