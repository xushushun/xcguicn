package tmpl

import (
	"github.com/xushushun/xcguicn/objectbase"
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// ListItemTemplate 列表项模板.
type ListItemTemplate struct {
	基.W基本
}

// 模板_创建, 创建项模板.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func New(nType xcc.ListItemTemp_Type_) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_Create(nType))
	return p
}

// NewByHandle 从句柄创建对象.
func NewByHandle(handle int) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(handle)
	return p
}

// 模板_加载从文件, 列表项模板文件载入.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func NewByXML(nType xcc.ListItemTemp_Type_, pFileName string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_Load(nType, pFileName))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewByZip(nType int, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword))
	return p
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_LoadZipMem(nType, data, pFileName, pPassword))
	return p
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func NewByString(nType xcc.ListItemTemp_Type_, pStringXML string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_LoadFromString(nType, pStringXML))
	return p
}

// 模板_克隆, 复制一份新的项模板, 返回模板对象.
func (l *ListItemTemplate) Clone() *ListItemTemplate {
	p := &ListItemTemplate{}
	p.W置句柄(xc.XTemp_Clone(l.W句柄))
	return p
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
func (l *ListItemTemplate) GetType() xcc.ListItemTemp_Type_ {
	return xc.XTemp_GetType(l.W句柄)
}

// 模板_销毁, 项模板销毁.
func (l *ListItemTemplate) Destroy() bool {
	return xc.XTemp_Destroy(l.W句柄)
}

// 模板_添加根节点.
//
// pNode: 节点指针.
func (l *ListItemTemplate) AddNodeRoot(pNode int) bool {
	return xc.XTemp_AddNodeRoot(l.W句柄, pNode)
}

// 模板_取列表中的节点.
//
// index: 节点位置索引.
func (l *ListItemTemplate) List_GetNode(index int) int {
	return xc.XTemp_List_GetNode(l.W句柄, index)
}

// 项模板_列表_插入节点.
//
// index: 插入位置索引.
//
// pNode: 节点指针.
func (l *ListItemTemplate) List_InsertNode(index int, pNode int) bool {
	return xc.XTemp_List_InsertNode(l.W句柄, index, pNode)
}

// 项模板_列表_删除节点.
//
// index: 位置索引.
func (l *ListItemTemplate) List_DeleteNode(index int) bool {
	return xc.XTemp_List_DeleteNode(l.W句柄, index)
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
func (l *ListItemTemplate) List_GetCount() int {
	return xc.XTemp_List_GetCount(l.W句柄)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//
// iColSrc: 源列索引.
//
// iColDest: 目标列索引.
func (l *ListItemTemplate) List_MoveColumn(iColSrc int, iColDest int) bool {
	return xc.XTemp_List_MoveColumn(l.W句柄, iColSrc, iColDest)
}

// 模板_加载从文件扩展, 加载列表项模板从文件.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadEx(nType, pFileName, pOutTemp1, pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipEx(nType, pZipFile, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipMemEx(nType, data, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromStringEx(nType, pStringXML, pOutTemp1, pOutTemp2)
}
