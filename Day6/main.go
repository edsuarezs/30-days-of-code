package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)
	Loop(n, reader)
}

func Loop(n int, reader *bufio.Reader) {
	for i := 0; i < int(n); i++ {
		splitSortString(strings.TrimSpace(readLine(reader)))
	}
}

func splitSortString(word string) {
	var splitedWord = strings.Split(word, "")
	var parSplitedWord []string
	var imparSplitedWord []string
	for i, v := range splitedWord {
		if i == 0 || i%2 == 0 {
			parSplitedWord = append(parSplitedWord, v)
		} else {
			imparSplitedWord = append(imparSplitedWord, v)
		}
	}
	fmt.Println(strings.Join(parSplitedWord, ""), strings.Join(imparSplitedWord, ""))

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
