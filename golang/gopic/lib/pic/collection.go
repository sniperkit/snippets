// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pic

// Collection holds a list of File ids which are Path independed
type Collection struct {
	ID       string `storm:"id,unique"`
	HashList []string
}
