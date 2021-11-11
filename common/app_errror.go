package common

import (
	"errors"
	"log"
)

var ErrRecordNotFound = errors.New("record not found")

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
