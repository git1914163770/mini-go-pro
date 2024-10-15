package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestWeb(t *testing.T) {
	pattern := "/web/example/run"
	vs := strings.Split(pattern, "/")
	fmt.Println(vs)
	for _, item := range vs {
		fmt.Print(item)
	}
}
