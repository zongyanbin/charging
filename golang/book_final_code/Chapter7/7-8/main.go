package main

import "sync"

/**
延迟函数 主流程干完了，后面辅助流程需要干的事情
 */
var mu sync.Mutex
var m map[string]int

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func main()  {
	m:=make(map[string]int)
	m["abc"] = 123
	m["m"] = 3
	m["g"] = 80
	lookup("g")
}
