package main

import "fmt"

func main()  {
	s:=[]int{168,68,98,1888}
	for _, item:=range s{
		if item==1888{
			fmt.Println("整单一折")
		}
	}
}
