package main

import (
	"flag"
	"os"
	"viscrypt/single"
)

func main() {

	n := flag.Int("n" , 2 , "number of transparencies") 
	multi := flag.Bool("multi-mode" , false , "wether to use multi mode") 
	
	if *multi == false {
		single.Encrypt(os.Args[3] , *n)
	}
}
