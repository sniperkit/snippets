package disk

type fsItem struct {
	fsid uint64
	name string
}

var fsItems = []fsItem {
	{ fsid: 0x00006969, name: "nfs"    },
	{ fsid: 0x0000EF53, name: "ext"    },
	{ fsid: 0x00011954, name: "ufs"    },
	{ fsid: 0x01021994, name: "tmpfs"  },
	{ fsid: 0x52654973, name: "reiser" },
	{ fsid: 0x9123683E, name: "btrfs"  },
}

/* Filesystem magic id to string
 * http://man7.org/linux/man-pages/man3/statfs.3.html
 */
func FsName(fsid uint64) string {
	for _, x := range fsItems {
		if x.fsid == fsid {
			return x.name
		}
	}
	return "unknown"
}
