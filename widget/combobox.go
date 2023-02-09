package widget

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// 下拉组合框.
type ComboBox struct {
	Edit
}

// 组合框_创建.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func NewComboBox(x int, y int, cx int, cy int, hParent int) *ComboBox {
	p := &ComboBox{}
	p.W置句柄(xc.XComboBox_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewComboBoxByHandle(handle int) *ComboBox {
	p := &ComboBox{}
	p.W置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewComboBoxByName(name string) *ComboBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ComboBox{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewComboBoxByUID(nUID int) *ComboBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ComboBox{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewComboBoxByUIDName(name string) *ComboBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ComboBox{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 组合框_置选择项.
//
// iIndex: 项索引.
func (c *ComboBox) SetSelItem(iIndex int) bool {
	return xc.XComboBox_SetSelItem(c.W句柄, iIndex)
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
func (c *ComboBox) CreateAdapter() int {
	return xc.XComboBox_CreateAdapter(c.W句柄)
}

// 组合框_绑定数据适配器.
//
// hAdapter: 适配器句柄.
func (c *ComboBox) BindAdapter(hAdapter int) int {
	return xc.XComboBox_BindAdapter(c.W句柄, hAdapter)
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
func (c *ComboBox) GetAdapter() int {
	return xc.XComboBox_GetAdapter(c.W句柄)
}

// 组合框_置绑定名称.
//
// pName: 字段名.
func (c *ComboBox) SetBindName(pName string) int {
	return xc.XComboBox_SetBindName(c.W句柄, pName)
}

// 组合框_取下拉按钮坐标.
//
// pRect: 坐标.
func (c *ComboBox) GetButtonRect(pRect *xc.RECT) int {
	return xc.XComboBox_GetButtonRect(c.W句柄, pRect)
}

// 组合框_置下拉按钮大小.
//
// size: 大小.
func (c *ComboBox) SetButtonSize(size int) int {
	return xc.XComboBox_SetButtonSize(c.W句柄, size)
}

// 组合框_置下拉列表高度.
//
// height: 高度, -1自动计算高度.
func (c *ComboBox) SetDropHeight(height int) int {
	return xc.XComboBox_SetDropHeight(c.W句柄, height)
}

// 组合框_取下拉列表高度.
func (c *ComboBox) GetDropHeight() int {
	return xc.XComboBox_GetDropHeight(c.W句柄)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// pXmlFile: 项模板文件.
func (c *ComboBox) SetItemTemplateXML(pXmlFile string) int {
	return xc.XComboBox_SetItemTemplateXML(c.W句柄, pXmlFile)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// pStringXML: 字符串.
func (c *ComboBox) SetItemTemplateXMLFromString(pStringXML string) int {
	return xc.XComboBox_SetItemTemplateXMLFromString(c.W句柄, pStringXML)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// bEnable: 是否绘制.
func (c *ComboBox) EnableDrawButton(bEnable bool) int {
	return xc.XComboBox_EnableDrawButton(c.W句柄, bEnable)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// bEdit: TRUE可编辑.
func (c *ComboBox) EnableEdit(bEdit bool) int {
	return xc.XComboBox_EnableEdit(c.W句柄, bEdit)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// bEnable: 是否启用.
func (c *ComboBox) EnableDropHeightFixed(bEnable bool) int {
	return xc.XComboBox_EnableDropHeightFixed(c.W句柄, bEnable)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
func (c *ComboBox) GetSelItem() int {
	return xc.XComboBox_GetSelItem(c.W句柄)
}

// 组合框_取状态, 返回: ComboBox_State_.
func (c *ComboBox) GetState() xcc.ComboBox_State_ {
	return xc.XComboBox_GetState(c.W句柄)
}

// 组合框_添加项文本, 返回项索引.
//
// pText:.
func (c *ComboBox) AddItemText(pText string) int {
	return xc.XComboBox_AddItemText(c.W句柄, pText)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) AddItemTextEx(pName string, pText string) int {
	return xc.XComboBox_AddItemTextEx(c.W句柄, pName, pText)
}

// 组合框_添加项图片, 返回项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) AddItemImage(hImage int) int {
	return xc.XComboBox_AddItemImage(c.W句柄, hImage)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) AddItemImageEx(pName string, hImage int) int {
	return xc.XComboBox_AddItemImageEx(c.W句柄, pName, hImage)
}

// 组合框_插入项文本, 返回项索引.
//
// iItem: 项索引.
//
// pText: 文本.
func (c *ComboBox) InsertItemText(iItem int, pText string) int {
	return xc.XComboBox_InsertItemText(c.W句柄, iItem, pText)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) InsertItemTextEx(iItem int, pName string, pText string) int {
	return xc.XComboBox_InsertItemTextEx(c.W句柄, iItem, pName, pText)
}

// 组合框_插入项图片, 返回项索引.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) InsertItemImage(iItem int, hImage int) int {
	return xc.XComboBox_InsertItemImage(c.W句柄, iItem, hImage)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XComboBox_InsertItemImageEx(c.W句柄, iItem, pName, hImage)
}

// 组合框_置项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pText: 文本.
func (c *ComboBox) SetItemText(iItem int, iColumn int, pText string) bool {
	return xc.XComboBox_SetItemText(c.W句柄, iItem, iColumn, pText)
}

// 组合框_置项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) SetItemTextEx(iItem int, pName string, pText string) bool {
	return xc.XComboBox_SetItemTextEx(c.W句柄, iItem, pName, pText)
}

// 组合框_置项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (c *ComboBox) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XComboBox_SetItemImage(c.W句柄, iItem, iColumn, hImage)
}

// 组合框_置项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XComboBox_SetItemImageEx(c.W句柄, iItem, pName, hImage)
}

// 组合框_置项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func (c *ComboBox) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XComboBox_SetItemInt(c.W句柄, iItem, iColumn, nValue)
}

// 组合框_置项指数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// nValue: 整数值.
func (c *ComboBox) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XComboBox_SetItemIntEx(c.W句柄, iItem, pName, nValue)
}

// 组合框_置项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func (c *ComboBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {
	return xc.XComboBox_SetItemFloat(c.W句柄, iItem, iColumn, fFloat)
}

// 组合框_置项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// fFloat: 浮点数.
func (c *ComboBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {
	return xc.XComboBox_SetItemFloatEx(c.W句柄, iItem, pName, fFloat)
}

// 组合框_取项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) GetItemText(iItem int, iColumn int) string {
	return xc.XComboBox_GetItemText(c.W句柄, iItem, iColumn)
}

// 组合框_取项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) GetItemTextEx(iItem int, pName string) string {
	return xc.XComboBox_GetItemTextEx(c.W句柄, iItem, pName)
}

// 组合框_取项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) GetItemImage(iItem int, iColumn int) int {
	return xc.XComboBox_GetItemImage(c.W句柄, iItem, iColumn)
}

// 组合框_取项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) GetItemImageEx(iItem int, pName string) int {
	return xc.XComboBox_GetItemImageEx(c.W句柄, iItem, pName)
}

// 组合框_取项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) GetItemInt(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XComboBox_GetItemInt(c.W句柄, iItem, iColumn, pOutValue)
}

// 组合框_取项整数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) GetItemIntEx(iItem int, pName string, pOutValue *int) bool {
	return xc.XComboBox_GetItemIntEx(c.W句柄, iItem, pName, pOutValue)
}

