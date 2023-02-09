package svg

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
)

// SVG矢量图形.
type Svg struct {
	基.W基本
}

// SVG_加载从文件, 返回Svg对象.
//
// pFileName: 文件名.
func NewByFile(pFileName string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadFile(pFileName))
	return p
}

// SVG_加载从字符串, 返回Svg对象.
//
// pString: 字符串.
func NewByString(pString string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadString(pString))
	return p
}

// SVG_加载从字符串W.
//
// pString: 字符串.
func NewByStringW(pString string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadStringW(pString))
	return p
}

// SVG_加载从字符串UTF8.
//
// pString: 字符串.
func NewByStringUtf8(pString string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadStringUtf8(pString))
	return p
}

// SVG_加载从ZIP, 返回Svg对象.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewByZip(pZipFileName, pFileName, pPassword string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadZip(pZipFileName, pFileName, pPassword))
	return p
}

// SVG_加载从内存ZIP, 返回Svg对象.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewByZipMem(data []byte, pFileName, pPassword string) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadZipMem(data, pFileName, pPassword))
	return p
}

// SVG_加载从资源, 返回Svg对象.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func NewByRes(id int, pType string, hModule int) *Svg {
	p := &Svg{}
	p.W置句柄(xc.XSvg_LoadRes(id, pType, hModule))
	return p
}

// SVG_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Svg) SetSize(nWidth, nHeight int) int {
	return xc.XSvg_SetSize(s.W句柄, nWidth, nHeight)
}

// SVG_取大小.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (s *Svg) GetSize(pWidth, pHeight *int) int {
	return xc.XSvg_GetSize(s.W句柄, pWidth, pHeight)
}

// SVG_取宽度.
func (s *Svg) GetWidth() int {
	return xc.XSvg_GetWidth(s.W句柄)
}

// SVG_取高度.
func (s *Svg) GetHeight() int {
	return xc.XSvg_GetHeight(s.W句柄)
}

// SVG_置偏移.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) SetPosition(x, y int) int {
	return xc.XSvg_SetPosition(s.W句柄, x, y)
}

// SVG_取偏移.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) GetPosition(pX, pY *int) int {
	return xc.XSvg_GetPosition(s.W句柄, pX, pY)
}

// SVG_置偏移F.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) SetPositionF(x, y float32) int {
	return xc.XSvg_SetPositionF(s.W句柄, x, y)
}

// SVG_取偏移F.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) GetPositionF(pX, pY *float32) int {
	return xc.XSvg_GetPositionF(s.W句柄, pX, pY)
}

// SVG_取视图框.
//
// pViewBox: 接收返回视图框.
func (s *Svg) GetViewBox(pViewBox *xc.RECT) int {
	return xc.XSvg_GetViewBox(s.W句柄, pViewBox)
}

// SVG_启用自动销毁.
//
// bEnable: 是否自动销毁.
func (s *Svg) EnableAutoDestroy(bEnable bool) int {
	return xc.XSvg_EnableAutoDestroy(s.W句柄, bEnable)
}

// SVG_增加引用计数.
func (s *Svg) AddRef() int {
	return xc.XSvg_AddRef(s.W句柄)
}

// SVG_释放引用计数.
func (s *Svg) Release() int {
	return xc.XSvg_Release(s.W句柄)
}

// SVG_取引用计数.
func (s *Svg) GetRefCount() int {
	return xc.XSvg_GetRefCount(s.W句柄)
}

// SVG_销毁.
func (s *Svg) Destroy() int {
	return xc.XSvg_Destroy(s.W句柄)
}

// SVG_置透明度.
//
// alpha: 透明度.
func (s *Svg) SetAlpha(alpha byte) int {
	return xc.XSvg_SetAlpha(s.W句柄, alpha)
}

// SVG_取透明度, 返回透明度.
func (s *Svg) GetAlpha() int {
	return xc.XSvg_GetAlpha(s.W句柄)
}

// SVG_置用户填充颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func (s *Svg) SetUserFillColor(color int, bEnable bool) int {
	return xc.XSvg_SetUserFillColor(s.W句柄, color, bEnable)
}

// SVG_置用户笔触颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// strokeWidth: 笔触宽度.
//
// bEnable: 是否有效.
func (s *Svg) SetUserStrokeColor(color int, strokeWidth float32, bEnable bool) int {
	return xc.XSvg_SetUserStrokeColor(s.W句柄, color, strokeWidth, bEnable)
}

// SVG_取用户填充颜色.
//
// pColor: 返回颜色值, AGBR颜色.
func (s *Svg) GetUserFillColor(pColor *int) bool {
	return xc.XSvg_GetUserFillColor(s.W句柄, pColor)
}

// SVG_取用户笔触颜色.
//
// pColor: 返回颜色值, AGBR颜色.
//
// pStrokeWidth: .
func (s *Svg) GetUserStrokeColor(pColor *int, pStrokeWidth *float32) bool {
	return xc.XSvg_GetUserStrokeColor(s.W句柄, pColor, pStrokeWidth)
}

// SVG_置旋转角度, 默认以自身中心点旋转.
//
// angle: 转角度.
func (s *Svg) SetRotateAngle(angle float32) int {
	return xc.XSvg_SetRotateAngle(s.W句柄, angle)
}

// SVG_取旋转角度, 返回旋转角度.
func (s *Svg) GetRotateAngle() float32 {
	return xc.XSvg_GetRotateAngle(s.W句柄)
}

// SVG_置旋转.
//
// angle: 角度.
//
// x: 旋转中心点X.
//
// y: 旋转中心点Y.
//
// bOffset: TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) SetRotate(angle float32, x float32, y float32, bOffset bool) int {
	return xc.XSvg_SetRotate(s.W句柄, angle, x, y, bOffset)
}

// SVG_取旋转.
//
// pAngle: 返回角度.
//
// pX: 返回 旋转中心点X.
//
// pY: 返回 旋转中心点Y.
//
// pbOffset: 返回TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) GetRotate(pAngle *float32, pX *float32, pY *float32, pbOffset *bool) int {
	return xc.XSvg_GetRotate(s.W句柄, pAngle, pX, pY, pbOffset)
}

// SVG_显示, 显示或隐藏.
//
// bShow: 是否显示.
func (s *Svg) Show(bShow bool) int {
	return xc.XSvg_Show(s.W句柄, bShow)
}
