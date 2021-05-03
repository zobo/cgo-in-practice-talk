package main

/*

#include <stdio.h>

void test1(char *text)
{
    printf("Native write: %s\n", text);
    free(text); // HL
}

*/
import "C"

func main() {
	C.test1(C.CString("Hello World!")) // HL
}
