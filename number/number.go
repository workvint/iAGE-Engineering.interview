package main

import (
	"flag"
	"fmt"
	"github.com/workvint/numberutil"
	"strconv"
)

func main() {
	flag.Parse()

	if number, err := strconv.Atoi(flag.Arg(0)); err == nil {
		if result, err := numberutil.Verbal(number); err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("usage: $ number [valid integer]")
	}
}
