package ani

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
)

// AnimaItem 动画项.
type AnimaItem struct {
	基.W基本
}

// 动画项_启用完成释放, 当动画项完成后自动释放.
//
// 例如对多个动画序列进行渐近式延迟, 在动画序列头标添加延时项(时间差), 当延时项完成时自动释放, 后续动画循环就形成一种时间差(因为对齐的时间差销毁了, 他们永远无法对齐时间).
//
// bEnable: 是否启用.
func (a *AnimaItem) EnableCompleteRelease(bEnable bool) int {
	return xc.XAnimaItem_EnableCompleteRelease(a.W句柄, bEnable)
}

// 动画项_置回调.
//
// callback: 回调函数.
func (a *AnimaItem) SetCallback(callback interface{}) int {
	return xc.XAnimaItem_SetCallback(a.W句柄, callback)
}

// 动画项_置用户数据.
//
// nUserData: 用户数据.
func (a *AnimaItem) SetUserData(nUserData int) int {
	return xc.XAnimaItem_SetUserData(a.W句柄, nUserData)
}

// 动画项_取用户数据, 返回用户数据.
func (a *AnimaItem) GetUserData() int {
	return xc.XAnimaItem_GetUserData(a.W句柄)
}

// 动画项_启用自动销毁.
//
// bEnable: 是否启用.
func (a *AnimaItem) EnableAutoDestroy(bEnable bool) int {
	return xc.XAnimaItem_EnableAutoDestroy(a.W句柄, bEnable)
}
