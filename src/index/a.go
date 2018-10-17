package main

import "fmt"

func main() {
	//这是我们使用和切片一样的数组来实现
	nums:=[]int{2,3,4}
	sum:=0
	for _num:=range nums{
		sum+=_num
	}
	fmt.Println("sum:",sum)

	//在数祖上使用range 将传入index和连个变量

	//求出索引
	for i,num:=range nums{
		if num==3 {
			fmt.Println("index:",i)
		}
	}

	//range也可以用在map的键值对上 迭代数组
	kvs:=map[string]string{"a":"apple","b":"banana"}
	for k,v:=range kvs{
		fmt.Println(k,v)
	}
}
