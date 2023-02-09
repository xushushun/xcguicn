package ani

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
)

// AnimaRotate 动画旋转项.
type AnimaRotate struct {
	基.W基本
}

// 动画旋转_置中心, 设置旋转中心点坐标.
//
// x: 坐标X.
//
// y: 坐标Y.
//
// bOffset: TRUE: 相对于自身中心点偏移, FALSE: 绝对坐标.
func (a *AnimaRotate) SetCenter(x float32, y float32, bOffset bool) bool {
	return xc.XAnimaRotate_SetCenter(a.W句柄, x, y, bOffset)
}
