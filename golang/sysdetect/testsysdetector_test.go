package sysdetect

type TestSysDetector struct {
}

func (sd *TestSysDetector) AddFile(filepath, content string) {
}

func (sd *TestSysDetector) ReadFile(filename string) (string, error) {
	return "", nil
}

func (sd *TestSysDetector) Sysname() string {
	return ""
}

func (sd *TestSysDetector) RunCommand(name string, arg ...string) (string, error) {
	return "", nil
}

func (sd *TestSysDetector) LookupEnv(key string) (string, bool) {
	return "", false
}

var _ SysDetector = &TestSysDetector{}
