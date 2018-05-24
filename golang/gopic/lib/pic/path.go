// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pic

// Path holds a list of File ids
type Path struct {
	Hash     string `storm:"id"`
	Path     string `storm:"index"`
	HashList []string
}
