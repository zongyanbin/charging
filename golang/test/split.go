package main

import "strings"

// Split 将s按照sep进行分割，返回会一个字符串的切片
// Split("我爱你","爱") => ["我","你"]
// 1不容易观察 2留存
func Split(s, sep string)(ret []string){
	ret = make([]string,0,strings.Count(s,sep)+1) // 1 allocs/op  优化后
	 idx := strings.Index(s,sep)
	 for idx > -1{
	 	ret = append(ret,s[:idx])
	 	s = s[idx+len(sep):]
	 	idx = strings.Index(s,sep)
	 }
	 ret = append(ret, s)
	 return
}
