package main

import "fmt"

var res [][]int
func combine(n int, k int) [][]int {
	res=[][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}

func backtrack(n,k,start int,track []int){
	if len(track)==k{
		temp:=make([]int,k)
		copy(temp,track)
		res=append(res,temp)
	}
	if len(track)+n-start+1 < k {
		return
	}
	for i:=start;i<=n;i++{
		track=append(track,i)
		backtrack(n,k,i+1,track)
	}
}

func main() {
	combine(4,2)
	fmt.Println(res)
}
