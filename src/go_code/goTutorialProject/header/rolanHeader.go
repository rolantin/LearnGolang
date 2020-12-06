package ro

import (
	"bufio"
	"fmt"
	"os"
)

func Print(t string) string {
	fmt.Println(t)
	return t
}

func PrintInput()  {
	s:= bufio.NewScanner(os.Stdin)
	s.Scan()
}
