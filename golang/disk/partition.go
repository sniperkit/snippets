package disk

import (
	"bytes"
	"io"
	"strings"
	"io/ioutil"
	"bufio"
	"fmt"
)

func readFile(file string, handler func(string) bool) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if !handler(string(line)) {
			break
		}
	}

	return nil
}

// http://linux.die.net/man/3/getmntent
func ScanMounts() error {
	err := readFile("/etc/mtab", func(line string) bool {
		fields := strings.Fields(line)
		fmt.Println(fields)
/*
		fs := FileSystem{}
		fs.DevName = fields[0]
		fs.DirName = fields[1]
		fs.SysTypeName = fields[2]
		fs.Options = fields[3]

		fslist = append(fslist, fs)
*/

		return true
	})

	return err
}
