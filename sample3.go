package main

/*
#include <stdio.h>
#include <stdlib.h> // HL

void test1(char *text)
{
    printf("Native write: %s\n", text);
}
*/
import "C"
import "unsafe"

func main() {
	s := C.CString("Hello World!")
	defer C.free(unsafe.Pointer(s)) // HL
	C.test1()
}
