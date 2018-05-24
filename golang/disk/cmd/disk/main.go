package main

import (
	"os"
	"fmt"
	"log"
	"github.com/xor-gate/disk"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("disk: <path>")
	}

	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	disk.Scan()
	disk.ScanMounts()

	d, err := disk.NewDisk(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	err = d.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("d: %+v\n", d)
	fmt.Printf("d.Size: %+v\n", d.Size)
	fmt.Printf("d.Geometry: %+v\n", d.Geometry)
}
