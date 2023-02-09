package 基

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// W基本 炫彩对象基类.
type W基本 struct {
	W句柄 int // 句柄.
}

// 给本类的Handle赋值.
func (o *W基本) W置句柄(句柄 int) {
	o.W句柄 = 句柄
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
func (o *W基本) W取类型() xcc.XC_OBJECT_TYPE {
	return xc.XObj_GetType(o.W句柄)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
func (o *W基本) W取类型基础() int {
	return xc.XObj_GetTypeBase(o.W句柄)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *W基本) W取类型扩展() xcc.XC_OBJECT_TYPE_EX {
	return xc.XObj_GetTypeEx(o.W句柄)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *W基本) W置类型扩展(nType xcc.XC_OBJECT_TYPE_EX) int {
	return xc.XObj_SetTypeEx(o.W句柄, nType)
}
