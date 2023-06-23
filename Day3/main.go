package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	lowerLimit  = 2
	upperLimit1 = 5
	upperLimit2 = 20
	Weird       = "Weird"
	NotWeird    = "NotWeird"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	checkNumber(N)
}

func checkNumber(N int32) {
	var answer string
	if N%2 == 0 {
		// Número par
		if N >= lowerLimit && N <= upperLimit1 {
			answer = NotWeird
		} else if N >= lowerLimit && N <= upperLimit2 {
			answer = Weird
		} else if N > upperLimit2 {
			answer = NotWeird
		}
	} else {
		answer = Weird
	}
	fmt.Println(answer)
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
