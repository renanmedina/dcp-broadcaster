package main

import "flag"

const (
	MODE_WEBSERVER    = "webserver"
	MODE_WORKER       = "worker"
	MODE_QUEUE_WORKER = "queue-worker"
	MODES_HELP        = "worker, queue-worker or webserver"
)

func getModeFlag() string {
	mode := flag.String("mode", MODE_WORKER, MODES_HELP)
	flag.Parse()
	return *mode
}
