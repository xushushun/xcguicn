package wapi

import (
	"errors"
	"github.com/xushushun/xcguicn/common"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

// waitOpenClipboard 循环打开剪贴板，最多等待一秒钟.
//
//	@return bool
func waitOpenClipboard() bool {
	started := time.Now()
	limit := started.Add(time.Second)
	for time.Now().Before(limit) {
		if OpenClipboard(0) {
			return true
		}
		time.Sleep(time.Millisecond)
	}
	return false
}

// GetClipboardText 获取剪贴板中的文本.
//
//	@return string
//	@return error
func GetClipboardText() (string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 确定剪贴板是否包含指定格式的数据
	if !IsClipboardFormatAvailable(CF_UNICODETEXT) {
		return "", nil // 剪贴板中不包含Unicode文本数据
	}

	// 打开剪贴板
	if !waitOpenClipboard() {
		return "", errors.New("打开剪贴板失败")
	}
	defer CloseClipboard() // 关闭剪贴板

	// 取剪贴板数据句柄
	hMem := GetClipboardData(CF_UNICODETEXT)
	if hMem == 0 {
		return "", errors.New("未获取到剪贴板数据句柄")
	}

	// 锁定
	lpData := GlobalLock(hMem)
	if lpData == 0 {
		return "", errors.New("锁定剪贴板数据内存失败")
	}
	defer GlobalUnlock(hMem) // 解锁

	// 获取数据大小
	nSize := GlobalSize(hMem)
	if nSize == 0 {
		return "", errors.New("获取到剪贴板文本数据尺寸为0")
	}

	buf := make([]uint16, nSize)
	RtlMoveMemory(common.Uint16SliceDataPtr(&buf), lpData, nSize)
	return syscall.UTF16ToString(buf), nil
}

// SetClipboardText 将文本置入剪贴板.
//
//	@param text 要置入的文本.
//	@return error
func SetClipboardText(text string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 打开剪贴板
	if !waitOpenClipboard() {
		return errors.New("打开剪贴板失败")
	}
	defer CloseClipboard() // 关闭剪贴板

	// 清空剪贴板
	if !EmptyClipboard() {
		return errors.New("清空剪贴板失败")
	}

	// 转换
	bytes, _ := syscall.UTF16FromString(text)
	dwLength := len(bytes) * int(unsafe.Sizeof(bytes[0]))

	// 分配全局内存
	hGlobalMemory := GlobalAlloc(GMEM_Moveable, uint64(dwLength))
	if hGlobalMemory == 0 {
		return errors.New("分配全局内存失败")
	}
	defer GlobalFree(hGlobalMemory) // 释放全局内存

	// 锁住内存区
	lpGlobalMemory := GlobalLock(hGlobalMemory)
	if lpGlobalMemory == 0 {
		return errors.New("锁住内存区失败")
	}
	defer GlobalUnlock(hGlobalMemory) // 解锁内存区

	// 内存复制
	h := LstrcpyW(lpGlobalMemory, uintptr(unsafe.Pointer(&bytes[0])))
	if h == 0 {
		return errors.New("内存复制失败")
	}

	// 设置剪贴板数据
	if SetClipboardData(CF_UNICODETEXT, hGlobalMemory) == 0 {
		return errors.New("设置剪贴板数据失败")
	}

	return nil
}
