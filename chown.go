// +build !linux

package timberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
