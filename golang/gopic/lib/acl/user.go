// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package acl

import (
	"time"
)

type User struct {
	Email string `storm:"id,unique"`
	Name string
	Password []byte
	Groups []Group
	Active bool
	RegisteredAt time.Time
}
