package env

import "os"

func IsTest() bool {
	return os.Getenv("MOVIEDEMO_ENV") == "TEST"
}
