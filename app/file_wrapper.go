package app

import (
	"fmt"
	"os"
)

type FileWrapper interface {
	Stat(filename string) (os.FileInfo, error)
	Open(filename string) (File, error)
}

type File interface {
	Close() error
	Read(p []byte) (n int, err error)
}

type FlyBehaviour interface {
	Fly()
}

type WingsFlyBehaviour struct{}

func (f WingsFlyBehaviour) Fly() {
	fmt.Println("I have some wings!")
}

type NoFlyBehaviour struct{}

func (f NoFlyBehaviour) Fly() {
	fmt.Println("I cannot fly!")
}

func run(flyBehaviour FlyBehaviour) {
	flyBehaviour.Fly()
}
