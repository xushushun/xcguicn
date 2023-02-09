package widget

import (
	"github.com/xushushun/xcguicn/xc"
	"github.com/xushushun/xcguicn/xcc"
)

// 列表树元素.
type Tree struct {
	ScrollView
}

// 列表树_创建, 创建树元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func NewTree(x int, y int, cx int, cy int, hParent int) *Tree {
	p := &Tree{}
	p.W置句柄(xc.XTree_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewTreeByHandle(handle int) *Tree {
	p := &Tree{}
	p.W置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewTreeByName(name string) *Tree {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Tree{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewTreeByUID(nUID int) *Tree {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Tree{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewTreeByUIDName(name string) *Tree {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Tree{}
		p.W置句柄(handle)
		return p
	}
	return nil
}

// 列表树_启用拖动项, 启用拖动项功能.
//
// bEnable: 是否启用.
func (t *Tree) EnableDragItem(bEnable bool) int {
	return xc.XTree_EnableDragItem(t.W句柄, bEnable)
}

// 列表树_启用连接线, 启用或禁用显示项的连接线.
//
// bEnable: 是否启用.
//
// bSolid: 实线或虚线; TRUE: 实线, FALSE: 虚线.
func (t *Tree) EnableConnectLine(bEnable bool, bSolid bool) int {
	return xc.XTree_EnableConnectLine(t.W句柄, bEnable, bSolid)
}

// 列表树_启用展开, 启动或关闭默认展开功能, 如果开启新插入的项将自动展开.
//
// bEnable: 是否启用.
func (t *Tree) EnableExpand(bEnable bool) int {
	return xc.XTree_EnableExpand(t.W句柄, bEnable)
}

// 列表树_启用模板复用.
//
// bEnable: 是否启用.
func (t *Tree) EnableTemplateReuse(bEnable bool) int {
	return xc.XTree_EnableTemplateReuse(t.W句柄, bEnable)
}

// 列表树_置连接线颜色.
//
// color: ABGR 颜色.
func (t *Tree) SetConnectLineColor(color int) int {
	return xc.XTree_SetConnectLineColor(t.W句柄, color)
}

// 列表树_置展开按钮大小, 设置展开按钮占用空间大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (t *Tree) SetExpandButtonSize(nWidth int, nHeight int) int {
	return xc.XTree_SetExpandButtonSize(t.W句柄, nWidth, nHeight)
}

// 列表树_置连接线长度, 设置连线绘制长度, 展开按钮与项内容之间的连线.
//
// nLength: 连线绘制长度.
func (t *Tree) SetConnectLineLength(nLength int) int {
	return xc.XTree_SetConnectLineLength(t.W句柄, nLength)
}

// 列表树_置拖动项插入位置颜色, 设置拖动项插入位置颜色提示.
//
// color: ABGR 颜色.
func (t *Tree) SetDragInsertPositionColor(color int) int {
	return xc.XTree_SetDragInsertPositionColor(t.W句柄, color)
}

// 列表树_置项模板文件.
//
// pXmlFile: 文件名.
func (t *Tree) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XTree_SetItemTemplateXML(t.W句柄, pXmlFile)
}

// 列表树_置选择项模板文件, 设置项模板文件, 项选中状态.
//
// pXmlFile: 文件名.
func (t *Tree) SetItemTemplateXMLSel(pXmlFile string) bool {
	return xc.XTree_SetItemTemplateXMLSel(t.W句柄, pXmlFile)
}

// 列表树_置项模板.
//
// hTemp: 模板句柄.
func (t *Tree) SetItemTemplate(hTemp int) bool {
	return xc.XTree_SetItemTemplate(t.W句柄, hTemp)
}

// 列表树_置选择项模板, 设置列表项模板, 项选中状态.
//
// hTemp: 模板句柄.
func (t *Tree) SetItemTemplateSel(hTemp int) bool {
	return xc.XTree_SetItemTemplateSel(t.W句柄, hTemp)
}

// 列表树_置项模板从字符串, 设置项模板文件.
//
// pStringXML: 字符串.
func (t *Tree) SetItemTemplateXMLFromString(pStringXML string) bool {
	return xc.XTree_SetItemTemplateXMLFromString(t.W句柄, pStringXML)
}

// 列表树_置选择项模板从字符串, 设置项模板文件, 项选中状态.
//
// pStringXML: 字符串.
func (t *Tree) SetItemTemplateXMLSelFromString(pStringXML string) bool {
	return xc.XTree_SetItemTemplateXMLSelFromString(t.W句柄, pStringXML)
}

// 列表树_置项背景绘制标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func (t *Tree) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XTree_SetDrawItemBkFlags(t.W句柄, nFlags)
}

// 列表树_置项数据, 设置项用户数据.
//
// nID: 项ID.
//
// nUserData: 用户数据.
func (t *Tree) SetItemData(nID int, nUserData int) bool {
	return xc.XTree_SetItemData(t.W句柄, nID, nUserData)
}

// 列表树_取项数据, 获取项用户数据.
//
// nID: 项ID.
func (t *Tree) GetItemData(nID int) int {
	return xc.XTree_GetItemData(t.W句柄, nID)
}

// 列表树_置选择项.
//
// nID: 项ID.
func (t *Tree) SetSelectItem(nID int) bool {
	return xc.XTree_SetSelectItem(t.W句柄, nID)
}

// 列表树_取选择项, 返回项ID.
func (t *Tree) GetSelectItem() int {
	return xc.XTree_GetSelectItem(t.W句柄)
}

// 列表树_可视指定项, 滚动视图让指定项可见.
//
// nID: 项索引.
func (t *Tree) VisibleItem(nID int) int {
	return xc.XTree_VisibleItem(t.W句柄, nID)
}

// 列表树_判断展开.
//
// nID: 项ID.
func (t *Tree) IsExpand(nID int) bool {
	return xc.XTree_IsExpand(t.W句柄, nID)
}

// 列表树_展开项, 判断项是否展开.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) ExpandItem(nID int, bExpand bool) bool {
	return xc.XTree_ExpandItem(t.W句柄, nID, bExpand)
}

// 列表树_展开全部子项, 展开所有的子项.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) ExpandAllChildItem(nID int, bExpand bool) bool {
	return xc.XTree_ExpandAllChildItem(t.W句柄, nID, bExpand)
}

