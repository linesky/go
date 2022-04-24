package main

import ("fmt"
)
func main(){
	var group1=[6]int {-2,-1,0,1,2,3};
	var group2=[]int{};
	var group3=[]int{}
	for n := range group1{
		if group1[n]/2*2==group1[n]{
			group2=append(group2,group1[n]);
		}else{
			group3=append(group3,group1[n]);
		}
	}
	fmt.Println("group :");
	fmt.Printf("%d\n",group1);
	fmt.Println("pare group :");
	fmt.Printf("%d\n",group2);
	fmt.Println("not pare group :");
	fmt.Printf("%d\n",group3);
}

