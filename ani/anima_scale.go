package ani

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// AnimaScale 动画缩放项.
type AnimaScale struct {
	基.W基本
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// position: 位置, Position_Flag_.
func (a *AnimaScale) SetPosition(position xcc.Position_Flag_) bool {
	return xc.XAnimaScale_SetPosition(a.W句柄, position)
}