// 列表树_测试点击项, 检测坐标点所在项, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) HitTest(pPt *xc.POINT) int {
	return xc.XTree_HitTest(t.W句柄, pPt)
}

// 列表树_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移坐标, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) HitTestOffset(pPt *xc.POINT) int {
	return xc.XTree_HitTestOffset(t.W句柄, pPt)
}

// 列表树_取第一个子项, 获取第一个子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) GetFirstChildItem(nID int) int {
	return xc.XTree_GetFirstChildItem(t.W句柄, nID)
}

// 列表树_取末尾子项, 获取末尾子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) GetEndChildItem(nID int) int {
	return xc.XTree_GetEndChildItem(t.W句柄, nID)
}

// 列表树_取上一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) GetPrevSiblingItem(nID int) int {
	return xc.XTree_GetPrevSiblingItem(t.W句柄, nID)
}

// 列表树_取下一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) GetNextSiblingItem(nID int) int {
	return xc.XTree_GetNextSiblingItem(t.W句柄, nID)
}

// 列表树_取父项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) GetParentItem(nID int) int {
	return xc.XTree_GetParentItem(t.W句柄, nID)
}

// 列表树_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (t *Tree) CreateAdapter() int {
	return xc.XTree_CreateAdapter(t.W句柄)
}

// 列表树_绑定数据适配器.
//
// hAdapter: 数据适配器句柄, XAdTree.
func (t *Tree) BindAdapter(hAdapter int) int {
	return xc.XTree_BindAdapter(t.W句柄, hAdapter)
}

