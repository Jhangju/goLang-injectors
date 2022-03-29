package main

import (
	"fmt"
	"go-inject/inject"
	"time"
)

func main() {

	time.Sleep(5 * time.Second)

	fmt.Println("[+] Detecting hooks.")
	originalSyscall := 3100740428

	moduleHandle := inject.GetModuleHandleA("ntdll.dll")
	fmt.Printf("[+] Got module handle ntdll.dll %d\n", moduleHandle)

	address := inject.GetProcAddress(moduleHandle, "NtAllocateVirtualMemory")
	fmt.Printf("[+] Got process address of NtAllocateVirtualMemory %d\n", address)

	tmpCheck := inject.RtlMoveMemory(address, 4)
	fmt.Println("[+] Moved memory from address.")

	if int(tmpCheck) == originalSyscall {
		fmt.Printf("Not hooked! %d\n", tmpCheck)
	} else {
		fmt.Printf("Hooked! %d\n", tmpCheck)
	}
}
