package main

import (
	"fmt"
	"go_logger/writer"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("start")
	logger := log.New(os.Stdout, "test", log.Ldate|log.Ltime|log.LUTC)
	fileWriter, _ := writer.NewFileWriter("test", "/mnt/d/Projects/ubuntu/go/go_logger/writer", writer.FileRotationPolicy{})
	fmt.Println("file open")
	n, err:= fileWriter.Write([]byte("sd\n"))
	if err !=nil{
		fmt.Println("error wirte", err)
	}
	fmt.Println("characters written", n)

	n, err= fileWriter.Write([]byte("sd\n"))
	if err !=nil{
		fmt.Println("error wirte", err)
	}
	fmt.Println("characters written", n)
	time.Sleep(50*time.Second)
	logger.Printf("hi")
}

// func main() {
// 	fmt.Printf("Start \n")

// 	//reader := bufio.NewReader(os.Stdin)
// 	_, err := logger.NewLogger(os.Stdout, "logger: %lda, %lti", true, 2)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
