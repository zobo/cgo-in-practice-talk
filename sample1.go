package main

/*

#include <stdio.h>

void test1(char *text)
{
    printf("Native write: %s\n", text);
}

*/
import "C"

func main() {
        C.test1("Hello World!")
}
