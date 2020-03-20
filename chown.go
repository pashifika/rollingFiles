// +build !linux

package rollingFiles

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
