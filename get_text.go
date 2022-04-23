package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

/* Get the name of OS/Distro or piped input */
func getDisplayText() (displayText string) {

	/* Check for piped input */
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		return string(bytes)
	}

	/* In case of no piped input, get OS name */
	f, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return strings.ToTitle(runtime.GOOS)
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

	return strings.ToTitle(runtime.GOOS)
}
