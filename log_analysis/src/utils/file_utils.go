package utils

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func ReadFileLineToList(file_name string) (lines []string) {
	inputFile, inputError := os.Open(file_name)
	if inputError != nil {
		fmt.Println("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		//fmt.Println("the input was: %s", inputString)
		lines = append(lines, inputString)
		if readerError == io.EOF {
			return
		}
	}

}

func WriteObjToJsonFile(obj map[string]interface{}, file_name string) {

}
