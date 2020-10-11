package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Module rsc.io/quote demo");
	fmt.Println(quote.Glass());
	fmt.Println(quote.Go());
	fmt.Println(quote.Hello());
	fmt.Println(quote.Opt());
}

