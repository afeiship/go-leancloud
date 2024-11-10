package leancloud

import (
	"fmt"
	"log"
	"testing"

	"github.com/afeiship/go-leancloud"
)

// 初始化 LeanCloud 客户端
var client = leancloud.New()

func TestGet(f *testing.T) {
	// init
	// res := leancloud.Get("60f77c8e85071346450995d3")
	// fmt.Println("result: ", res)

	// 获取指定键的完整数据
	result, err := client.Get("60f77c8e85071346450995d3")
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	fmt.Printf("Full result: %+v\n", result)
}

func TestValue(f *testing.T) {
	// init
	// 获取指定键的值
	value, err := client.Value("60f77c8e85071346450995d3")
	if err != nil {
		log.Fatalf("Failed to get value: %v", err)
	}
	fmt.Printf("Value: %v\n", value)
}
