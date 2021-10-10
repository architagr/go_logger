package writer

import (
	"fmt"
	"os"
	"time"
)

type FileRotationPolicy struct {
	maxFileSize    int32
	maxDaysPerFile int
}
type fileWriter struct {
	fileNamePrefix string
	logFilePath    string
	rotationPolicy FileRotationPolicy
	file           *os.File
}

func (fw *fileWriter) Write(p []byte) (n int, err error) {
	return fw.file.Write(p)
}
func fileRotation(file *os.File){
    time.AfterFunc(20*time.Second, func() {
        x, _ := file.Stat()
        fmt.Println("filename:", x.Name(), "size:", x.Size())
    })
}
func NewFileWriter(prefix, path string, rotationPolicy FileRotationPolicy) (Writer, error) {
	timeNow := time.Now().UTC().Unix()
	mkdirErr:=os.Mkdir(path, os.ModeDir)
	if mkdirErr!=nil{
		fmt.Printf("mkdir error %+v\n", mkdirErr)
	}
	file, fileOpenError := os.OpenFile(fmt.Sprintf("%s/%s_%d.txt", path, prefix, timeNow), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if fileOpenError != nil {
		return nil, fileOpenError
	}
	go fileRotation(file)

	return &fileWriter{
		fileNamePrefix: prefix,
		logFilePath:    path,
		rotationPolicy: rotationPolicy,
		file:           file,
	}, nil
}
