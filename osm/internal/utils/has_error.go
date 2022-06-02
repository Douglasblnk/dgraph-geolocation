package utils

import "log"

func HasError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
