package main

import ("fmt"
"sort"
)
func main(){
	group1:=[]int {1,20,18,17,13,12};
	fmt.Println("group :");
	fmt.Printf("%d\n",group1);
	sort.Ints(group1);
	fmt.Println("sort group :");
	fmt.Printf("%d\n",group1);
}

