// +build darwin

package configdir

import (
	"os"
	"path/filepath"
)

var (
	// SystemWide points to system-wide configuration path
	// On Darwin systems, it is /Library/Application Support
	SystemWide = "/Library/Application Support"

	// PerUser points to per-user configuration location
	// On Darwin systems, it is $HOME/Library/Application Support
	PerUser = filepath.Join(os.Getenv("HOME"), SystemWide)
)
