package main

import (
	"fmt"
	"flag"
	"os"
	"viscrypt/single"
	"viscrypt/multi"
)

func main() {

	n := flag.Int("n" , 2 , "number of transparencies, can't be used with multi-mode on") 
	isMulti := flag.Bool("multi-mode" , false , "wether to use multi-mode (default false)") 
	flag.Parse()
	
	if *isMulti == false {
		single.Encrypt(os.Args[3] , *n)
	} else { 
		fmt.Println(os.Args[3:])
		multi.Encrypt(os.Args[3:])
	}
}
