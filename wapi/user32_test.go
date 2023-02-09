package wapi_test

import (
	"fmt"
	"testing"

	"github.com/xushushun/xcguicn/app"
	"github.com/xushushun/xcguicn/wapi"
	"github.com/xushushun/xcguicn/widget"
	"github.com/xushushun/xcguicn/window"
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

func TestGetDesktopWindow(t *testing.T) {
	fmt.Println(wapi.GetDesktopWindow())
}

func TestMessageBoxW(t *testing.T) {
	id := wapi.MessageBoxW(0, "context", "title", wapi.MB_CanaelTryContinue|wapi.MB_IconInformation)
	switch id {
	case wapi.ID_Cancel:
		fmt.Println("Cancel")
	case wapi.ID_TryAgain:
		fmt.Println("TryAgain")
	case wapi.ID_Continue:
		fmt.Println("Continue")
	default:
		fmt.Println(id)
	}
}

func TestFindWindowExW(t *testing.T) {
	fmt.Println(wapi.FindWindowExW(0, 0, "", "任务管理器"))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", ""))
	fmt.Println(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	// 取任务管理器的窗口标题长度
	fmt.Println(wapi.GetWindowTextLengthW(wapi.FindWindowExW(0, 0, "TaskManagerWindow", "")))
}

func TestClientToScreen(t *testing.T) {
	a := app.New(true)
	w := window.New(0, 0, 300, 300, "", 0, xcc.Window_Style_Default)

	pt := xc.POINT{X: 0, Y: 0}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	fmt.Println(pt)

	a.ShowAndRun(w.W句柄)
	a.Exit()
}

func TestGetCursorPos(t *testing.T) {
	a := app.New(true)
	w := window.New(0, 0, 300, 300, "", 0, xcc.Window_Style_Default)

	widget.W新按钮(20, 50, 100, 30, "GetCursorPos", w.W句柄).W事件按钮单击(func(pbHandled *bool) int {
		var pt xc.POINT
		wapi.GetCursorPos(&pt)
		a.MessageBox("GetCursorPos", fmt.Sprintf("x: %d  y: %d", pt.X, pt.Y), xcc.MessageBox_Flag_Ok, w.GetHWND(), xcc.Window_Style_Default)
		return 0
	})

	a.ShowAndRun(w.W句柄)
	a.Exit()
}
