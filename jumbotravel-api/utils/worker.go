package utils

import "os"

func IsWorker() bool {
	return os.Getenv("PWS_WORKER") != ""
}
