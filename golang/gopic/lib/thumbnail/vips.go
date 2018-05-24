// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build vips

package thumbnail

import (
	"github.com/h2non/bimg"
)

func Create(dst, src string, maxWidth int) error {
	buf, err := bimg.Read(src)
	if err != nil {
		return err
	}

	out, err := bimg.NewImage(buf).Resize(maxWidth, 0)
	if err != nil {
		return err
	}

	createDstDir(dst)
	bimg.Write(dst, out)

	return nil
}
