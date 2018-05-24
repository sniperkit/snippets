// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

// IndexerStat keeps statistics of the PicFile index
type IndexerStat struct {
	timeStarted time.Time
	timeFinished time.Time
	duplicateFiles int64
	duplicateSize int64
	indexedFiles int64
	indexedSize int64
	indexedNewFiles int64
	indexedNewSize  int64
}

// String writes a summary of the stats
func (s IndexerStat) String() string {
	summary := `
Summary:
	 Started at: %v
	Finished at: %v
	       Took: %v

	  New unique files: %v
	Total unique files: %v
	   Duplicate files: %v

	       Total files: %v`

	totalUniqueFiles := s.indexedNewFiles + s.indexedFiles
	totalFiles       := totalUniqueFiles + s.duplicateFiles

	return fmt.Sprintf(summary, s.timeStarted, s.timeFinished, time.Since(s.timeStarted), s.indexedNewFiles, totalUniqueFiles, s.duplicateFiles, totalFiles)
}

// Start the indexer time measurement
func (s *IndexerStat) Start() {
	s.timeStarted = time.Now()
}

// Finish the indexer time measurement
func (s *IndexerStat) Finish() {
	s.timeFinished = time.Now()
}

// Reset all stats
func (s *IndexerStat) Reset() {
	s.duplicateFiles = 0
	s.duplicateSize = 0
	s.indexedFiles = 0
	s.indexedSize = 0
	s.indexedNewFiles = 0
	s.indexedNewSize = 0
}

// AddFile adds a single file and size (in bytes) to the current stats
func (s *IndexerStat) AddFile(size int64) {
	s.indexedFiles++
	s.indexedSize += size
}

// AddDuplicateFile adds a single file with given size (in bytes)
func (s *IndexerStat) AddDuplicateFile(size int64) {
	s.duplicateFiles++
	s.duplicateSize += size
}

// AddNewFile adds a single new file with given size (in bytes)
func (s *IndexerStat) AddNewFile(size int64) {
	s.indexedNewFiles++
	s.indexedNewSize += size
}