// 列表树_取数据适配器, 返回数据适配器句柄.
func (t *Tree) GetAdapter() int {
	return xc.XTree_GetAdapter(t.W句柄)
}

// 列表树_刷新数据, 刷新所有项模板, 以便更新UI.
func (t *Tree) RefreshData() int {
	return xc.XTree_RefreshData(t.W句柄)
}

// 列表树_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// nID: 项ID.
func (t *Tree) RefreshItem(nID int) int {
	return xc.XTree_RefreshItem(t.W句柄, nID)
}

// 列表树_置缩进, 设置缩进大小.
//
// nWidth: 缩进宽度.
func (t *Tree) SetIndentation(nWidth int) int {
	return xc.XTree_SetIndentation(t.W句柄, nWidth)
}

// 列表树_取缩进, 返回缩进值大小.
func (t *Tree) GetIndentation() int {
	return xc.XTree_GetIndentation(t.W句柄)
}

// 列表树_置项默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) SetItemHeightDefault(nHeight int, nSelHeight int) int {
	return xc.XTree_SetItemHeightDefault(t.W句柄, nHeight, nSelHeight)
}

// 列表树_取项默认高度.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) GetItemHeightDefault(pHeight *int, pSelHeight *int) int {
	return xc.XTree_GetItemHeightDefault(t.W句柄, pHeight, pSelHeight)
}

// 列表树_置项高度.
//
// nID: 项ID.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) SetItemHeight(nID int, nHeight int, nSelHeight int) int {
	return xc.XTree_SetItemHeight(t.W句柄, nID, nHeight, nSelHeight)
}

// 列表树_取项高度.
//
// nID: 项ID.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) GetItemHeight(nID int, pHeight *int, pSelHeight *int) int {
	return xc.XTree_GetItemHeight(t.W句柄, nID, pHeight, pSelHeight)
}

// 列表树_置行间距.
//
// nSpace: 行间隔大小.
func (t *Tree) SetRowSpace(nSpace int) int {
	return xc.XTree_SetRowSpace(t.W句柄, nSpace)
}

// 列表树_取行间距.
func (t *Tree) GetRowSpace() int {
	return xc.XTree_GetRowSpace(t.W句柄)
}

// 列表树_移动项, 移动项的位置.
//
// nMoveItem: 要移动的项ID.
//
// nDestItem: 目标项ID, 参照位置.
//
// nFlag: 0:目标前面, 1:目标后面, 2:目标子项首, 3:目标子项尾.
func (t *Tree) MoveItem(nMoveItem int, nDestItem int, nFlag int) bool {
	return xc.XTree_MoveItem(t.W句柄, nMoveItem, nDestItem, nFlag)
}

// 列表树_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// nID: 树项ID.
//
// nTempItemID: 模板项ID.
func (t *Tree) GetTemplateObject(nID int, nTempItemID int) int {
	return xc.XTree_GetTemplateObject(t.W句柄, nID, nTempItemID)
}

// 列表树_取对象所在项, 获取当前对象所在模板实例, 属于列表树中哪一个项. 成功返回项ID, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (t *Tree) GetItemIDFromHXCGUI(hXCGUI int) int {
	return xc.XTree_GetItemIDFromHXCGUI(t.W句柄, hXCGUI)
}

// 列表树_插入项文本.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) InsertItemText(pValue string, nParentID int, insertID int) int {
	return xc.XTree_InsertItemText(t.W句柄, pValue, nParentID, insertID)
}

// 列表树_插入项文本扩展.
//
// pName:.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) InsertItemTextEx(pName string, pValue string, nParentID int, insertID int) int {
	return xc.XTree_InsertItemTextEx(t.W句柄, pName, pValue, nParentID, insertID)
}

// 列表树_插入项图片.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) InsertItemImage(hImage int, nParentID int, insertID int) int {
	return xc.XTree_InsertItemImage(t.W句柄, hImage, nParentID, insertID)
}

