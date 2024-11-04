package pkg

import (
	"runtime"
	"time"
)

var (
	Version   = "localVersion"
	BuildDate = time.Now().Format("2006-01-02 15:04:05")
	GoVersion = runtime.Version()
)
