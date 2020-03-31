package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"
	"tkestack.io/tke/cmd/tke-logagent-api/app"

	"tkestack.io/tke/pkg/util/log"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	log.Infof("====================================================================")
	app.NewApp("tke-logagent-api").Run()
}