// 列表树_插入项图片扩展.
//
// pName:.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) InsertItemImageEx(pName string, hImage int, nParentID int, insertID int) int {
	return xc.XTree_InsertItemImageEx(t.W句柄, pName, hImage, nParentID, insertID)
}

// 列表树_取项数量.
func (t *Tree) GetCount() int {
	return xc.XTree_GetCount(t.W句柄)
}

// 列表树_取列数量.
func (t *Tree) GetCountColumn() int {
	return xc.XTree_GetCountColumn(t.W句柄)
}

// 列表树_置项文本.
//
// nID:.
//
// iColumn:.
//
// pValue:.
func (t *Tree) SetItemText(nID int, iColumn int, pValue string) bool {
	return xc.XTree_SetItemText(t.W句柄, nID, iColumn, pValue)
}

// 列表树_置项文本扩展.
//
// nID:.
//
// pName:.
//
// pValue:.
func (t *Tree) SetItemTextEx(nID int, pName string, pValue string) bool {
	return xc.XTree_SetItemTextEx(t.W句柄, nID, pName, pValue)
}

// 列表树_置项图片.
//
// nID:.
//
// iColumn:.
//
// hImage:.
func (t *Tree) SetItemImage(nID int, iColumn int, hImage int) bool {
	return xc.XTree_SetItemImage(t.W句柄, nID, iColumn, hImage)
}

// 列表树_置项图片扩展.
//
// nID:.
//
// pName:.
//
// hImage:.
func (t *Tree) SetItemImageEx(nID int, pName string, hImage int) bool {
	return xc.XTree_SetItemImageEx(t.W句柄, nID, pName, hImage)
}

// 列表树_取项文本.
//
// nID:.
//
// iColumn:.
func (t *Tree) GetItemText(nID int, iColumn int) string {
	return xc.XTree_GetItemText(t.W句柄, nID, iColumn)
}

// 列表树_取项文本扩展.
//
// nID:.
//
// pName:.
func (t *Tree) GetItemTextEx(nID int, pName string) string {
	return xc.XTree_GetItemTextEx(t.W句柄, nID, pName)
}

// 列表树_取项图片.
//
// nID:.
//
// iColumn:.
func (t *Tree) GetItemImage(nID int, iColumn int) int {
	return xc.XTree_GetItemImage(t.W句柄, nID, iColumn)
}

// 列表树_取项图片扩展.
//
// nID:.
//
// pName:.
func (t *Tree) GetItemImageEx(nID int, pName string) int {
	return xc.XTree_GetItemImageEx(t.W句柄, nID, pName)
}

// 列表树_删除项.
//
// nID:.
func (t *Tree) DeleteItem(nID int) bool {
	return xc.XTree_DeleteItem(t.W句柄, nID)
}

// 列表树_删除全部项.
func (t *Tree) DeleteItemAll() int {
	return xc.XTree_DeleteItemAll(t.W句柄)
}

// 列表树_删除列全部.
func (t *Tree) DeleteColumnAll() int {
	return xc.XTree_DeleteColumnAll(t.W句柄)
}

// 列表树_置分割线颜色.
//
// color: ABGR 颜色值.
func (t *Tree) SetSplitLineColor(color int) int {
	return xc.XTree_SetSplitLineColor(t.W句柄, color)
}

/*
以下都是事件
*/

