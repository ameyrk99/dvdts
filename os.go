package main

import (
	"io/ioutil"
	"runtime"
	"strings"
)

/* Get the name of OS/Distro to display as text */
func getOsName() (osName string) {
	f, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return strings.Title(runtime.GOOS)
	}

	osRelease := make(map[string]string)

	for _, l := range strings.Split(string(f), "\n") {
		kv := strings.SplitN(l, "=", 2)
		if len(kv) != 2 {
			continue
		}

		osRelease[kv[0]] = strings.Trim(kv[1], `"`)
	}

	if v, ok := osRelease["NAME"]; ok {
		return v
	}

	if v, ok := osRelease["PRETTY_NAME"]; ok {
		return v
	}

	return strings.Title(runtime.GOOS)
}
