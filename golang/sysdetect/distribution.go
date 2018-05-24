package sysdetect

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// nolint
const (
	DistributionDebian  = "Debian"
	DistributionFreeBSD = "FreeBSD"
	DistributionFreeNAS = "FreeNAS"
	DistributionMacOSX  = "Mac OS X"
)

const (
	// DistributionFileOSRelease from https://www.freedesktop.org/software/systemd/man/os-release.html
	DistributionFileOSRelease = "/etc/os-release"
	// DistributionFileOSReleaseAlt alternative path for DistributionFileOSRelease
	DistributionFileOSReleaseAlt = "/usr/lib/os-release"
	// DistributionFileMachineID for linux machines with systemd
	DistributionFileMachineID = "/etc/machine-id"
	// DistributionFileDarwinSystemVersion for Darwin
	DistributionFileDarwinSystemVersion = "/System/Library/CoreServices/SystemVersion.plist"
	// DistributionFileDarwinServerVersion for Darwin server edition
	DistributionFileDarwinServerVersion = "/System/Library/CoreServices/ServerVersion.plist"
	// DistributionFileFreeBSDFreeNasConf for FreeNAS (FreeBSD) systems
	DistributionFileFreeBSDFreeNasConf = "/etc/freenas.conf"
	// DistributionFileFreeBSDFreeNasVersion for FreeNAS version
	DistributionFileFreeBSDFreeNasVersion = "/etc/version"
)

func parseDistributionFileOSRelease(r io.Reader) map[string]string {
	entries := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		values := strings.Split(scanner.Text(), "=")
		if len(values) != 2 {
			continue
		}
		values[1] = strings.Trim(values[1], `"`)
		entries[values[0]] = values[1]
	}
	return entries
}

func distributionFileParse(filename string, r io.Reader) {
	switch filename {
	case DistributionFileOSReleaseAlt:
		fallthrough
	case DistributionFileOSRelease:
		entries := parseDistributionFileOSRelease(r)
		for k, v := range entries {
			fmt.Printf(`"%s"="%s"%c`, k, v, '\n')
		}
	case DistributionFileDarwinSystemVersion:
	case DistributionFileDarwinServerVersion:
	}
}
