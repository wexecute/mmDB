package mm_database

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type Requester struct {
	Data map[string]string
}

func (r *Requester) Load(name string) error {
	file, err := syscall.Open(name, syscall.O_CREAT | syscall.O_RDONLY | syscall.O_CLOEXEC, 1<<32-1)
	if err != nil {
		return err
	}

	stat, _ := os.Stat(name)
	size := stat.Size()

	buff := make([]byte, 1)
	for ; ; {
		syscall.Read(file, buff)
		if buff[0] == 0 {
			
		}
	}

	return nil
}

func (r *Requester) Save(name string) error {
	path, _ := syscall.UTF16PtrFromString(name)
	syscall.DeleteFile(path)

	file, err := syscall.Open(name, syscall.O_CREAT | syscall.O_WRONLY | syscall.O_CLOEXEC, 1<<32-1)
	if err != nil {
		return err
	}

	for k, v := range r.Data {
		syscall.Write(file, []byte(k + "\000" + v + "\000"))
	}

	return nil
}