package appversion

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var (
	// Version is the application's version which must be set by the build system
	Version = "unknown"
	// BuildDate is the application' build timestamp wich must be set by the build system
	BuildDate = "unknown"
)

// String formats the application's version and build date and returns the result as a string
func String() string {
	return fmt.Sprintf("%s %s (built with %s at %s)", path.Base(os.Args[0]), Version, runtime.Version(), BuildDate)
}

// Print formats the application's version and build date and prints the result
func Print() {
	fmt.Println(String())
}
