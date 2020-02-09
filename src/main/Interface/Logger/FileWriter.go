package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
)

/*
文件写入器是日志写入器的一种
他将日志写入到文件中
*/

type fileWriter struct {
	file *os.File
}

//创建一个空的文件写入器
func NewFileWriter() *fileWriter {
	return &fileWriter{}
}

func (f *fileWriter) Write(data interface{}) error {
	/*
		写入文件,如果文件未打开,则报错
	*/
	if f.file == nil {
		return errors.New("file not open")
	}

	//格式化文本
	str := fmt.Sprintf("%v\n", data)

	//写入文件中
	_, err := f.file.Write([]byte(str))
	return err

}

func (f *fileWriter) SetFile(filename string) (err error) {
	//设置文件写入位置
	env := os.Getenv("GOPATH")
	if f.file != nil {
		f.file.Close()
	}

	f.file, err = os.Create(path.Join(env, filename))
	return err
}

func (f *fileWriter) CloseFile() error {
	if f.file == nil {
		return errors.New("file not open")
	}
	return f.file.Close()
}
