// +build windows

package configdir

import (
	"os"
)

var (
	// SystemWide points to system-wide configuration path
	// On Windows systems, it is %PROGRAMDATA%
	SystemWide = os.Getenv("PROGRAMDATA")

	// PerUser points to per-user configuration location
	// On Windows systems, it is %APPDATA%
	PerUser = os.Getenv("APPDATA")
)