type XE_TREE_TEMP_CREATE func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int                // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_TREE_TEMP_CREATE1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int     // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_TREE_TEMP_CREATE_END func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int            // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_TREE_TEMP_CREATE_END1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_TREE_TEMP_DESTROY func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int               // 列表树元素-项模板销毁,模板复用机制需先启用;.
type XE_TREE_TEMP_DESTROY1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int    // 列表树元素-项模板销毁,模板复用机制需先启用;.
type XE_TREE_TEMP_ADJUST_COORDINATE func(pItem *xc.Tree_Item_, pbHandled *bool) int                // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.Tree_Item_, pbHandled *bool) int     // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_DRAWITEM func(hDraw int, pItem *xc.Tree_Item_, pbHandled *bool) int                   // 树元素,绘制项.
type XE_TREE_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.Tree_Item_, pbHandled *bool) int        // 树元素,绘制项.
type XE_TREE_SELECT func(nItemID int, pbHandled *bool) int                                         // 树元素,项选择事件.
type XE_TREE_SELECT1 func(hEle int, nItemID int, pbHandled *bool) int                              // 树元素,项选择事件.
type XE_TREE_EXPAND func(id int, bExpand bool, pbHandled *bool) int                                // 树元素,项展开收缩事件.
type XE_TREE_EXPAND1 func(hEle int, id int, bExpand bool, pbHandled *bool) int                     // 树元素,项展开收缩事件.
type XE_TREE_DRAG_ITEM_ING func(pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int                    // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM_ING1 func(hEle int, pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int         // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM func(pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int                        // 树元素,拖动项事件.
type XE_TREE_DRAG_ITEM1 func(hEle int, pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int             // 树元素,拖动项事件.

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (t *Tree) Event_TREE_TEMP_CREATE(pFun XE_TREE_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (t *Tree) Event_TREE_TEMP_CREATE1(pFun XE_TREE_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (t *Tree) Event_TREE_TEMP_CREATE_END(pFun XE_TREE_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (t *Tree) Event_TREE_TEMP_CREATE_END1(pFun XE_TREE_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-项模板销毁,模板复用机制需先启用;.
func (t *Tree) Event_TREE_TEMP_DESTROY(pFun XE_TREE_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_TEMP_DESTROY, pFun)
}

// 列表树元素-项模板销毁,模板复用机制需先启用;.
func (t *Tree) Event_TREE_TEMP_DESTROY1(pFun XE_TREE_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_TEMP_DESTROY, pFun)
}

// 树元素,项模板,调整项坐标. 已停用.
func (t *Tree) Event_TREE_TEMP_ADJUST_COORDINATE(pFun XE_TREE_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,项模板,调整项坐标. 已停用.
func (t *Tree) Event_TREE_TEMP_ADJUST_COORDINATE1(pFun XE_TREE_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,绘制项.
func (t *Tree) Event_TREE_DRAWITEM(pFun XE_TREE_DRAWITEM) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_DRAWITEM, pFun)
}

// 树元素,绘制项.
func (t *Tree) Event_TREE_DRAWITEM1(pFun XE_TREE_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_DRAWITEM, pFun)
}

// 树元素,项选择事件.
func (t *Tree) Event_TREE_SELECT(pFun XE_TREE_SELECT) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_SELECT, pFun)
}

// 树元素,项选择事件.
func (t *Tree) Event_TREE_SELECT1(pFun XE_TREE_SELECT1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_SELECT, pFun)
}

// 树元素,项展开收缩事件.
func (t *Tree) Event_TREE_EXPAND(pFun XE_TREE_EXPAND) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_EXPAND, pFun)
}

// 树元素,项展开收缩事件.
func (t *Tree) Event_TREE_EXPAND1(pFun XE_TREE_EXPAND1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_EXPAND, pFun)
}

// 树元素,用户正在拖动项, 可对参数值修改.
func (t *Tree) Event_TREE_DRAG_ITEM_ING(pFun XE_TREE_DRAG_ITEM_ING) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,用户正在拖动项, 可对参数值修改.
func (t *Tree) Event_TREE_DRAG_ITEM_ING1(pFun XE_TREE_DRAG_ITEM_ING1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,拖动项事件.
func (t *Tree) Event_TREE_DRAG_ITEM(pFun XE_TREE_DRAG_ITEM) bool {
	return xc.XEle_RegEventC(t.W句柄, xcc.XE_TREE_DRAG_ITEM, pFun)
}

// 树元素,拖动项事件.
func (t *Tree) Event_TREE_DRAG_ITEM1(pFun XE_TREE_DRAG_ITEM1) bool {
	return xc.XEle_RegEventC1(t.W句柄, xcc.XE_TREE_DRAG_ITEM, pFun)
}
