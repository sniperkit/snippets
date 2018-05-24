package disk

import (
	"os"
	"fmt"
	"log"
	"path"
	"strings"
	"io/ioutil"
	"unsafe"
	"syscall"
)

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/hdreg.h
type hd_geometry struct {
	heads uint8
	sectors uint8
	cylinders uint16
	start uint64
	__padding uint16
}

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/hdreg.h
type hd_driveid struct {
	config		uint16		/* lots of obsolete bit flags */
	cyls		uint16		/* Obsolete, "physical" cyls */
	reserved2	uint16	/* reserved (word 2) */
	heads		uint16	/* Obsolete, "physical" heads */
	track_bytes	uint16	/* unformatted bytes per track */
	sector_bytes	uint16	/* unformatted bytes per sector */
	sectors		uint16	/* Obsolete, "physical" sectors per track */
	vendor0		uint16	/* vendor unique */
	vendor1		uint16	/* vendor unique */
	vendor2		uint16	/* Retired vendor unique */
	serial_no	[20]uint8	/* 0 = not_specified */
	buf_type	uint16	/* Retired */
	buf_size	uint16	/* Retired, 512 byte increments
					 * 0 = not_specified
					 */
	ecc_bytes	uint16	/* for r/w long cmds 0 = not_specified */
	fw_rev	[8]uint8	/* 0 = not_specified */
	model	[40]uint8	/* 0 = not_specified */
	max_multsect	uint8	/* 0=not_implemented */
	vendor3		uint8	/* vendor unique */
	dword_io	uint16	/* 0=not_implemented 1=implemented */
	vendor4		uint8	/* vendor unique */
	capability	uint8	/* (upper byte of word 49)
					 *  3:	IORDYsup
					 *  2:	IORDYsw
					 *  1:	LBA
					 *  0:	DMA
					 */
	reserved50	uint16	/* reserved (word 50) */
	vendor5		uint8	/* Obsolete, vendor unique */
	tPIO		uint8	/* Obsolete, 0=slow, 1=medium, 2=fast */
	vendor6		uint8	/* Obsolete, vendor unique */
	tDMA		uint8	/* Obsolete, 0=slow, 1=medium, 2=fast */
	field_valid	uint16	/* (word 53)
					 *  2:	ultra_ok	word  88
					 *  1:	eide_ok		words 64-70
					 *  0:	cur_ok		words 54-58
					 */
	cur_cyls	uint16	/* Obsolete, logical cylinders */
	cur_heads	uint16	/* Obsolete, l heads */
	cur_sectors	uint16	/* Obsolete, l sectors per track */
	cur_capacity0	uint16	/* Obsolete, l total sectors on drive */
	cur_capacity1	uint16	/* Obsolete, (2 words, misaligned int)     */
	multsect	uint8	/* current multiple sector count */
	multsect_valid	uint8	/* when (bit0==1) multsect is ok */
	lba_capacity	uint32	/* Obsolete, total number of sectors */
	dma_1word	uint16	/* Obsolete, single-word dma info */
	dma_mword	uint16	/* multiple-word dma info */
	eide_pio_modes	uint16	/* bits 0:mode3 1:mode4 */
	eide_dma_min	uint16	/* min mword dma cycle time (ns) */
	eide_dma_time	uint16	/* recommended mword dma cycle time (ns) */
	eide_pio	uint16	/* min cycle time (ns), no IORDY  */
	eide_pio_iordy	uint16	/* min cycle time (ns), with IORDY */
	words69_70	[2]uint16	/* reserved words 69-70
					 * future command overlap and queuing
					 */
	words71_74	[4]uint16	/* reserved words 71-74
					 * for IDENTIFY PACKET DEVICE command
					 */
	queue_depth	uint16	/* (word 75)
					 * 15:5	reserved
					 *  4:0	Maximum queue depth -1
					 */
	words76_79	[4]uint16	/* reserved words 76-79 */
	major_rev_num	uint16	/* (word 80) */
	minor_rev_num	uint16	/* (word 81) */
	command_set_1	uint16	/* (word 82) supported
					 * 15:	Obsolete
					 * 14:	NOP command
					 * 13:	READ_BUFFER
					 * 12:	WRITE_BUFFER
					 * 11:	Obsolete
					 * 10:	Host Protected Area
					 *  9:	DEVICE Reset
					 *  8:	SERVICE Interrupt
					 *  7:	Release Interrupt
					 *  6:	look-ahead
					 *  5:	write cache
					 *  4:	PACKET Command
					 *  3:	Power Management Feature Set
					 *  2:	Removable Feature Set
					 *  1:	Security Feature Set
					 *  0:	SMART Feature Set
					 */
	command_set_2	uint16	/* (word 83)
					 * 15:	Shall be ZERO
					 * 14:	Shall be ONE
					 * 13:	FLUSH CACHE EXT
					 * 12:	FLUSH CACHE
					 * 11:	Device Configuration Overlay
					 * 10:	48-bit Address Feature Set
					 *  9:	Automatic Acoustic Management
					 *  8:	SET MAX security
					 *  7:	reserved 1407DT PARTIES
					 *  6:	SetF sub-command Power-Up
					 *  5:	Power-Up in Standby Feature Set
					 *  4:	Removable Media Notification
					 *  3:	APM Feature Set
					 *  2:	CFA Feature Set
					 *  1:	READ/WRITE DMA QUEUED
					 *  0:	Download MicroCode
					 */
	cfsse		uint16	/* (word 84)
					 * cmd set-feature supported extensions
					 * 15:	Shall be ZERO
					 * 14:	Shall be ONE
					 * 13:6	reserved
					 *  5:	General Purpose Logging
					 *  4:	Streaming Feature Set
					 *  3:	Media Card Pass Through
					 *  2:	Media Serial Number Valid
					 *  1:	SMART selt-test supported
					 *  0:	SMART error logging
					 */
	cfs_enable_1	uint16	/* (word 85)
					 * command set-feature enabled
					 * 15:	Obsolete
					 * 14:	NOP command
					 * 13:	READ_BUFFER
					 * 12:	WRITE_BUFFER
					 * 11:	Obsolete
					 * 10:	Host Protected Area
					 *  9:	DEVICE Reset
					 *  8:	SERVICE Interrupt
					 *  7:	Release Interrupt
					 *  6:	look-ahead
					 *  5:	write cache
					 *  4:	PACKET Command
					 *  3:	Power Management Feature Set
					 *  2:	Removable Feature Set
					 *  1:	Security Feature Set
					 *  0:	SMART Feature Set
					 */
	cfs_enable_2	uint16	/* (word 86)
					 * command set-feature enabled
					 * 15:	Shall be ZERO
					 * 14:	Shall be ONE
					 * 13:	FLUSH CACHE EXT
					 * 12:	FLUSH CACHE
					 * 11:	Device Configuration Overlay
					 * 10:	48-bit Address Feature Set
					 *  9:	Automatic Acoustic Management
					 *  8:	SET MAX security
					 *  7:	reserved 1407DT PARTIES
					 *  6:	SetF sub-command Power-Up
					 *  5:	Power-Up in Standby Feature Set
					 *  4:	Removable Media Notification
					 *  3:	APM Feature Set
					 *  2:	CFA Feature Set
					 *  1:	READ/WRITE DMA QUEUED
					 *  0:	Download MicroCode
					 */
	csf_default	uint16	/* (word 87)
					 * command set-feature default
					 * 15:	Shall be ZERO
					 * 14:	Shall be ONE
					 * 13:6	reserved
					 *  5:	General Purpose Logging enabled
					 *  4:	Valid CONFIGURE STREAM executed
					 *  3:	Media Card Pass Through enabled
					 *  2:	Media Serial Number Valid
					 *  1:	SMART selt-test supported
					 *  0:	SMART error logging
					 */
	dma_ultra	uint16	/* (word 88) */
	trseuc		uint16	/* time required for security erase */
	trsEuc		uint16	/* time required for enhanced erase */
	CurAPMvalues	uint16	/* current APM values */
	mprc		uint16	/* master password revision code */
	hw_config	uint16	/* hardware config (word 93)
					 * 15:	Shall be ZERO
					 * 14:	Shall be ONE
					 * 13:
					 * 12:
					 * 11:
					 * 10:
					 *  9:
					 *  8:
					 *  7:
					 *  6:
					 *  5:
					 *  4:
					 *  3:
					 *  2:
					 *  1:
					 *  0:	Shall be ONE
					 */
	acoustic	uint16	/* (word 94)
					 * 15:8	Vendor's recommended value
					 *  7:0	current value
					 */
	msrqs		uint16	/* min stream request size */
	sxfert		uint16	/* stream transfer time */
	sal		uint16	/* stream access latency */
	spg		uint32	/* stream performance granularity */
	lba_capacity_2	uint64	/* 48-bit total number of sectors */
	words104_125	[22]uint16	/* reserved words 104-125 */
	last_lun	uint16	/* (word 126) */
	word127	uint16	/* (word 127) Feature Set
					 * Removable Media Notification
					 * 15:2	reserved
					 *  1:0	00 = not supported
					 *	01 = supported
					 *	10 = reserved
					 *	11 = reserved
					 */
	dlf		uint16	/* (word 128)
					 * device lock function
					 * 15:9	reserved
					 *  8	security level 1:max 0:high
					 *  7:6	reserved
					 *  5	enhanced erase
					 *  4	expire
					 *  3	frozen
					 *  2	locked
					 *  1	en/disabled
					 *  0	capability
					 */
	csfo		uint16	/*  (word 129)
					 * current set features options
					 * 15:4	reserved
					 *  3:	auto reassign
					 *  2:	reverting
					 *  1:	read-look-ahead
					 *  0:	write cache
					 */
	words130_155	[26]uint16	/* reserved vendor words 130-155 */
	word156	uint16	/* reserved vendor word 156 */
	words157_159	[3]uint16	/* reserved vendor words 157-159 */
	cfa_power	uint16	/* (word 160) CFA Power Mode
					 * 15 word 160 supported
					 * 14 reserved
					 * 13
					 * 12
					 * 11:0
					 */
	words161_175	[15]uint16	/* Reserved for CFA */
	words176_205	[30]uint16	/* Current Media Serial Number */
	words206_254	[49]uint16	/* reserved words 206-254 */
	integrity_word	uint16	/* (word 255)
					 * 15:8 Checksum
					 *  7:0 Signature
					 */
}

