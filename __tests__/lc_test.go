package leancloud

import (
	"fmt"
	"testing"

	"github.com/afeiship/go-leancloud"
)

func TestGet(f *testing.T) {
	// init
	res := leancloud.Get("60f77c8e85071346450995d3")
	fmt.Println("result: ", res)
}
