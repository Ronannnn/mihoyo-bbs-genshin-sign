package config

import "time"

func init() {
	time.Local = time.FixedZone("utc", 0) // set timezone to utc
}
