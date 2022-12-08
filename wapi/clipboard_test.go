package wapi_test

import (
	"fmt"
	"github.com/xushushun/xcguicn/wapi"
	"testing"
)

func TestSetClipboardText(t *testing.T) {
	err := wapi.SetClipboardText("SetClipboardText")
	if err != nil {
		fmt.Println(err)
	}
	text, err := wapi.GetClipboardText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
