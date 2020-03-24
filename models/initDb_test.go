package models

import (
	"fmt"
	"log"
	"testing"
)

func TestSetConnect(t *testing.T) {
	mc := SetConnect()
	err := mc.Ping(nil, nil)
	if err != nil {
		fmt.Println("err", err.Error())
		log.Print(err)
		t.Fail()
	} else {
		fmt.Println("ok")
	}

}
