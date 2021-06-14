/**
结构体标签在json中的应用
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string 		`json:"title"`
	Year int			`json:"year"`
	Price int			`json:"price"`
	Actors []string		`json:"actors"`
}

func main()  {
	// 定义结构体
	movie := Movie{"喜剧之王你",200,10,[]string{"xingye","zhangbozhi"}}

	// 编码的过程 结构体->json
	jsonStr, err := json.Marshal(movie)
	if err != nil{
		fmt.Println("json marshal error",err)
		return
	}

	fmt.Printf("jsonStr=%s\n", jsonStr)

	//	解码的过程 jsonStr ->结构体
	//jsonStr={"title":"喜剧之王你","year":200,"price":10,"actors":["xingye","zhangbozhi"]}
	my_Movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_Movie)
	if err !=nil{
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n",my_Movie)

}