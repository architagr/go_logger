package main

import (
	"fmt"
	"go_logger/logger"
	"go_logger/writer"

	"time"
)

func main() {
	fmt.Println("start")

	fileWriter, writererr := writer.NewFileWriter("test", "/home/architagr/repo/golang/go_logger/writer/log", writer.FileRotationPolicy{
		MaxFileSize:    150,
		MaxDaysPerFile: 10,
	})

	if writererr != nil {
		fmt.Println(writererr)
		panic(writererr)
	}
	loggerW, err := logger.NewLogger(fileWriter, fmt.Sprintf("%s|%s:%s|%s|%s", logger.Llongfile, logger.Ldate, logger.Ltime, logger.LType, logger.Lmsg), true, 1)
	// logger := log.New(fileWriter, "test", log.Ldate|log.Ltime|log.LUTC)
	if err == nil {
		loggerW.Debug("sd1")

		loggerW.Debug("sd2")
		time.Sleep(2 * time.Second)
		loggerW.Debug("sd3")

		time.Sleep(50 * time.Second)
		loggerW.Debug("sd new file")
	}
	//logger.Printf("hi")
}

// func main() {
// 	fmt.Printf("Start \n")

// 	//reader := bufio.NewReader(os.Stdin)
// 	_, err := logger.NewLogger(os.Stdout, "logger: %lda, %lti", true, 2)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
