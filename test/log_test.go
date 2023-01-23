package test

import (
	"GoLog/log"
	"testing"
)

func TestLog(t *testing.T) {
	msg := "hello, world"
	log.I(msg)
	log.S(msg)
	log.W(msg)
	log.E(msg)
	log.F(msg)
	log.S(msg)
}
