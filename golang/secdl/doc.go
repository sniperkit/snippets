// Copyright 2015 Jerry Jacobs. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package secdl implements the famous Lighttpd SecDownload secure download URL
//  algorithm
//
// Overview
//
// A expirable secure download URL is created as follows (without error checking):
//
//  s := secdl.New()
//
//  s.SetSecret("my-secret-token")
//  s.SetBaseURL("https://example.com")
//  s.SetPrefix("/dl/")
//  s.SetFilename("/dir1/dir2/dir3/my-secret-file.txt")
//
//  url, _ := s.Encode(ExpireHour)
//  // Output: https://example.com/dl/3990f415f215d5d7f20307ddf8a4b387/572ba80c/dir1/dir2/dir3/my-secret-file.txt
//
package secdl
