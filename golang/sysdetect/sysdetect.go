package sysdetect

import (
	"fmt"
	"strings"
)

// SysDetector for
type SysDetector interface {
	// ReadFile from filename
	ReadFile(filename string) (string, error)

	// LookupEnv retrieves the value of the environment variable named by the key.
	// If the variable is present in the environment the value (which may be empty) is
	//  returned and the boolean is true. Otherwise the returned value will be empty
	//  and the boolean will be false.
	LookupEnv(key string) (string, bool)

	// RunCommand executes a command with name and optional arg
	RunCommand(name string, arg ...string) (string, error)
}

// Detect system information with a detector sd
func Detect(sd SysDetector) map[string]string {
	result := make(map[string]string)
	sysname := Sysname(sd)

	switch sysname {
	case SysnameDarwin:
		detectDarwin(sd, result)
	case SysnameFreeBSD:
		detectFreeBSD(sd, result)
	case SysnameLinux:
		detectLinux(sd)
	case SysnameWindows:
	default:
		return nil
	}

	result["SYSNAME"] = sysname
	if machineID, err := MachineID(sd, sysname); err == nil {
		result["MACHINE_ID"] = machineID
	}
	if arch, err := sd.RunCommand("uname", "-m"); err == nil && arch != "" {
		result["ARCHITECTURE"] = arch
	}
	detectLookupEnv(sd, result)
	return result
}

// Sysname gets the system name
func Sysname(sd SysDetector) string {
	if uname, err := sd.RunCommand("uname", "-s"); err == nil {
		return strings.Trim(uname, " \r\n")
	}

	ver, err := sd.RunCommand("cmd", "/K", "ver")
	if err == nil {
		if strings.Contains(ver, "Windows") {
			return SysnameWindows
		}
	}
	fmt.Println(err)

	return SysnameUnknown
}

func detectLookupEnv(sd SysDetector, result map[string]string) {
	if user, ok := sd.LookupEnv("USER"); ok {
		result["USER"] = user
	}
	if shell, ok := sd.LookupEnv("SHELL"); ok {
		result["SHELL"] = shell
	}
	if home, ok := sd.LookupEnv("HOME"); ok {
		result["HOME"] = home
	}
}

func detectDarwin(sd SysDetector, result map[string]string) {
	if productVersion, err := sd.RunCommand("sw_vers", "-productVersion"); err == nil {
		result["VERSION_ID"] = productVersion
	}
	if buildVersion, err := sd.RunCommand("sw_vers", "-buildVersion"); err == nil {
		result["BUILD_ID"] = buildVersion
	}
	if productName, err := sd.RunCommand("sw_vers", "-productName"); err == nil {
		result["NAME"] = productName
	}
}

func detectFreeBSD(sd SysDetector, result map[string]string) {
	if version, err := sd.ReadFile(DistributionFileFreeBSDFreeNasVersion); err == nil {
		result["VERSION"] = version
	}
}

func detectLinux(sd SysDetector) {
	osrelease, err := sd.ReadFile(DistributionFileOSRelease)
	if err == nil {
		distributionFileParse(DistributionFileOSRelease, strings.NewReader(osrelease))
	}
}
