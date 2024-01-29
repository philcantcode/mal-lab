package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modDbgHelp  = windows.NewLazySystemDLL("DbgHelp.dll")
	modGdi32    = windows.NewLazySystemDLL("Gdi32.dll")
	modKernel32 = windows.NewLazySystemDLL("Kernel32.dll")
	modUser32   = windows.NewLazySystemDLL("User32.dll")
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")
	modpsapi    = windows.NewLazySystemDLL("psapi.dll")

	procMiniDumpWriteDump                 = modDbgHelp.NewProc("MiniDumpWriteDump")
	procBitBlt                            = modGdi32.NewProc("BitBlt")
	procCreateCompatibleBitmap            = modGdi32.NewProc("CreateCompatibleBitmap")
	procCreateCompatibleDC                = modGdi32.NewProc("CreateCompatibleDC")
	procDeleteDC                          = modGdi32.NewProc("DeleteDC")
	procDeleteObject                      = modGdi32.NewProc("DeleteObject")
	procGetDIBits                         = modGdi32.NewProc("GetDIBits")
	procSelectObject                      = modGdi32.NewProc("SelectObject")
	procGlobalAlloc                       = modKernel32.NewProc("GlobalAlloc")
	procGlobalFree                        = modKernel32.NewProc("GlobalFree")
	procGlobalLock                        = modKernel32.NewProc("GlobalLock")
	procGlobalUnlock                      = modKernel32.NewProc("GlobalUnlock")
	procGetDC                             = modUser32.NewProc("GetDC")
	procGetDesktopWindow                  = modUser32.NewProc("GetDesktopWindow")
	procReleaseDC                         = modUser32.NewProc("ReleaseDC")
	procCreateProcessWithLogonW           = modadvapi32.NewProc("CreateProcessWithLogonW")
	procImpersonateLoggedOnUser           = modadvapi32.NewProc("ImpersonateLoggedOnUser")
	procLogonUserW                        = modadvapi32.NewProc("LogonUserW")
	procLookupPrivilegeDisplayNameW       = modadvapi32.NewProc("LookupPrivilegeDisplayNameW")
	procLookupPrivilegeNameW              = modadvapi32.NewProc("LookupPrivilegeNameW")
	procRegSaveKeyW                       = modadvapi32.NewProc("RegSaveKeyW")
	procCreateProcessW                    = modkernel32.NewProc("CreateProcessW")
	procCreateRemoteThread                = modkernel32.NewProc("CreateRemoteThread")
	procCreateThread                      = modkernel32.NewProc("CreateThread")
	procDeleteProcThreadAttributeList     = modkernel32.NewProc("DeleteProcThreadAttributeList")
	procGetExitCodeThread                 = modkernel32.NewProc("GetExitCodeThread")
	procGetProcessHeap                    = modkernel32.NewProc("GetProcessHeap")
	procHeapAlloc                         = modkernel32.NewProc("HeapAlloc")
	procHeapFree                          = modkernel32.NewProc("HeapFree")
	procHeapReAlloc                       = modkernel32.NewProc("HeapReAlloc")
	procHeapSize                          = modkernel32.NewProc("HeapSize")
	procInitializeProcThreadAttributeList = modkernel32.NewProc("InitializeProcThreadAttributeList")
	procModule32FirstW                    = modkernel32.NewProc("Module32FirstW")
	procPssCaptureSnapshot                = modkernel32.NewProc("PssCaptureSnapshot")
	procQueueUserAPC                      = modkernel32.NewProc("QueueUserAPC")
	procUpdateProcThreadAttribute         = modkernel32.NewProc("UpdateProcThreadAttribute")
	procVirtualAllocEx                    = modkernel32.NewProc("VirtualAllocEx")
	procVirtualProtectEx                  = modkernel32.NewProc("VirtualProtectEx")
	procWriteProcessMemory                = modkernel32.NewProc("WriteProcessMemory")
	procRtlCopyMemory                     = modntdll.NewProc("RtlCopyMemory")
	procGetProcessMemoryInfo              = modpsapi.NewProc("GetProcessMemoryInfo")
)

