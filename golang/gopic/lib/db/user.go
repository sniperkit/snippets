// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/xor-gate/gopic/lib/acl"
)

var ErrorUserIsInactive = fmt.Errorf("user is inactive")

func GetUsers() (*[]acl.User, error) {
	var users []acl.User
	if err := dbSession.All(&users); err != nil {
		return nil, err
	}
	return &users, nil
}

func UserGet(email string) (*acl.User, error) {
	var u acl.User
	if err := dbSession.One("Email", email, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func UserVerify(email, password string) error {
	var u acl.User
	if err := dbSession.One("Email", email, &u); err != nil {
		return err
	}

	if !u.Active {
		return ErrorUserIsInactive
	}

	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}

func UserSave(email, password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &acl.User{Email: email, Password: pass, RegisteredAt: time.Now()}
	return dbSession.Save(u)
}
