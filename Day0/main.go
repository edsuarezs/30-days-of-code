package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		_ = fmt.Errorf(" %v error", err)
		return
	}
	fmt.Printf("Hello, World.\n%v\n", stdin)
}
