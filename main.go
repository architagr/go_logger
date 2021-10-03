package main

import (
	"fmt"
	logger "go_logger/logger"
	"os"
)

func main() {
	fmt.Printf("Start \n")
	//reader := bufio.NewReader(os.Stdin)
	_, err := logger.NewLogger(os.Stdout, "logger: %lda, %lti", true, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
}
