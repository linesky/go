package main

import ("fmt"
)
func main(){
	var nums int32=0;
	var num2 float32=0.0;
	var group1=[6]float32 {-2.2,-1.0,0.0,1.1,2.0,3.1};
	var group2=[]float32{};
	var group3=[]float32{};
	for n := range group1{
		nums=int32(group1[n]);
		num2=float32(nums);
		if num2==group1[n]{
			group2=append(group2,group1[n]);
		}else{
			group3=append(group3,group1[n]);
		}
	}
	fmt.Println("group :");
	fmt.Printf("%f\n",group1);
	fmt.Println("integer group :");
	fmt.Printf("%f\n",group2);
	fmt.Println("not integer group :");
	fmt.Printf("%f\n",group3);
}

