package writer

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type FileRotationPolicy struct {
	MaxFileSize    int64
	MaxDaysPerFile int
}
type fileWriter struct {
	fileNamePrefix string
	logFilePath    string
	rotationPolicy FileRotationPolicy
	file           *os.File
	fileCreateDate time.Time
	mu             sync.Mutex
}

var logWriter fileWriter
var stopLogfileRotation = make(chan struct{})

func (fw *fileWriter) Write(p []byte) (n int, err error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	validateFileWithRotationPolicy(len(p))
	return fw.file.Write(p)
}

func validateFileWithRotationPolicy(lengthNewData int) {
	fileCreatingDays := time.Now().UTC().Sub(logWriter.fileCreateDate).Hours() / 24
	//fileCreatingDays := time.Now().UTC().Sub(logWriter.fileCreateDate).Seconds()
	fileStat, _ := logWriter.file.Stat()
	size := fileStat.Size() + int64(lengthNewData)
	if size >= logWriter.rotationPolicy.MaxFileSize || fileCreatingDays >= float64(logWriter.rotationPolicy.MaxDaysPerFile) {
		createTime := time.Now().UTC()
		file, fileOpenError := createFile(logWriter.fileNamePrefix, logWriter.logFilePath, createTime)

		if fileOpenError == nil {
			logWriter.file.Close()
			if fileStat.Size() == 0 {
				os.Remove(logWriter.file.Name())
			}
			logWriter.file = file
			logWriter.fileCreateDate = createTime
		}

	}
}
func fileRotation() {
	ticker := time.NewTicker(1 * time.Hour)

	for {
		select {
		case <-ticker.C:
			validateFileWithRotationPolicy(0)
		case <-stopLogfileRotation:
			ticker.Stop()
		}
	}
}

// createFile function will create the new log file
func createFile(prefix, path string, createTime time.Time) (*os.File, error) {
	timeNow := createTime.Unix()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(path, "does not exist")
		os.Mkdir(path, os.ModeDir|0775)
	}
	filePath := fmt.Sprintf("%s/%s_%d.txt", path, prefix, timeNow)
	file, fileOpenError := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if fileOpenError != nil {
		return nil, fileOpenError
	}

	return file, nil
}

func NewFileWriter(prefix, path string, rotationPolicy FileRotationPolicy) (Writer, error) {
	createTime := time.Now().UTC()

	file, fileOpenError := createFile(prefix, path, createTime)

	if fileOpenError != nil {
		return nil, fileOpenError
	}
	go fileRotation()
	logWriter = fileWriter{
		fileNamePrefix: prefix,
		logFilePath:    path,
		rotationPolicy: rotationPolicy,
		file:           file,
		fileCreateDate: createTime,
	}
	return &logWriter, nil
}
