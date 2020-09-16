package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	u := flag.String("u", "", "Upper")
	l := flag.String("l", "", "lower")
	flag.Parse()

	if *u == "" && *l == "" {
		fmt.Println("-u か -l を入れてね。")
		os.Exit(1)
	}

	if *u != "" {
		fmt.Println(strings.ToUpper(*u))
	}

	if *l != "" {
		fmt.Println(strings.ToLower(*l))
	}

	os.Exit(0)

}
