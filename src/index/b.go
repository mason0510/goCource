package main

import "fmt"

func main() {
	var countCapitalMap map[string]string /*创建集合*/
	countCapitalMap=make(map[string]string)

	//map插入key
	countCapitalMap["France"]="paris";
	countCapitalMap["Italy"]="罗马";

//迭代数组
	for country:=range countCapitalMap{
		fmt.Println(country,"首都是",countCapitalMap[country])
	}
	//检查元素是否存在
	capical,ok:=countCapitalMap["美国"]
	if (ok) {
		fmt.Println("美国首都是",capical)
	}else {
		fmt.Println("美国首都不存在")
	}

	//删除元素
	delete(countCapitalMap,"France")

	for country:=range countCapitalMap{
		fmt.Println(country,"首都是",countCapitalMap[country])
	}

}
