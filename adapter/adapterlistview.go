package 适配器

import (
	"github.com/xushushun/xcguicn/xc"
)

// 数据适配器-列表视元素.
type AdapterListView struct {
	adapter
}

// 数据适配器列表视_创建, 创建列表视元素数据适配器.
func NewAdapterListView() *AdapterListView {
	p := &AdapterListView{}
	p.W置句柄(xc.XAdListView_Create())
	return p
}

// 从句柄创建对象.
func NewAdapterListViewByHandle(handle int) *AdapterListView {
	p := &AdapterListView{}
	p.W置句柄(handle)
	return p
}

// 数据适配器列表视_组添加列, 组操作, 添加数据列, 返回列索引.
//
// pName: 字段称.
func (a *AdapterListView) Group_AddColumn(pName string) int {
	return xc.XAdListView_Group_AddColumn(a.W句柄, pName)
}

// 数据适配器列表视_添加组文本, 组操作, 添加组, 数据默认填充到第一列, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemText(pValue string, iPos int) int {
	return xc.XAdListView_Group_AddItemText(a.W句柄, pValue, iPos)
}

// 数据适配器列表视_添加组文本扩展, 组操作, 添加组, 数据填充指定列, 返回组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemTextEx(pName string, pValue string, iPos int) int {
	return xc.XAdListView_Group_AddItemTextEx(a.W句柄, pName, pValue, iPos)
}

// 数据适配器列表视_添加组图片, 组操作, 添加组, 数据默认填充第一列.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemImage(hImage int, iPos int) int {
	return xc.XAdListView_Group_AddItemImage(a.W句柄, hImage, iPos)
}

// 数据适配器列表视_添加组图片扩展, 组操作, 添加组, 数据填充指定列.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Group_AddItemImageEx(pName string, hImage int, iPos int) int {
	return xc.XAdListView_Group_AddItemImageEx(a.W句柄, pName, hImage, iPos)
}

// 数据适配器列表视_组设置文本, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) Group_SetText(iGroup int, iColumn int, pValue string) bool {
	return xc.XAdListView_Group_SetText(a.W句柄, iGroup, iColumn, pValue)
}

// 数据适配器列表视_组设置文本扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func (a *AdapterListView) Group_SetTextEx(iGroup int, pName string, pValue string) bool {
	return xc.XAdListView_Group_SetTextEx(a.W句柄, iGroup, pName, pValue)
}

// 数据适配器列表视_组设置图片, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) Group_SetImage(iGroup int, iColumn int, hImage int) bool {
	return xc.XAdListView_Group_SetImage(a.W句柄, iGroup, iColumn, hImage)
}

// 数据适配器列表视_组设置图片扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (a *AdapterListView) Group_SetImageEx(iGroup int, pName string, hImage int) bool {
	return xc.XAdListView_Group_SetImageEx(a.W句柄, iGroup, pName, hImage)
}

// 数据适配器列表视_项添加列, 项操作, 添加列.
//
// pName: 字段称.
func (a *AdapterListView) Item_AddColumn(pName string) int {
	return xc.XAdListView_Item_AddColumn(a.W句柄, pName)
}

// 数据适配器列表视_取组数量, 组操作, 获取组数量.
func (a *AdapterListView) Group_GetCount() int {
	return xc.XAdListView_Group_GetCount(a.W句柄)
}

// 数据适配器列表视_取组中项数量, 项操作, 获取指定组中项数量.
//
// iGroup: 组索引.
func (a *AdapterListView) Item_GetCount(iGroup int) int {
	return xc.XAdListView_Item_GetCount(a.W句柄, iGroup)
}

// 数据适配器列表视_添加项文本, 项操作, 添加项.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemText(iGroup int, pValue string, iPos int) int {
	return xc.XAdListView_Item_AddItemText(a.W句柄, iGroup, pValue, iPos)
}

// 数据适配器列表视_添加项文本扩展, 项操作, 添加项, 数据填充指定列.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemTextEx(iGroup int, pName string, pValue string, iPos int) int {
	return xc.XAdListView_Item_AddItemTextEx(a.W句柄, iGroup, pName, pValue, iPos)
}

// 数据适配器列表视_添加项图片, 项操作, 添加项.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemImage(iGroup int, hImage int, iPos int) int {
	return xc.XAdListView_Item_AddItemImage(a.W句柄, iGroup, hImage, iPos)
}

// 数据适配器列表视_添加项图片扩展, 项操作, 添加项, 填充指定列数据.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) Item_AddItemImageEx(iGroup int, pName string, hImage int, iPos int) int {
	return xc.XAdListView_Item_AddItemImageEx(a.W句柄, iGroup, pName, hImage, iPos)
}

