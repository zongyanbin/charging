/**
数组 切片 属于单元素容器，它们属于连续存储，每个元素都是类型相同

字典map
不同他存储的不是单一元素
他存储的是键值对，索引列表  真正的字
所以键是一个类型， 所有的值是一个类型，键和值可以是不同类型
单独的键和单独的值必须是一个类型

字典的键 必须是可以通过 == 进行比较的数据类型
数组类型
字符串类型

模拟正删改查
*/
package main

import "fmt"

func main() {
var m map[string]int = map[string]int{"红烧肉":88, "清蒸鱼":98, "溜大虾":128, "蒸螃蟹":198, "鲍鱼粥":68}
fmt.Println(m)

foodsMap := make(map[string]int)
foodsMap["红烧肉"] = 88
foodsMap["清蒸鱼"] = 98
foodsMap["溜大虾"] = 128
foodsMap["蒸螃蟹"] = 198
foodsMap["鲍鱼粥"] = 68
fmt.Println(foodsMap)
}
