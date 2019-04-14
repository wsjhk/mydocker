package main

/*
#include<stdio.h>
__attribute__((constructor)) void before_main() {
   printf("before main\n");
}

__attribute__((destructor)) void after_main() {
   printf("after main\n");
}
*/
import "C"

import "log"
func main()  {
	log.Printf("hello world!")
}