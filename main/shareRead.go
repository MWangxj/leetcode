package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_open(char *name) {
    return shm_open(name, O_RDWR, FILE_MODE);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const SHM_NAME1 = "my_shm"
const SHM_SIZE1 = 4 * 1 << 30

type MyData1 struct {
	Col1 int
	Col2 int
	Col3 int
}

func main() {
	fd, err := C.my_shm_open(C.CString(SHM_NAME1))
	if err != nil {
		fmt.Println(err)
		return
	}

	ptr, err := C.mmap(nil, SHM_SIZE1, C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	C.close(fd)

	data := (*MyData1)(unsafe.Pointer(ptr))

	fmt.Println(data)
}
