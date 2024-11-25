package trash

import (
	"fmt"
	"io"
	"os"
	"trash/internal/lib"
)

type service struct {
	colorSrv lib.Service
}

func NewService(colorSrv lib.Service) *service {
	return &service{
		colorSrv: colorSrv,
	}
}



func DoSmt(colorSrv lib.Service, content string) (lib.Color, error) {
	srv := NewService(colorSrv)

	file, err := os.CreateTemp("", "test-file-*.tmp")
	defer func(){
		fmt.Println("закрылся файл + удалился")
		file.Close()
		os.Remove(file.Name())
	}()
	if err != nil {
		return "", err
	}
	file.WriteString(content)
	file.Seek(0, io.SeekStart)

	fileSrv := srv.colorSrv.MakeFromFile()
	
	_, closeErr := file.Stat()
	fmt.Println("закрыт ли файл перед моком?", closeErr)
	
	color, err := fileSrv.FileColor(file)
	fmt.Println("FileColor do")
	if err != nil {
		return "", err
	}
	
	_, closeErr = file.Stat()
	fmt.Println("закрыт ли файл после мока?", closeErr)
	
	return color, err
}