func MiniDumpWriteDump(hProcess windows.Handle, pid uint32, hFile uintptr, dumpType uint32, exceptionParam uintptr, userStreamParam uintptr, callbackParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procMiniDumpWriteDump.Addr(), 7, uintptr(hProcess), uintptr(pid), uintptr(hFile), uintptr(dumpType), uintptr(exceptionParam), uintptr(userStreamParam), uintptr(callbackParam), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func BitBlt(hdc windows.Handle, x uint32, y uint32, cx uint32, cy uint32, hdcSrc windows.Handle, x1 uint32, y1 uint32, rop int32) (BOOL int, err error) {
	r0, _, e1 := syscall.Syscall9(procBitBlt.Addr(), 9, uintptr(hdc), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hdcSrc), uintptr(x1), uintptr(y1), uintptr(rop))
	BOOL = int(r0)
	if BOOL == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateCompatibleBitmap(hdc windows.Handle, cx int, cy int) (HBITMAP windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleBitmap.Addr(), 3, uintptr(hdc), uintptr(cx), uintptr(cy))
	HBITMAP = windows.Handle(r0)
	if HBITMAP == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateCompatibleDC(hdc windows.Handle) (HDC windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleDC.Addr(), 1, uintptr(hdc), 0, 0)
	HDC = windows.Handle(r0)
	if HDC == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeleteDC(hdc windows.Handle) (BOOL uint32, err error) {
	r0, _, e1 := syscall.Syscall(procDeleteDC.Addr(), 1, uintptr(hdc), 0, 0)
	BOOL = uint32(r0)
	if BOOL == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeleteObject(ho windows.Handle) (BOOL uint32, err error) {
	r0, _, e1 := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(ho), 0, 0)
	BOOL = uint32(r0)
	if BOOL == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDIBits(hdc windows.Handle, hbm windows.Handle, start uint32, cLines uint32, lpvBits uintptr, lpbmi uintptr, usage int) (ret int, err error) {
	r0, _, e1 := syscall.Syscall9(procGetDIBits.Addr(), 7, uintptr(hdc), uintptr(hbm), uintptr(start), uintptr(cLines), uintptr(lpvBits), uintptr(lpbmi), uintptr(usage), 0, 0)
	ret = int(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func SelectObject(hdc windows.Handle, h windows.Handle) (HGDIOBJ windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procSelectObject.Addr(), 2, uintptr(hdc), uintptr(h), 0)
	HGDIOBJ = windows.Handle(r0)
	if HGDIOBJ == 0 {
		err = errnoErr(e1)
	}
	return
}

func GlobalAlloc(uFlags uint, dwBytes uintptr) (HGLOBAL windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGlobalAlloc.Addr(), 2, uintptr(uFlags), uintptr(dwBytes), 0)
	HGLOBAL = windows.Handle(r0)
	if HGLOBAL == 0 {
		err = errnoErr(e1)
	}
	return
}

func GlobalFree(hMem windows.Handle) (HGLOBAL windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGlobalFree.Addr(), 1, uintptr(hMem), 0, 0)
	HGLOBAL = windows.Handle(r0)
	if HGLOBAL == 0 {
		err = errnoErr(e1)
	}
	return
}

func GlobalLock(hMem windows.Handle) (LPVOID uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGlobalLock.Addr(), 1, uintptr(hMem), 0, 0)
	LPVOID = uintptr(r0)
	if LPVOID == 0 {
		err = errnoErr(e1)
	}
	return
}

func GlobalUnlock(hMem windows.Handle) (BOOL uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGlobalUnlock.Addr(), 1, uintptr(hMem), 0, 0)
	BOOL = uint32(r0)
	if BOOL == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDC(HWND windows.Handle) (HDC windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetDC.Addr(), 1, uintptr(HWND), 0, 0)
	HDC = windows.Handle(r0)
	if HDC == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDesktopWindow() (HWND windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetDesktopWindow.Addr(), 0, 0, 0, 0)
	HWND = windows.Handle(r0)
	if HWND == 0 {
		err = errnoErr(e1)
	}
	return
}

