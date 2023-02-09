package widget

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// LayoutFrame 布局框架.
type LayoutFrame struct {
	ScrollView
}

// NewLayoutFrame 布局框架_创建.
//
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutFrame
func NewLayoutFrame(x int, y int, cx int, cy int, hParent int) *LayoutFrame {
	p := &LayoutFrame{}
	p.W置句柄(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
	return p
}

// NewLayoutFrameByHandle 从句柄创建对象.
//
//	@param handle
//	@return *LayoutFrame
func NewLayoutFrameByHandle(handle int) *LayoutFrame {
	p := &LayoutFrame{}
	p.W置句柄(handle)
	return p
}

// NewLayoutFrameByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
func NewLayoutFrameByName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUID 从UID创建对象, 失败返回nil.
//
//	@param nUID
//	@return *LayoutFrame
func NewLayoutFrameByUID(nUID int) *LayoutFrame {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutFrame{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
func NewLayoutFrameByUIDName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// ShowLayoutFrame 布局框架_显示布局边界.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayoutFrame_ShowLayoutFrame(l.W句柄, bEnable)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.W句柄, bEnable)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.W句柄, bEnable)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.W句柄, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutFrame) SetAlignH(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignH(l.W句柄, nAlign)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutFrame) SetAlignV(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignV(l.W句柄, nAlign)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
func (l *LayoutFrame) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {
	return xc.XLayoutBox_SetAlignBaseline(l.W句柄, nAlign)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
func (l *LayoutFrame) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.W句柄, nSpace)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
func (l *LayoutFrame) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.W句柄, nSpace)
}
