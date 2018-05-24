package disk

import (
	"os"
	"fmt"
	"errors"
	"syscall"
	"path/filepath"
)

type Geometry struct {
	Sectors uint64
	SectorSize uint64 // Sector size in bytes
	Heads uint64
	Cylinders uint16
	Start uint64
}

type Size struct {
	Bytes uint64 // Disk size in bytes
	Blksize uint64
	Blocks uint64
}

type Disk struct {
	Path string
	Vendor string
	Model string
	Serial string
	Id uint64
	Major uint
	Minor uint
	Size *Size
	Geometry *Geometry
	file *os.File
}

var errNoDisk = errors.New("Not a disk")

// Find mountpoint of file or folder in the filesystem
// http://stackoverflow.com/questions/4453602/how-to-find-the-mountpoint-a-file-resides-on/34443698#34443698
func Mountpoint(path string) string {
	pi, err := os.Stat(path)
	if err != nil {
		return ""
	}

	odev := pi.Sys().(*syscall.Stat_t).Dev

	for path != "/" {
		_path := filepath.Dir(path)

		in, err := os.Stat(_path)
		if err != nil {
			return ""
		}

		if odev != in.Sys().(*syscall.Stat_t).Dev {
			break
		}

		path = _path
	}

	return path
}

func Scan() {
	scan()
}

func isDiskAvailable(name string) (*os.File,error) {
	f, err := os.OpenFile(name, os.O_RDONLY, 0600)
	if err != nil {
		return nil,err
	}

	fi, _ := f.Stat()

	if fi.Mode() & os.ModeDevice == 0 {
		return nil,errNoDisk
	}

	return f,nil
}

func NewDisk(path string) (*Disk,error) {
	f,err := isDiskAvailable(path)
	if err != nil {
		return nil,err
	}

	return &Disk{
		Path: path,
		Size: &Size{},
		Geometry: &Geometry{},
		file: f,
	},nil
}

func (d *Disk) Read() error {
	err := read(d)
	if err != nil {
		return err
	}

	stat := syscall.Stat_t{}
	err = syscall.Stat(d.Path, &stat)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", stat)
	//d.Size.Bytes   = uint64(stat.Size)
	d.Size.Blksize = uint64(stat.Blksize)
	d.Size.Blocks  = uint64(stat.Blocks)
	d.Id = uint64(stat.Dev)
	d.Major = uint(stat.Rdev/256)
	d.Minor = uint(stat.Rdev%256)

	return nil
}

func (d *Disk) Close() {
	d.file.Close()
}
