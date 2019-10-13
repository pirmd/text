// +build !windows,!darwin

package configdir

import (
	"os"
	"path/filepath"
	"strings"
)

//TODO(pirmd): or /usr/local/etc (?)

var (
	// SystemWide points to system-wide configuration path
	// On Unix systems, it is either /etc or first XDG_CONFIG_DIRS
	SystemWide = "/etc"

	// PerUser points to per-user configuration location
	// On Unix systems, it is either $HOME/.config or XDG_CONFIG_HOME
	PerUser = filepath.Join(os.Getenv("HOME"), ".config")
)

func init() {
	//TODO(pirmd): switch to go1.13 os.UserConfigDir()
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		PerUser = os.Getenv("XDG_CONFIG_HOME")
	}

	if os.Getenv("XDG_CONFIG_DIRS") != "" {
		SystemWide = strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")[0]
	}
}
