package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"compress/gzip"
)

func WriteGzFile(filename string, data []byte) error {
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz := gzip.NewWriter(fi)
	defer fz.Close()

	fz.Write(data)

	return nil
}

func ReadGzFile(filename string) ([]byte, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	s, err := ioutil.ReadAll(fz)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func main() {
	err := WriteGzFile("hello-world.gz", []byte("Hello World!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := ReadGzFile("hello-world.gz")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(c))
}
