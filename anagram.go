package main

import (
	"fmt"
)
func testIsAnagram() {
	s1 := "anagram"
	t1 := "nagaram"
	s2 := "rat"
	t2 := "car"
	isAnagram(s1, t1)
	isAnagram(s2, t2)
}

func isAnagram(s string, t string) (bool){
	record := make(map[int32]int, 26)
	for _, v := range s {
		// fmt.Println(reflect.TypeOf(v))
		// fmt.Println(v)         // range 字符串后得到int32类型的数字
		// fmt.Println(string(v)) // 可以将int32强制转成string
		record[v]++
	}
	for _, v := range t {
		record[v]--
	}
	for k, v := range record {
		if v != 0 {
			fmt.Printf("not anagram, character: %s \n", string(k))
			return false
		}
	}
	fmt.Printf("is anagram \n")
	return true
}

