// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pic

// IDSize is the maximum hex encoded string of the picture hash ID
const IDSize = 64

// IDMinSize is the minimum hex encoded string of the picture hash ID
const IDMinSize = 11

type ID [IDSize]byte

func (id ID) String() string {
	return string(id[:])
}
