package xc_test

import (
	"testing"

	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

func TestXInitXCGUI(t *testing.T) {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	err := xc.SetXcguiPath(`C:\Users\xushushun\GolandProjects\xu-shushun-go\XCGUI.dll`)
	if err != nil {
		panic(err)
	}
	xc.LoadXCGUI()
	xc.XInitXCGUI(false)
	hWindow := xc.XWnd_Create(0, 0, 1000, 500, "go开发游戏测试", 0, xcc.Window_Style_Default)
	//加入动画
	背景资源 := xc.XImage_LoadFile("背景.jpg")
	xc.XImage_SetDrawType(背景资源, xcc.Image_Draw_Type_Stretch)
	背景句柄 := xc.XShapeGif_Create(0, 0, 1000, 500, hWindow)
	xc.XShapeGif_SetImage(背景句柄, 背景资源)
	怪物资源 := xc.XImage_LoadFile("怪物.gif")
	xc.XImage_EnableTranColor(怪物资源, true)
	怪物句柄 := xc.XShapeGif_Create(450, 100, 100, 200, hWindow)
	xc.XShapeGif_SetImage(怪物句柄, 怪物资源)
	宠物资源 := xc.XImage_LoadFile("宠物.gif")
	宠物句柄 := xc.XShapeGif_Create(100, 100, 100, 200, hWindow)
	xc.XShapeGif_SetImage(宠物句柄, 宠物资源)
	人物资源 := xc.XImage_LoadFile("人物.gif")
	人物句柄 := xc.XShapeGif_Create(200, 100, 280, 200, hWindow)
	xc.XShapeGif_SetImage(人物句柄, 人物资源)
	xc.XWnd_EnableDragWindow(hWindow, true)
	xc.XWnd_ShowWindow(hWindow, xcc.SW_SHOW)
	xc.XRunXCGUI()
	xc.XExitXCGUI()
}

func TestWriteDll(t *testing.T) {
	err := xc.WriteDll([]byte("0"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(xc.GetXcguiPath())
}
