package disk

import (
	"os"
	"log"
	"unsafe"
	"syscall"
)

// https://www.freebsd.org/doc/en_US.ISO8859-1/books/handbook/disk-organization.html#disks-naming
func scan() {

}

// https://github.com/freebsd/freebsd/blob/master/sys/sys/disk.h
// https://github.com/freebsd/freebsd/blob/master/sys/sys/disklabel.h
// http://git.cgsecurity.org/cgit/testdisk/tree/src/hdaccess.c#n639


const (
        DIOCGSECTORSIZE = 0x40046480
	 /* Get the sector size of the device in bytes.  The sector size is the
	 * smallest unit of data which can be transferred from this device.
	 * Usually this is a power of 2 but it might not be (i.e. CDROM audio).
	 */
        DIOCGMEDIASIZE  = 0x40086481 /* Get media size in bytes */
	/*
	 * Get the size of the entire device in bytes.  This should be a
	 * multiple of the sector size.
	 */
        DIOCGFWSECTORS  = 0x40046482
        DIOCGFWHEADS    = 0x40046483
)

func diocgmediasize(f *os.File) (uint64, error) {
	var size uint64
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), DIOCGMEDIASIZE, uintptr(unsafe.Pointer(&size)))
	if e != 0 {
		log.Fatal(e)
		return 0, e
	}
	return size,nil
}

func diocgsectorsize(f *os.File) (uint64, error) {
	var size uint64
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), DIOCGSECTORSIZE, uintptr(unsafe.Pointer(&size)))
	if e != 0 {
		log.Fatal(e)
		return 0, e
	}
	return size,nil
}

func diocgfwsectors(f *os.File) (uint64, error) {
	var size uint64
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), DIOCGFWSECTORS, uintptr(unsafe.Pointer(&size)))
	if e != 0 {
		log.Fatal(e)
		return 0, e
	}
	return size,nil
}

func diocgfwheads(f *os.File) (uint64, error) {
	var size uint64
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), DIOCGFWHEADS, uintptr(unsafe.Pointer(&size)))
	if e != 0 {
		log.Fatal(e)
		return 0, e
	}
	return size,nil
}

func read(d *Disk) error {
	d.Size.Bytes,_ = diocgmediasize(d.file)
	d.Geometry.Sectors, _ = diocgfwsectors(d.file)
	d.Geometry.Heads, _ = diocgfwheads(d.file)
	d.Geometry.SectorSize, _ = diocgsectorsize(d.file)

	return nil
}
