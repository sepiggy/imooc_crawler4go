package common

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
