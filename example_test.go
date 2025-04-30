package timberjack

import (
	"log"
	"time"
)

// To use timberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func Example() {
	log.SetOutput(&Logger{
		Filename:         "/var/log/myapp/foo.log",
		MaxSize:          500,            // megabytes
		MaxBackups:       3,              // number of backups
		MaxAge:           28,             // days
		Compress:         true,           // disabled by default
		LocalTime:        true,           // use the local timezone
		RotationInterval: time.Hour * 24, // rotate daily
	})
}
