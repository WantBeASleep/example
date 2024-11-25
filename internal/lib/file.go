package lib

import (
	"io"
	"os"
)

type Service interface {
	MakeFromFile() File
}

func NewService() Service {
	return new(service)
}

type service struct {}
func (*service) MakeFromFile() File {
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