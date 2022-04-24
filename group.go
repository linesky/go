package main

import ("fmt"
"strings"
)
func main(){
	group1:=strings.Split("hello;world;hi;...",";");
	fmt.Println(group1);
}
