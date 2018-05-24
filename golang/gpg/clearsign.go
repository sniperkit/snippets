package main

import (
        "fmt"
	"encoding/hex"
	"strconv"
	"strings"
//        "os"
	"bytes"
        "golang.org/x/crypto/openpgp"
        //"golang.org/x/crypto/openpgp/armor"
        "golang.org/x/crypto/openpgp/clearsign"
)

func main() {
        var e *openpgp.Entity
        e, err := openpgp.NewEntity("Foo Bar", "", "foo@bar.com", nil)
        if err != nil {
                fmt.Println(err)
                return
        }

        // Add more identities here if you wish
        // Sign all the identities
/*
        for _, id := range e.Identities {
                err := id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, nil)
                if err != nil {
                        fmt.Println(err)
                        return
                }
        }

        w, err := armor.Encode(os.Stdout, openpgp.PublicKeyType, nil)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer w.Close()
*/

        //e.Serialize(w)

	var buf bytes.Buffer

	plaintext, err := clearsign.Encode(&buf, e.PrivateKey, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello signed world!"

	if _, err = plaintext.Write([]byte(msg)); err != nil {
		fmt.Printf("error from Write: %s", err)
		return
	}
	if err = plaintext.Close(); err != nil {
		fmt.Printf("error from Close: %s", err)
		return
	}

	fmt.Printf("---\nName: %+v\n", e.Identities)
	keyid := strings.ToUpper(strconv.FormatUint(e.PrimaryKey.KeyId, 16))
	fp := strings.ToUpper(hex.EncodeToString(e.PrimaryKey.Fingerprint[:]))
	fmt.Printf("Fingerprint: %s\n", fp)
	fmt.Printf("Long KeyId: %s\n", keyid)
	keyidShort := keyid[len(keyid)-8:]
	fmt.Printf("Short KeyId: %s\n", keyidShort)
	fmt.Printf("---")

	fmt.Printf("%s", buf.String())
}
