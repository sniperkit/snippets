// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"github.com/xor-gate/gopic/lib/pic"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/asdine/storm/codec/json"
)

var dbSession *storm.DB

// Open creates a single database
func Open(dbFile string) error {
	if dbSession != nil {
		return nil
	}

	db, err := storm.Open(dbFile, storm.Codec(json.Codec))
	if err != nil {
		return err
	}

	dbSession = db

	return nil
}

// Get returns the instance to the database
func Get() *storm.DB {
	return dbSession
}

// GetPicFileByHash fetches a single File from a hash or shorthash
func GetPicFileByHash(hash string) (*pic.File, error) {
	// TODO: check hash for invalid non-hex chars (0-9,a-f,A-F)
	// TODO: Normalize hash to lowercase

	if len(hash) < pic.IDMinSize {
		return nil, fmt.Errorf("hash is to short")
	}

	if len(hash) == pic.IDSize {
		var f pic.File
		if err := dbSession.One("Hash", hash, &f); err != nil {
			return nil, err
		}
		return &f, nil
	}

	// Fetch single File object based on shorthash
	var pfiles []pic.File

	if err := dbSession.Select(q.Re("Hash", "^" + hash)).Find(&pfiles); err != nil {
		return nil, err
	}
	if len(pfiles) != 1 {
		return nil, fmt.Errorf("to many File results, shorthash collision")
	}
	return &pfiles[0],nil
}

// GetPicFilePath fetches a path by hash or shorthash
func GetPicFilePath(hash string) (*pic.Path, error) {
	if len(hash) == pic.IDSize {
		var f pic.Path
		if err := dbSession.One("Hash", hash, &f); err != nil {
			return nil, err
		}
		return &f, nil
	}

	// Fetch single FilePath object based on shorthash
	var pfiles []pic.Path

	if err := dbSession.Select(q.Re("Hash", "^" + hash)).Find(&pfiles); err != nil {
		return nil, err
	}
	if len(pfiles) != 1 {
		return nil, fmt.Errorf("to many FilePath results, shorthash collision")
	}
	return &pfiles[0],nil
}

// GetPicFilePaths fetches all indexed paths
func GetPicFilePaths() (*[]pic.Path, error) {
	var p []pic.Path

	if err := dbSession.All(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
