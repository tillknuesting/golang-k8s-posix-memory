package main

import (
	"github.com/fabiokung/shm"
	"log"
	"os"
)
func main(){
	file, err := shm.Open("/dev/shm", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	// syscall.Ftruncate if new, etc
	defer file.Close()
	defer shm.Unlink(file.Name())

	b := []byte("hello")

	_, err = file.Write(b)

	if err != nil {
		log.Fatal(err)
	}

}
