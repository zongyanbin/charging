package main

import "fmt"

func main() {
	Desc("老王", "划水")
}
func Desc(name string, fav string) {

}

func ShowFoods(foods []string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}
