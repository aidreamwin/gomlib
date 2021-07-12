// +build windows

package disk

import (
	"log"
	"syscall"
	"unsafe"
)

func getUsage() (int64, int64, int64, error) {
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		log.Panic(err)
		return 0, 0, 0, err
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	if err != nil {
		log.Panic(err)
		return 0, 0, 0, err
	}

	available := int64(0)
	free := int64(0)
	total := int64(0)
	res, err := syscall.UTF16PtrFromString("C:")
	if err != nil {
		log.Fatalln(err)
		return 0, 0, 0, err
	}
	r1, r2, err := syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(res)),
		uintptr(unsafe.Pointer(&available)),
		uintptr(unsafe.Pointer(&total)),
		uintptr(unsafe.Pointer(&free)), 0, 0)

	//if err != nil {
	//	log.Fatalln(err)
	//	return 0, 0, 0, err
	//}

	log.Println(r1, r2, err)
	return available, free, total, nil
}