func ReleaseDC(hWnd windows.Handle, hDC windows.Handle) (int uint32, err error) {
	r0, _, e1 := syscall.Syscall(procReleaseDC.Addr(), 2, uintptr(hWnd), uintptr(hDC), 0)
	int = uint32(r0)
	if int == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateProcessWithLogonW(username *uint16, domain *uint16, password *uint16, logonFlags uint32, appName *uint16, commandLine *uint16, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfoEx, outProcInfo *windows.ProcessInformation) (err error) {
	r1, _, e1 := syscall.Syscall12(procCreateProcessWithLogonW.Addr(), 11, uintptr(unsafe.Pointer(username)), uintptr(unsafe.Pointer(domain)), uintptr(unsafe.Pointer(password)), uintptr(logonFlags), uintptr(unsafe.Pointer(appName)), uintptr(unsafe.Pointer(commandLine)), uintptr(creationFlags), uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(currentDir)), uintptr(unsafe.Pointer(startupInfo)), uintptr(unsafe.Pointer(outProcInfo)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ImpersonateLoggedOnUser(hToken windows.Token) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateLoggedOnUser.Addr(), 1, uintptr(hToken), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LogonUser(lpszUsername *uint16, lpszDomain *uint16, lpszPassword *uint16, dwLogonType uint32, dwLogonProvider uint32, phToken *windows.Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procLogonUserW.Addr(), 6, uintptr(unsafe.Pointer(lpszUsername)), uintptr(unsafe.Pointer(lpszDomain)), uintptr(unsafe.Pointer(lpszPassword)), uintptr(dwLogonType), uintptr(dwLogonProvider), uintptr(unsafe.Pointer(phToken)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LookupPrivilegeDisplayNameW(systemName string, privilegeName *uint16, buffer *uint16, size *uint32, languageId *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	return _LookupPrivilegeDisplayNameW(_p0, privilegeName, buffer, size, languageId)
}

func _LookupPrivilegeDisplayNameW(systemName *uint16, privilegeName *uint16, buffer *uint16, size *uint32, languageId *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeDisplayNameW.Addr(), 5, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(privilegeName)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(languageId)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LookupPrivilegeNameW(systemName string, luid *uint64, buffer *uint16, size *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	return _LookupPrivilegeNameW(_p0, luid, buffer, size)
}

func _LookupPrivilegeNameW(systemName *uint16, luid *uint64, buffer *uint16, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeNameW.Addr(), 4, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(luid)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RegSaveKeyW(hKey windows.Handle, lpFile *uint16, lpSecurityAttributes *windows.SecurityAttributes) (err error) {
	r1, _, e1 := syscall.Syscall(procRegSaveKeyW.Addr(), 3, uintptr(hKey), uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(lpSecurityAttributes)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateProcess(appName *uint16, commandLine *uint16, procSecurity *windows.SecurityAttributes, threadSecurity *windows.SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfoEx, outProcInfo *windows.ProcessInformation) (err error) {
	var _p0 uint32
	if inheritHandles {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall12(procCreateProcessW.Addr(), 10, uintptr(unsafe.Pointer(appName)), uintptr(unsafe.Pointer(commandLine)), uintptr(unsafe.Pointer(procSecurity)), uintptr(unsafe.Pointer(threadSecurity)), uintptr(_p0), uintptr(creationFlags), uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(currentDir)), uintptr(unsafe.Pointer(startupInfo)), uintptr(unsafe.Pointer(outProcInfo)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateRemoteThread(hProcess windows.Handle, lpThreadAttributes *windows.SecurityAttributes, dwStackSize uint32, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uint32) (threadHandle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateRemoteThread.Addr(), 7, uintptr(hProcess), uintptr(unsafe.Pointer(lpThreadAttributes)), uintptr(dwStackSize), uintptr(lpStartAddress), uintptr(lpParameter), uintptr(dwCreationFlags), uintptr(unsafe.Pointer(lpThreadId)), 0, 0)
	threadHandle = windows.Handle(r0)
	if threadHandle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateThread(lpThreadAttributes *windows.SecurityAttributes, dwStackSize uint32, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32, lpThreadId *uint32) (threadHandle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateThread.Addr(), 6, uintptr(unsafe.Pointer(lpThreadAttributes)), uintptr(dwStackSize), uintptr(lpStartAddress), uintptr(lpParameter), uintptr(dwCreationFlags), uintptr(unsafe.Pointer(lpThreadId)))
	threadHandle = windows.Handle(r0)
	if threadHandle == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeleteProcThreadAttributeList(lpAttributeList *PROC_THREAD_ATTRIBUTE_LIST) {
	syscall.Syscall(procDeleteProcThreadAttributeList.Addr(), 1, uintptr(unsafe.Pointer(lpAttributeList)), 0, 0)
	return
}

func GetExitCodeThread(hTread windows.Handle, lpExitCode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetExitCodeThread.Addr(), 2, uintptr(hTread), uintptr(unsafe.Pointer(lpExitCode)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetProcessHeap() (procHeap windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcessHeap.Addr(), 0, 0, 0, 0)
	procHeap = windows.Handle(r0)
	if procHeap == 0 {
		err = errnoErr(e1)
	}
	return
}

func HeapAlloc(hHeap windows.Handle, dwFlags uint32, dwBytes uintptr) (lpMem uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procHeapAlloc.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(dwBytes))
	lpMem = uintptr(r0)
	if lpMem == 0 {
		err = errnoErr(e1)
	}
	return
}

func HeapFree(hHeap windows.Handle, dwFlags uint32, lpMem uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procHeapFree.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(lpMem))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func HeapReAlloc(hHeap windows.Handle, dwFlags uint32, lpMem uintptr, dwBytes uintptr) (lpRes uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procHeapReAlloc.Addr(), 4, uintptr(hHeap), uintptr(dwFlags), uintptr(lpMem), uintptr(dwBytes), 0, 0)
	lpRes = uintptr(r0)
	if lpRes == 0 {
		err = errnoErr(e1)
	}
	return
}

func HeapSize(hHeap windows.Handle, dwFlags uint32, lpMem uintptr) (res uint32, err error) {
	r0, _, e1 := syscall.Syscall(procHeapSize.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(lpMem))
	res = uint32(r0)
	if res == 0 {
		err = errnoErr(e1)
	}
	return
}

func InitializeProcThreadAttributeList(lpAttributeList *PROC_THREAD_ATTRIBUTE_LIST, dwAttributeCount uint32, dwFlags uint32, lpSize *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procInitializeProcThreadAttributeList.Addr(), 4, uintptr(unsafe.Pointer(lpAttributeList)), uintptr(dwAttributeCount), uintptr(dwFlags), uintptr(unsafe.Pointer(lpSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Module32FirstW(hSnapshot windows.Handle, lpme *MODULEENTRY32W) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(hSnapshot), uintptr(unsafe.Pointer(lpme)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func PssCaptureSnapshot(processHandle windows.Handle, captureFlags uint32, threadContextFlags uint32, snapshotHandle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall6(procPssCaptureSnapshot.Addr(), 4, uintptr(processHandle), uintptr(captureFlags), uintptr(threadContextFlags), uintptr(unsafe.Pointer(snapshotHandle)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueueUserAPC(pfnAPC uintptr, hThread windows.Handle, dwData uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procQueueUserAPC.Addr(), 3, uintptr(pfnAPC), uintptr(hThread), uintptr(dwData))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func UpdateProcThreadAttribute(lpAttributeList *PROC_THREAD_ATTRIBUTE_LIST, dwFlags uint32, attribute uintptr, lpValue *uintptr, cbSize uintptr, lpPreviousValue uintptr, lpReturnSize *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procUpdateProcThreadAttribute.Addr(), 7, uintptr(unsafe.Pointer(lpAttributeList)), uintptr(dwFlags), uintptr(attribute), uintptr(unsafe.Pointer(lpValue)), uintptr(cbSize), uintptr(lpPreviousValue), uintptr(unsafe.Pointer(lpReturnSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualAllocEx(hProcess windows.Handle, lpAddress uintptr, dwSize uintptr, flAllocationType uint32, flProtect uint32) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procVirtualAllocEx.Addr(), 5, uintptr(hProcess), uintptr(lpAddress), uintptr(dwSize), uintptr(flAllocationType), uintptr(flProtect), 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualProtectEx(hProcess windows.Handle, lpAddress uintptr, dwSize uintptr, flNewProtect uint32, lpflOldProtect *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualProtectEx.Addr(), 5, uintptr(hProcess), uintptr(lpAddress), uintptr(dwSize), uintptr(flNewProtect), uintptr(unsafe.Pointer(lpflOldProtect)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func WriteProcessMemory(hProcess windows.Handle, lpBaseAddress uintptr, lpBuffer *byte, nSize uintptr, lpNumberOfBytesWritten *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procWriteProcessMemory.Addr(), 5, uintptr(hProcess), uintptr(lpBaseAddress), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize), uintptr(unsafe.Pointer(lpNumberOfBytesWritten)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RtlCopyMemory(dest uintptr, src uintptr, dwSize uint32) {
	syscall.Syscall(procRtlCopyMemory.Addr(), 3, uintptr(dest), uintptr(src), uintptr(dwSize))
	return
}

func GetProcessMemoryInfo(process windows.Handle, ppsmemCounters *ProcessMemoryCounters, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(process), uintptr(unsafe.Pointer(ppsmemCounters)), uintptr(cb))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
