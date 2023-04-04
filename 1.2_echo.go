package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
