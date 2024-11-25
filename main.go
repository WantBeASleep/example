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
		fmt.Println("А Я ВЫПОЛНИЛСЯ АХАХХАХАХ")
		file.Close()
	}()
	if err != nil {
		return "", err
	}
	file.WriteString(content)
	file.Seek(0, io.SeekStart)

	_, closeErr := file.Stat()
	fmt.Println("ЗАКРЫТ ЛИ ФАЙЛ?:", closeErr)
	
	fileSrv := srv.colorSrv.MakeFromFile(file)
	fmt.Println("MAKE_FROM_FILE!")
	_, closeErr = file.Stat()
	fmt.Println("ЗАКРЫТ ЛИ ФАЙЛ?VV2:", closeErr)
	
	color, err := fileSrv.FileColor(file)
	if err != nil {
		return "", err
	}

	return color, err
}