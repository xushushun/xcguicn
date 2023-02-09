package widget

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// W按钮 按钮.
// 英文名:Button.
type W按钮 struct {
	Element
}

// W新按钮
// 英文名:XBtn_Create
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 标题.
//
// hParent: 父为窗口句柄或元素句柄.
func W新按钮(x int, y int, cx int, cy int, pName string, hParent int) *W按钮 {
	p := &W按钮{}
	p.W置句柄(xc.XBtn_Create(x, y, cx, cy, pName, hParent))
	return p
}

// W新按钮自句柄 从句柄创建对象.
// 英文名:W新按钮自句柄.
func W新按钮自句柄(handle int) *W按钮 {
	p := &W按钮{}
	p.W置句柄(handle)
	return p
}

// W新按钮自名字 从name创建对象, 失败返回nil.
// 英文名:NewButtonByName.
func W新按钮自名字(name string) *W按钮 {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &W按钮{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// W新按钮自UID 从UID创建对象, 失败返回nil.
// 英文名:XC_GetObjectByUID.
func W新按钮自UID(nUID int) *W按钮 {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &W按钮{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// W新按钮自UID名 从UID名称创建对象, 失败返回nil.
// 英文名:XC_GetObjectByUIDName
func W新按钮自UID名(name string) *W按钮 {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &W按钮{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// W是否选中 按钮_判断选中, 是否选中状态.
// 英文名:XBtn_IsCheck.
func (b *W按钮) W是否选中() bool {
	return xc.XBtn_IsCheck(b.W句柄)
}

// W置选中 按钮_置选中, 设置选中状态.
// 英文名: XBtn_SetCheck.
// bCheck: W是否选中.
func (b *W按钮) W置选中(bCheck bool) bool {
	return xc.XBtn_SetCheck(b.W句柄, bCheck)
}

// W置状态 按钮_置状态.
// 英文名: XBtn_SetState.
//
//	@param nState 按钮状态: xcc.Common_State3_.
//	@return int
func (b *W按钮) W置状态(nState xcc.Common_State3_) int {
	return xc.XBtn_SetState(b.W句柄, nState)
}

// W取状态 按钮_取状态.
// 英文名: XBtn_GetState.
//
//	@return xcc.Common_State3_
func (b *W按钮) W取状态() xcc.Common_State3_ {
	return xc.XBtn_GetState(b.W句柄)
}

// W取状态扩展 按钮_取状态扩展.
// 英文名: XBtn_GetStateEx.
//
//	@return xcc.Button_State_
func (b *W按钮) W取状态扩展() xcc.Button_State_ {
	return xc.XBtn_GetStateEx(b.W句柄)
}

// W置类型扩展 , 设置按钮类型并自动修改样式和文本对齐方式.
// 英文名: XBtn_SetTypeEx.
// nType: 按钮类型, Button_Type_ , element_type_ , xc_ex_error.
func (b *W按钮) W置类型扩展(nType xcc.XC_OBJECT_TYPE_EX) int {
	return xc.XBtn_SetTypeEx(b.W句柄, nType)
}

// W置组ID 按钮_置组ID.
// 英文名: XBtn_SetGroupID.
// nID: 组ID.
func (b *W按钮) W置组ID(nID int) int {
	return xc.XBtn_SetGroupID(b.W句柄, nID)
}

// W取组ID 按钮_取组ID.
// 英文名: XBtn_GetGroupID.
func (b *W按钮) W取组ID() int {
	return xc.XBtn_GetGroupID(b.W句柄)
}

// W置绑定元素 按钮_置绑定元素.
// 英文名: XBtn_SetBindEle.
// hBindEle: 将要绑定的元素.
func (b *W按钮) W置绑定元素(hBindEle int) int {
	return xc.XBtn_SetBindEle(b.W句柄, hBindEle)
}

// W取绑定元素 按钮_取绑定元素, 返回: 绑定的元素句柄.
// 英文名: XBtn_GetBindEle.
func (b *W按钮) W取绑定元素() int {
	return xc.XBtn_GetBindEle(b.W句柄)
}

// W置文本对齐 按钮_置文本对齐.
// 英文名: XBtn_SetTextAlign.
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
func (b *W按钮) W置文本对齐(nFlags xcc.TextFormatFlag_) int {
	return xc.XBtn_SetTextAlign(b.W句柄, nFlags)
}

// W取文本对齐 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
// 英文名: XBtn_GetTextAlign.
func (b *W按钮) W取文本对齐() xcc.TextFormatFlag_ {
	return xc.XBtn_GetTextAlign(b.W句柄)
}

// W置图标对齐 按钮_置图标对齐.
// 英文名: XBtn_SetIconAlign.
// align: 对齐方式, Button_Icon_Align_.
func (b *W按钮) W置图标对齐(align xcc.Button_Icon_Align_) int {
	return xc.XBtn_SetIconAlign(b.W句柄, align)
}

// W置文本偏移 按钮_置偏移, 设置按钮文本坐标偏移量.
// 英文名: XBtn_SetOffset.
// x: 偏移量.
//
// y: 偏移量.
func (b *W按钮) W置文本偏移(x int, y int) int {
	return xc.XBtn_SetOffset(b.W句柄, x, y)
}

// W置图标偏移 按钮_置图标偏移, 设置按钮图标坐标偏移量.
// 英文名: XBtn_SetOffsetIcon.
// x: 偏移量.
//
// y: 偏移量.
func (b *W按钮) W置图标偏移(x int, y int) int {
	return xc.XBtn_SetOffsetIcon(b.W句柄, x, y)
}

// W置图标间隔 按钮_置图标间隔, 设置图标与文本间隔大小.
// 英文名: XBtn_SetIconSpace.
// size: 间隔大小.
func (b *W按钮) W置图标间隔(size int) int {
	return xc.XBtn_SetIconSpace(b.W句柄, size)
}

// W置文本 按钮_置文本.
// 英文名: XBtn_SetText.
// pName: 文本内容.
func (b *W按钮) W置文本(pName string) int {
	return xc.XBtn_SetText(b.W句柄, pName)
}

// W取文本
// 英文名: XBtn_GetText.
func (b *W按钮) W取文本() string {
	return xc.XBtn_GetText(b.W句柄)
}

// W置图标 按钮_置图标.
// 英文名: XBtn_SetIcon.
// hImage: 图片句柄.
func (b *W按钮) W置图标(hImage int) int {
	return xc.XBtn_SetIcon(b.W句柄, hImage)
}

// W置禁用图标 按钮_置禁用图标.
// 英文名: XBtn_SetIconDisable.
// hImage: 图片句柄.
func (b *W按钮) W置禁用图标(hImage int) int {
	return xc.XBtn_SetIconDisable(b.W句柄, hImage)
}

// W取图标 按钮_取图标, 返回图标句柄.
// 英文名: XBtn_GetIcon.
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func (b *W按钮) W取图标(nType int) int {
	return xc.XBtn_GetIcon(b.W句柄, nType)
}

// W加动画帧 按钮_添加动画帧.
// 英文名: XBtn_AddAnimationFrame.
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.
func (b *W按钮) W加动画帧(hImage int, uElapse int) int {
	return xc.XBtn_AddAnimationFrame(b.W句柄, hImage, uElapse)
}

// W启用动画 按钮_启用动画, 开始或关闭图片动画的播放.
// 英文名: XBtn_EnableAnimation.
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.
func (b *W按钮) W启用动画(bEnable bool, bLoopPlay bool) int {
	return xc.XBtn_EnableAnimation(b.W句柄, bEnable, bLoopPlay)
}

/*
下面都是事件
*/

type W函数按钮点击 func(pbHandled *bool) int                         // 按钮点击事件.
type W函数按钮点击1 func(hEle int, pbHandled *bool) int              // 按钮点击事件.
type W函数按钮选中 func(bCheck bool, pbHandled *bool) int            // 按钮选中事件.
type W函数按钮选中1 func(hEle int, bCheck bool, pbHandled *bool) int // 按钮选中事件.

// W事件按钮单击 事件_按钮被单击.
// 英文名: XEle_RegEventC.
func (b *W按钮) W事件按钮单击(pFun W函数按钮点击) bool {
	return xc.XEle_RegEventC(b.W句柄, xcc.XE_BNCLICK, pFun)
}

// W事件按钮单击1 事件_按钮被单击1.
// 英文名: XEle_RegEventC1.
func (b *W按钮) W事件按钮单击1(pFun W函数按钮点击1) bool {
	return xc.XEle_RegEventC1(b.W句柄, xcc.XE_BNCLICK, pFun)
}

// 按钮选中事件.
// 英文名: XEle_RegEventC.
func (b *W按钮) W事件按钮选中(pFun W函数按钮选中) bool {
	return xc.XEle_RegEventC(b.W句柄, xcc.XE_BUTTON_CHECK, pFun)
}

// W事件按钮选中1 按钮选中事件.
// 英文名: XEle_RegEventC1.
func (b *W按钮) W事件按钮选中1(pFun W函数按钮选中1) bool {
	return xc.XEle_RegEventC1(b.W句柄, xcc.XE_BUTTON_CHECK, pFun)
}
