// https://gitlab.flux.utah.edu/a3/vmi/blob/13d0c81d8b0952f04b062e085edd2834c98b8135/vmprobes/examples/ioctlent.h
// https://github.com/bnoordhuis/strace/blob/master/ioctlsort.c
// http://sourceforge.net/p/strace/mailman/message/34335651/
// http://lxr.free-electrons.com/source/include/linux/dm-ioctl.h
// http://git.cgsecurity.org/cgit/testdisk/tree/src/hdaccess.c
// https://www.win.tue.nl/~aeb/linux/lk/lk-8.html
// http://www.win.tue.nl/~aeb/partitions/partition_types-1.html
// hdparm -I /dev/sda
// lshw -class disk
// smartctl -i /dev/sda
// libatasmart
// fdisk -l /dev/sda
// smartmontools
// http://lxr.free-electrons.com/source/include/uapi/linux/btrfs.h
// http://lxr.free-electrons.com/source/Documentation/ioctl/ioctl-number.txt
// http://www.linuxquestions.org/questions/linux-general-1/what-is-disk-identifier-740408/  < read disk id
// http://www.opensource.apple.com/source/diskdev_cmds/diskdev_cmds-557.3.1/fdisk.tproj/mbr.c
// http://www.opensource.apple.com/source/diskdev_cmds/diskdev_cmds-557.3.1/fdisk.tproj/part.c
// http://lxr.free-electrons.com/source/include/uapi/linux/magic.h
// https://www.kernel.org/doc/Documentation/ioctl/hdio.txt
// http://www.t13.org/documents/UploadedDocuments/docs2009/d2015r1a-ATAATAPI_Command_Set_-_2_ACS-2.pdf
// http://lxr.free-electrons.com/source/include/uapi/linux/hdreg.h
// http://www.makelinux.net/books/lkd2/ch12lev1sec9
// https://www.win.tue.nl/~aeb/linux/lk/lk-8.html
// TODO: disk label
// TODO: disk id
// TODO: MBR
// TODO: partition table
// TODO: partition types
// TODO: smart in seperate package
// TODO: vfsmount (instead of ugly /proc/mounts, /etc/mtab)
package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"io/ioutil"
	"unsafe"
	"syscall"
)

type disk struct {
	size uint64 // Disk size in bytes
}

/* converted from struct hd_geometry, add linux header ... for 64bit system! */
type hd_geometry struct {
	heads uint8
	sectors uint8
	cylinders uint16
	start uint64
	__padding uint16
}

const (
	HDIO_GETGEO  = 0x0301
	BLKGETSIZE64 = 0x80081272
)

func hdio_getgeo(f *os.File) (hd_geometry, error) {
	var geo hd_geometry
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), HDIO_GETGEO, uintptr(unsafe.Pointer(&geo)))
	if e != 0 {
		log.Fatal(e)
		return hd_geometry{}, e
	}
	return geo,nil
}

func blkgetsize64(f *os.File) (uint64, error) {
	var size uint64
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), BLKGETSIZE64, uintptr(unsafe.Pointer(&size)))
	if e != 0 {
		log.Fatal(e)
		return 0, e
	}
	return size,nil
}

func readMajorMinor(device string) {
	stat := syscall.Stat_t{}
	_ = syscall.Stat(device, &stat)
	fmt.Println("major:", uint64(stat.Rdev/256), "minor:", uint64(stat.Rdev%256))
}

func readVendorModel(blockdevice string) {
	var vendor string
	var model string

	buf, err := ioutil.ReadFile("/sys/block/" + blockdevice + "/device/vendor")
	if err == nil {
		vendor = strings.Trim(string(buf), " \n")
	}

	buf, err = ioutil.ReadFile("/sys/block/" + blockdevice + "/device/model")
	if err == nil {
		model = strings.Trim(string(buf), " \n")
	}

	fmt.Println("vendor:", vendor, "model:", model)

/*
 4223 open("/sys/block/sdf/device/vendor", O_RDONLY) = 4
 4224 fstat(4, {st_mode=S_IFREG|0444, st_size=4096, ...}) = 0
 4225 mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fba3733d000
 4226 read(4, "Generic \n", 4096)             = 9
 4227 close(4)                                = 0
 4228 munmap(0x7fba3733d000, 4096)            = 0
 4229 open("/sys/block/sdf/device/model", O_RDONLY) = 4
 4230 fstat(4, {st_mode=S_IFREG|0444, st_size=4096, ...}) = 0
 4231 mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fba3733d000
 4232 read(4, "USB Flash Disk  \n", 4096)     = 17
 4233 close(4)
*/
}

func main() {
	dev := "/dev/sda"

	f, err := os.OpenFile(dev, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fi, _ := f.Stat()

	if fi.Mode() & os.ModeDevice == 0 {
		log.Fatal("no device")
	}

	size, _ := blkgetsize64(f)
	geo,  _ := hdio_getgeo(f)

	fmt.Println("device node:", dev)
	readMajorMinor(dev)
	readVendorModel("sda")
	fmt.Println("sectors:", geo.sectors, "heads:", geo.heads, "cylinders:", geo.cylinders)
	fmt.Println("fd:", f.Fd())
	fmt.Println("size (bytes):", size)
}