// 数据适配器列表视_组删除项, 删除组, 自动删除子项.
//
// iGroup: 组索引.
func (a *AdapterListView) Group_DeleteItem(iGroup int) bool {
	return xc.XAdListView_Group_DeleteItem(a.W句柄, iGroup)
}

// 数据适配器列表视_删除全部子项, 删除指定组的所有子项.
//
// iGroup: 组索引.
func (a *AdapterListView) Group_DeleteAllChildItem(iGroup int) int {
	return xc.XAdListView_Group_DeleteAllChildItem(a.W句柄, iGroup)
}

// 数据适配器列表视_删除项, 删除项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (a *AdapterListView) Item_DeleteItem(iGroup int, iItem int) bool {
	return xc.XAdListView_Item_DeleteItem(a.W句柄, iGroup, iItem)
}

// 数据适配器列表视_删除全部, 删除所有(组,项,列).
func (a *AdapterListView) DeleteAll() int {
	return xc.XAdListView_DeleteAll(a.W句柄)
}

// 数据适配器列表视_删除全部组, 删除所有的组.
func (a *AdapterListView) DeleteAllGroup() int {
	return xc.XAdListView_DeleteAllGroup(a.W句柄)
}

// 数据适配器列表视_删除全部项, 删除所有的项.
func (a *AdapterListView) DeleteAllItem() int {
	return xc.XAdListView_DeleteAllItem(a.W句柄)
}

// 数据适配器列表视_组删除列, 删除组的列.
//
// iColumn: 列索引.
func (a *AdapterListView) DeleteColumnGroup(iColumn int) int {
	return xc.XAdListView_DeleteColumnGroup(a.W句柄, iColumn)
}

// 数据适配器列表视_项删除列, 删除项的列.
//
// iColumn: 列索引.
func (a *AdapterListView) DeleteColumnItem(iColumn int) int {
	return xc.XAdListView_DeleteColumnItem(a.W句柄, iColumn)
}

// 数据适配器列表视_项获取文本扩展, 项操作, 获取项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterListView) Item_GetTextEx(iGroup int, iItem int, pName string) string {
	return xc.XAdListView_Item_GetTextEx(a.W句柄, iGroup, iItem, pName)
}

// 数据适配器列表视_项获取图片扩展, 项操作, 获取项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterListView) Item_GetImageEx(iGroup int, iItem int, pName string) int {
	return xc.XAdListView_Item_GetImageEx(a.W句柄, iGroup, iItem, pName)
}

// 数据适配器列表视_项置文本, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) Item_SetText(iGroup int, iItem int, iColumn int, pValue string) bool {
	return xc.XAdListView_Item_SetText(a.W句柄, iGroup, iItem, iColumn, pValue)
}

// 数据适配器列表视_项置文本扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterListView) Item_SetTextEx(iGroup int, iItem int, pName string, pValue string) bool {
	return xc.XAdListView_Item_SetTextEx(a.W句柄, iGroup, iItem, pName, pValue)
}

// 数据适配器列表视_项置图片, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) Item_SetImage(iGroup int, iItem int, iColumn int, hImage int) bool {
	return xc.XAdListView_Item_SetImage(a.W句柄, iGroup, iItem, iColumn, hImage)
}

// 数据适配器列表视_项置图片扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterListView) Item_SetImageEx(iGroup int, iItem int, pName string, hImage int) bool {
	return xc.XAdListView_Item_SetImageEx(a.W句柄, iGroup, iItem, pName, hImage)
}

// 数据适配器列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Group_GetText(iGroup int, iColumn int) string {
	return xc.XAdListView_Group_GetText(a.W句柄, iGroup, iColumn)
}

// 数据适配器列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (a *AdapterListView) Group_GetTextEx(iGroup int, pName string) string {
	return xc.XAdListView_Group_GetTextEx(a.W句柄, iGroup, pName)
}

// 数据适配器列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Group_GetImage(iGroup int, iColumn int) int {
	return xc.XAdListView_Group_GetImage(a.W句柄, iGroup, iColumn)
}

// 数据适配器列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (a *AdapterListView) Group_GetImageEx(iGroup int, pName string) int {
	return xc.XAdListView_Group_GetImageEx(a.W句柄, iGroup, pName)
}

// 数据适配器列表视_项取文本. 项操作, 返回项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Item_GetText(iGroup int, iItem int, iColumn int) string {
	return xc.XAdListView_Item_GetText(a.W句柄, iGroup, iItem, iColumn)
}

// 数据适配器列表视_项取图片. 项操作, 返回项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterListView) Item_GetImage(iGroup int, iItem int, iColumn int) int {
	return xc.XAdListView_Item_GetImage(a.W句柄, iGroup, iItem, iColumn)
}