const (
	HDIO_GETGEO       = 0x0301
	HDIO_GET_IDENTITY = 0x030d
	BLKGETSIZE64      = 0x80081272
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

// https://github.com/noguxun/hdparm/blob/master/identify.c
func hdio_get_identity(f *os.File) (hd_driveid, error) {
	var id hd_driveid
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), HDIO_GET_IDENTITY, uintptr(unsafe.Pointer(&id)))
	if e != 0 {
		return hd_driveid{}, e
	}
	return id,nil
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

func readVendorModel(name string) (string, string) {
	var vendor string
	var model string

	// TODO fixup device name cciss/c0d0 -> cciss!c0d0
	// OMG!!! XXX
	if path.Dir(name) == "/dev/cciss" {
		name = strings.Replace(name, "cciss/", "cciss!", -1)
	}

	var blockdevice string = path.Base(name)

	buf, err := ioutil.ReadFile("/sys/block/" + blockdevice + "/device/vendor")
	if err == nil {
		vendor = strings.Trim(string(buf), " \n")
	}

	buf, err = ioutil.ReadFile("/sys/block/" + blockdevice + "/device/model")
	if err == nil {
		model = strings.Trim(string(buf), " \n")
	}

	return vendor,model
}

// Original code from: http://git.cgsecurity.org/cgit/testdisk/tree/src/hdaccess.c
func scan() {
	return // TODO scan for real...

	// IDE
	// - "/dev/hd[a-g]"
	for i := 'a'; i <= 'g'; i++ {
		fmt.Printf("/dev/hd%c\n", i)
	}
	// SCSI
	// - "/dev/sd[a-z]"
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("/dev/sd%c\n", i)
	}
	// RAID Compaq
	// - "/dev/ida/c[0-7]d[0-7]"
	// TODO check if dir /dev/id exists... saves 64 fstats!
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			fmt.Printf("/dev/ida/c%dd%d\n", i, j)
		}
	}
	// CCISS
	// - "/dev/cciss/c0d[0-7]"
	// TODO check if dir /dev/cciss exists... saves 8 fstats!
	for i := 0; i <= 7; i++ {
		fmt.Printf("/dev/cciss/c0d%d\n", i)
	}
	// Device RAID
	// - "/dev/rd/c0d[0-9]"
	// TODO check if dir /dev/ataraid exists... saves 10 fstats!
	for i := 0; i <= 9; i++ {
		fmt.Printf("/dev/rd/c0d%d\n", i)
	}
	// Device RAID IDE
	// - "/dev/ataraid/d[0-14]"
	// TODO check if dir /dev/ataraid exists... saves 15 fstats!
	for i := 0; i <= 14; i++ {
		fmt.Printf("/dev/ataraid/d%d\n", i)
	}
	// Parallel port IDE
	// - "/dev/pd[a-d]"
	for i := 'a'; i <= 'd'; i++ {
		fmt.Printf("/dev/pd%c\n", i)
	}
	// I2O
	// - "/dev/i2o/hd[a-z]"
	// TODO check if dir /dev/ataraid exists... saves 26 fstats!
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("/dev/i2o/hd%c\n", i)
	}
	// Memory card
	// - "/dev/mmcblk[0-9]"
	for i := 0; i <= 9; i++ {
		fmt.Printf("/dev/mmcblk%d\n", i)
	}
	//// Glob (TODO)
	// - "/dev/mapper/*"
	// Software Raid (partition level)
	// - "/dev/md*",
	// - "/dev/sr?"
	// Software (ATA)Raid configured (disk level) via dmraid
	// - "/dev/dm-*"
	// Xen virtual disks
	// - "/dev/xvd?"
}

func read(d *Disk) error {
	d.Vendor, d.Model = readVendorModel(d.Path)
	d.Size.Bytes, _ = blkgetsize64(d.file)

	id, err := hdio_get_identity(d.file)
	if err == nil {
		d.Serial = strings.Trim(string(id.serial_no[:]), " \n")
	}

	geometry, err := hdio_getgeo(d.file)
	if err == nil {
		d.Geometry.Heads     = uint64(geometry.heads)
		d.Geometry.Sectors   = uint64(geometry.sectors)
		d.Geometry.Cylinders = geometry.cylinders
		d.Geometry.Start     = geometry.start
	}

	return nil
}
