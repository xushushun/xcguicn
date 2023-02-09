package 适配器

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// AdapterTable 数据适配器-XList-XListBox.
type AdapterTable struct {
	adapter
}

// 数据适配器表_创建, 创建列表框元素数据适配器.
func NewAdapterTable() *AdapterTable {
	p := &AdapterTable{}
	p.W置句柄(xc.XAdTable_Create())
	return p
}

// 从句柄创建对象.
func NewAdapterTableByHandle(handle int) *AdapterTable {
	p := &AdapterTable{}
	p.W置句柄(handle)
	return p
}

// 数据适配器表_排序, 对内容进行排序.
//
// iColumn: 要排序的列索引。.
//
// bAscending: 是否按照升序方式排序.
func (a *AdapterTable) Sort(iColumn int, bAscending bool) int {
	return xc.XAdTable_Sort(a.W句柄, iColumn, bAscending)
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemDataType(iItem int, iColumn int) xcc.Adapter_Date_Type_ {
	return xc.XAdTable_GetItemDataType(a.W句柄, iItem, iColumn)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemDataTypeEx(iItem int, pName string) xcc.Adapter_Date_Type_ {
	return xc.XAdTable_GetItemDataTypeEx(a.W句柄, iItem, pName)
}

// 数据适配器表_添加列, 添加数据列.
//
// pName: 字段称.
func (a *AdapterTable) AddColumn(pName string) int {
	return xc.XAdTable_AddColumn(a.W句柄, pName)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
// pColName: 列名, 多个列名用逗号分开.
func (a *AdapterTable) SetColumn(pColName string) int {
	return xc.XAdTable_SetColumn(a.W句柄, pColName)
}

// 数据适配器表_添加项文本, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// pValue: 值.
func (a *AdapterTable) AddItemText(pValue string) int {
	return xc.XAdTable_AddItemText(a.W句柄, pValue)
}

// 数据适配器表_添加项文本扩展, 添加数据项, 填充指定列数据, 返回项索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) AddItemTextEx(pName string, pValue string) int {
	return xc.XAdTable_AddItemTextEx(a.W句柄, pName, pValue)
}

// 数据适配器表_添加项图片, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hImage: 图片句柄.
func (a *AdapterTable) AddItemImage(hImage int) int {
	return xc.XAdTable_AddItemImage(a.W句柄, hImage)
}

// 数据适配器表_添加项图片扩展, 添加数据项, 并填充指定列数据.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) AddItemImageEx(pName string, hImage int) int {
	return xc.XAdTable_AddItemImageEx(a.W句柄, pName, hImage)
}

// 数据适配器表_插入项文本, 插入数据项, 填充第一列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pValue: 值.
func (a *AdapterTable) InsertItemText(iItem int, pValue string) int {
	return xc.XAdTable_InsertItemText(a.W句柄, iItem, pValue)
}

// 数据适配器表_插入项文本扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) InsertItemTextEx(iItem int, pName string, pValue string) int {
	return xc.XAdTable_InsertItemTextEx(a.W句柄, iItem, pName, pValue)
}

// 数据适配器表_插入项图片, 插入数据项, 填充第一列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// hImage: 图片句柄.
func (a *AdapterTable) InsertItemImage(iItem int, hImage int) int {
	return xc.XAdTable_InsertItemImage(a.W句柄, iItem, hImage)
}

// 数据适配器表_插入项图片扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XAdTable_InsertItemImageEx(a.W句柄, iItem, pName, hImage)
}

// 数据适配器表_置项文本, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterTable) SetItemText(iItem int, iColumn int, pValue string) bool {
	return xc.XAdTable_SetItemText(a.W句柄, iItem, iColumn, pValue)
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTable) SetItemTextEx(iItem int, pName string, pValue string) bool {
	return xc.XAdTable_SetItemTextEx(a.W句柄, iItem, pName, pValue)
}

// 数据适配器表_置项整数值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 值.
func (a *AdapterTable) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XAdTable_SetItemInt(a.W句柄, iItem, iColumn, nValue)
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// nValue: 值.
func (a *AdapterTable) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XAdTable_SetItemIntEx(a.W句柄, iItem, pName, nValue)
}

// 数据适配器表_置项浮点值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fValue: 值.
func (a *AdapterTable) SetItemFloat(iItem int, iColumn int, fValue float32) bool {
	return xc.XAdTable_SetItemFloat(a.W句柄, iItem, iColumn, fValue)
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// fValue: 值.
func (a *AdapterTable) SetItemFloatEx(iItem int, pName string, fValue float32) bool {
	return xc.XAdTable_SetItemFloatEx(a.W句柄, iItem, pName, fValue)
}

// 数据适配器表_置项图片, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterTable) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XAdTable_SetItemImage(a.W句柄, iItem, iColumn, hImage)
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTable) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XAdTable_SetItemImageEx(a.W句柄, iItem, pName, hImage)
}

// 数据适配器表_删除项, 删除项.
//
// iItem: 项索引.
func (a *AdapterTable) DeleteItem(iItem int) bool {
	return xc.XAdTable_DeleteItem(a.W句柄, iItem)
}

// 数据适配器表_删除项扩展, 删除多个项.
//
// iItem: 项起始索引.
//
// nCount: 删除项数量.
func (a *AdapterTable) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XAdTable_DeleteItemEx(a.W句柄, iItem, nCount)
}

// 数据适配器表_删除项全部, 删除所有项.
func (a *AdapterTable) DeleteItemAll() int {
	return xc.XAdTable_DeleteItemAll(a.W句柄)
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.
func (a *AdapterTable) DeleteColumnAll() int {
	return xc.XAdTable_DeleteColumnAll(a.W句柄)
}

// 数据适配器表_取项数量, 获取项数量.
func (a *AdapterTable) GetCount() int {
	return xc.XAdTable_GetCount(a.W句柄)
}

// 数据适配器表_取列数量, 获取列数量.
func (a *AdapterTable) GetCountColumn() int {
	return xc.XAdTable_GetCountColumn(a.W句柄)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemText(iItem int, iColumn int) string {
	return xc.XAdTable_GetItemText(a.W句柄, iItem, iColumn)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemTextEx(iItem int, pName string) string {
	return xc.XAdTable_GetItemTextEx(a.W句柄, iItem, pName)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterTable) GetItemImage(iItem int, iColumn int) int {
	return xc.XAdTable_GetItemImage(a.W句柄, iItem, iColumn)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterTable) GetItemImageEx(iItem int, pName string) int {
	return xc.XAdTable_GetItemImageEx(a.W句柄, iItem, pName)
}

// 数据适配器表_取项整数值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemInt(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XAdTable_GetItemInt(a.W句柄, iItem, iColumn, pOutValue)
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemIntEx(iItem int, pName string, pOutValue *int) bool {
	return xc.XAdTable_GetItemIntEx(a.W句柄, iItem, pName, pOutValue)
}

// 数据适配器表_取项浮点值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XAdTable_GetItemFloat(a.W句柄, iItem, iColumn, pOutValue)
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.
func (a *AdapterTable) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XAdTable_GetItemFloatEx(a.W句柄, iItem, pName, pOutValue)
}
