# go-leancloud
> Leancloud sdk for golang.

## installation
```sh
go get -u github.com/afeiship/go-leancloud
```

## usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/afeiship/go-leancloud/leancloud"
)

var client = leancloud.New()

func main() {
	// get
	result, err := client.Get("my_key")
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	fmt.Printf("Full result: %+v\n", result)

	// set
	err = client.Set("my_key", "my_value")
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	// value
	value, err := result.Value()
	if err != nil {
		log.Fatalf("Failed to get value: %v", err)
	}
	fmt.Printf("Value: %v\n", value)
}
```

## resources
- https://github.com/afeiship/leancloud