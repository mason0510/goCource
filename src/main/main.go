package main

import "fmt"

//
//import (
//	"goLearn/src/cal"
//	"goLearn/src/cal/multi"
//	"fmt"
//)
//
//func main() {
//	result1 :=cal.Add(1,2)
//	result2:= cal.Subtract(110,2)
//	result3:=multi.Multi(2,4)
//	fmt.Printf("%d\n",result1)
//	fmt.Printf("%d\n",result2)
//	fmt.Printf("%d\n",result3)
//
//}

//行数
const LINES int = 10

// 杨辉三角
//func ShowYangHuiTriangle() {
//	nums := []int{}
//	for i := 0; i < LINES; i++ {
//		//补空白
//		for j := 0; j < (LINES - i); j++ {
//			fmt.Print(" ")
//		}
//		for j := 0; j < (i + 1); j++ {
//			var length = len(nums)
//			var value int
//			if j == 0 || j == i {
//				value = 1
//			} else {
//				value = nums[length-i] + nums[length-i-1]
//			}
//			nums = append(nums, value)
//			fmt.Print(value, " ")
//		}
//		fmt.Println("")
//	}
//}
//
//func main() {
//	ShowYangHuiTriangle()
//}


//func add(a, b int) int {
//	return a + b
//}
//
//func multiplicationTable() {
//	for i := 1; i <= 9; i++ {
//		for j := 1; j <= i; j++ {
//			var ret string
//			if i*j < 10 && j != 1 {
//				ret = " " + strconv.Itoa(i*j)
//			} else {
//				ret = strconv.Itoa(i * j)
//			}
//
//			fmt.Print(j, " * ", i, " = ", ret, "   ")
//		}
//		fmt.Print("\n")
//	}
//}
//
//func main() {
//	multiplicationTable()
//}


/* 声明全局变量 */
/* 声明全局变量 */
//var a int = 20;
//
//func main() {
//	/* main 函数中声明局部变量 */
//	var a int = 10
//	var b int = 20
//	var c int = 0
//
//	fmt.Printf("main()函数中 a = %d\n",  a);
//	c = sum( a, b);
//	fmt.Printf("main()函数中 c = %d\n",  c);
//}
//
///* 函数定义-两数相加 */
//func sum(a, b int) int {
//	fmt.Printf("sum() 函数中 a = %d\n",  a);
//	fmt.Printf("sum() 函数中 b = %d\n",  b);
//
//	return a + b;
//}
//var a int 20

//func main() {
//	//局部变量
//	var a  int=10
//	var b  int=20
//	var c int=0
//	fmt.Printf("main()函数",a)
//	c=sum(a,b)
//	println(a)
//	println(c)
//}
//func sum(a,b int)  int{
//	a=a+1
//	fmt.Printf("a=%d \n",a)
//	fmt.Printf("b=%d \n",b)
//	return a+b
//}
func main() {
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
//切片 函数 copy append 复制和拓展