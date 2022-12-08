package objectbase

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// ObjectBase 炫彩对象基类.
type ObjectBase struct {
	Handle int // 句柄.
}

// 给本类的Handle赋值.
func (o *ObjectBase) SetHandle(handle int) {
	o.Handle = handle
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
func (o *ObjectBase) GetType() xcc.XC_OBJECT_TYPE {
	return xc.XObj_GetType(o.Handle)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
func (o *ObjectBase) GetTypeBase() int {
	return xc.XObj_GetTypeBase(o.Handle)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) GetTypeEx() xcc.XC_OBJECT_TYPE_EX {
	return xc.XObj_GetTypeEx(o.Handle)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) SetTypeEx(nType xcc.XC_OBJECT_TYPE_EX) int {
	return xc.XObj_SetTypeEx(o.Handle, nType)
}
