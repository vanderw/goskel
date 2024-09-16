/*
Build:

	Build with flag
		-ldflags="-X 'mport/version.BuildGoVersion=$(shell go version)'"

Usage:

`

	func main() {
		flag.Parse()
		if version.CheckArg() {
			os.Exit(0)
		}
	}

`

Example output:
```
$ ./example.exe version
Version: 1.1.2
Build Date: 2023-01-03T14:57:12
Go Version: go version go1.19.3 windows/amd64
Build OS: MINGW64_NT-10.0-19045
Git branch: dev
Commit ID: 1933db9a83d2972c499709a38f2a044e712684b2
```
*/
package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	BuildVersion    string = "unset"
	BuildDate              = "unset"
	BuildGoVersion         = "unknown"
	BuildPlatform          = "unknown"
	BuildBranch            = "unknown"
	BuildCommitHash        = "unknown"

	// command flag
	printVer = false
)

func init() {
	flag.BoolVar(&printVer, "v", false, "print version string")
}

func CheckArg() bool {
	if printVer {
		DisplayVersion()
	}
	return printVer
}

func DisplayVersion() {
	fmt.Fprintf(os.Stdout, `
 Version: %v
 Build Date: %v
 Go Version: %v
 Build OS: %v
 Git branch: %v
 Commit ID: %v
 `, BuildVersion, BuildDate, BuildGoVersion, BuildPlatform, BuildBranch, BuildCommitHash)
}
