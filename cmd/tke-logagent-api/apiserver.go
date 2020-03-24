package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"
	"tkestack.io/tke/cmd/tke-monitor-api/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	app.NewApp("tke-logagent-api").Run()
}
