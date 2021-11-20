package main

import (
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	// log.Info("Hello from Seelog!")

	logger, err := log.LoggerFromConfigAsFile("seelog2.xml")
	if err != nil {
		return
	}

	log.ReplaceLogger(logger)
	log.Info("Hello from Seelog!")
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}
