package wapi_test

import (
	"fmt"
	"github.com/xushushun/xcguicn/wapi"
	"testing"
)

func TestSleep(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		wapi.Sleep(1000)
	}
}
