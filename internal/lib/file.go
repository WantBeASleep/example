package lib

import (
	"fmt"
	"io"
	"os"
)

type Service interface {
	MakeFromFile(*os.File) File
}

func NewService() Service {
	return new(service)
}

type service struct {}
func (*service) MakeFromFile(f *os.File) File {
	_, err := f.Stat()
	fmt.Println(err)

	return new(file)
}

type Color string

const (
	Blue Color = "blue"
	Red Color = "red"
)

type File interface {
	FileColor(*os.File) (Color, error)
}

type file struct {}
func (*file) FileColor(file *os.File) (Color, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if len(content) > 10 {
		return Blue, nil
	}
	return Red, nil
}