// 组合框_取项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloat(c.W句柄, iItem, iColumn, pOutValue)
}

// 组合框_取项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloatEx(c.W句柄, iItem, pName, pOutValue)
}

// 组合框_删除项.
//
// iItem: 项索引.
func (c *ComboBox) DeleteItem(iItem int) bool {
	return xc.XComboBox_DeleteItem(c.W句柄, iItem)
}

// 组合框_删除项扩展.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func (c *ComboBox) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XComboBox_DeleteItemEx(c.W句柄, iItem, nCount)
}

// 组合框_删除项全部.
func (c *ComboBox) DeleteItemAll() int {
	return xc.XComboBox_DeleteItemAll(c.W句柄)
}

// 组合框_删除列全部.
func (c *ComboBox) DeleteColumnAll() int {
	return xc.XComboBox_DeleteColumnAll(c.W句柄)
}

// 组合框_取项数量.
func (c *ComboBox) GetCount() int {
	return xc.XComboBox_GetCount(c.W句柄)
}

// 组合框_取列数量.
func (c *ComboBox) GetCountColumn() int {
	return xc.XComboBox_GetCountColumn(c.W句柄)
}

// 组合框_弹出下拉列表.
func (c *ComboBox) PopupDropList() int {
	return xc.XComboBox_PopupDropList(c.W句柄)
}

// 组合框_设置项模板.
//
// hTemp: 模板句柄.
func (c *ComboBox) SetItemTemplate(hTemp int) int {
	return xc.XComboBox_SetItemTemplate(c.W句柄, hTemp)
}

/*
下面都是事件
*/

type XE_COMBOBOX_SELECT_END func(iItem int, pbHandled *bool) int                            // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT_END1 func(hEle int, iItem int, pbHandled *bool) int                 // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT func(iItem int, pbHandled *bool) int                                // 组合框下拉列表项选择事件.
type XE_COMBOBOX_SELECT1 func(hEle int, iItem int, pbHandled *bool) int                     // 组合框下拉列表项选择事件.
type XE_COMBOBOX_POPUP_LIST func(hWindow int, hListBox int, pbHandled *bool) int            // 组合框下拉列表弹出事件.
type XE_COMBOBOX_POPUP_LIST1 func(hEle int, hWindow int, hListBox int, pbHandled *bool) int // 组合框下拉列表弹出事件.
type XE_COMBOBOX_EXIT_LIST func(pbHandled *bool) int                                        // 组合框下拉列表退出事件.
type XE_COMBOBOX_EXIT_LIST1 func(hEle int, pbHandled *bool) int                             // 组合框下拉列表退出事件.

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) Event_ComboBox_Select_End(pFun XE_COMBOBOX_SELECT_END) bool {
	return xc.XEle_RegEventC(c.W句柄, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) Event_ComboBox_Select_End1(pFun XE_COMBOBOX_SELECT_END1) bool {
	return xc.XEle_RegEventC1(c.W句柄, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) Event_COMBOBOX_SELECT(pFun XE_COMBOBOX_SELECT) bool {
	return xc.XEle_RegEventC(c.W句柄, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) Event_COMBOBOX_SELECT1(pFun XE_COMBOBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(c.W句柄, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) Event_COMBOBOX_POPUP_LIST(pFun XE_COMBOBOX_POPUP_LIST) bool {
	return xc.XEle_RegEventC(c.W句柄, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) Event_COMBOBOX_POPUP_LIST1(pFun XE_COMBOBOX_POPUP_LIST1) bool {
	return xc.XEle_RegEventC1(c.W句柄, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) Event_COMBOBOX_EXIT_LIST(pFun XE_COMBOBOX_EXIT_LIST) bool {
	return xc.XEle_RegEventC(c.W句柄, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) Event_COMBOBOX_EXIT_LIST1(pFun XE_COMBOBOX_EXIT_LIST1) bool {
	return xc.XEle_RegEventC1(c.W句柄, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}
