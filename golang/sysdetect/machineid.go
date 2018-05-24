package sysdetect

// Copyright 2017 Jerry Jacobs. All rights reserved.
//   Integrate into github.com/xor-gate/sysdetect
// Copyright 2017 Denis Brodbeck. All rights reserved.
//   Original work located at https://github.com/denisbrodbeck/machineid
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
	// FreeNAS hostid file
	hostidPath = "/etc/hostid"
	// Windows registry key for MachineGuid
	registryKey = `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography`
)

func machineIDForBSD(sd SysDetector) (string, error) {
	kenv, err := sd.RunCommand("kenv", "-q", "smbios.system.uuid")
	if err == nil && len(kenv) > 0 {
		return strings.TrimSpace(kenv), nil
	}

	hostid, err := sd.ReadFile(hostidPath)
	if err == nil {
		return strings.Trim(hostid, "\n"), err
	}

	return "", err
}

func machineIDForLinux(sd SysDetector) (string, error) {
	id, err := sd.ReadFile(dbusPath)
	if err != nil {
		// try fallback path
		id, err = sd.ReadFile(dbusPathEtc)
	}

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(id), nil
}

func machineIDForDarwin(sd SysDetector) (string, error) {
	buf, err := sd.RunCommand("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(buf, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `"`), nil
			}
		}
	}
	return "", nil
}

func machineIDForWindows(sd SysDetector) (string, error) {
	_, err := sd.RunCommand("regedit", "/e", "boem.reg", registryKey)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadFile("boem.reg")
	if err != nil {
		return "", err
	}
	os.Remove("boem.reg")

	re1, err := regexp.Compile(`"MachineGuid"="(.*)"`)
	if err != nil {
		return "", err
	}

	result := re1.FindStringSubmatch(string(data))
	if len(result) == 2 {
		// XXX: Probably this is not the most reliable way to fetch the MachineGuid match
		return result[1], nil
	}
	return "", fmt.Errorf("unable to fetch MachineGuid")
}

// MachineID gets the identifier of the machine (e.g UUID)
func MachineID(sd SysDetector, sysname string) (string, error) {
	switch sysname {
	case SysnameDarwin:
		return machineIDForDarwin(sd)
	case SysnameLinux:
		return machineIDForLinux(sd)
	case SysnameFreeBSD, SysnameNetBSD, SysnameOpenBSD:
		return machineIDForBSD(sd)
	case SysnameWindows:
		return machineIDForWindows(sd)
	}
	return "", nil
}
