// +build plan9 solaris

package termsize

import (
	"fmt"
)

func Width() (int, error) {
	return -1, fmt.Errorf("Not supported")
}
