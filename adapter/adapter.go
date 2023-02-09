package 适配器 // 数据适配器包 英文名:adapter

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
)

// 数据适配器.
type adapter struct {
	基.W基本
}

// 数据适配器_增加引用计数.
// 英文名:AddRef
func (a *adapter) AddRef() int {
	return xc.XAd_AddRef(a.W句柄)
}

// 数据适配器_释放引用计数.英文名:Release
func (a *adapter) Release() int {
	return xc.XAd_Release(a.W句柄)
}

// 数据适配器_取引用计数. 英文名:GetRefCount
func (a *adapter) GetRefCount() int {
	return xc.XAd_GetRefCount(a.W句柄)
}

// 数据适配器_销毁. 英文名:Destroy
func (a *adapter) Destroy() int {
	return xc.XAd_Destroy(a.W句柄)
}

// 数据适配器_启用自动销毁.
//
// bEnable: 是否启用. 英文名:EnableAutoDestroy
func (a *adapter) EnableAutoDestroy(bEnable bool) int {
	return xc.XAd_EnableAutoDestroy(a.W句柄, bEnable)
}
