package main

import (
	"fmt"
	"sort"
)

func solve() {
	var n int 
	var k int 
	var ele int
	var arr[] int  
	var freq[] int

	fmt.Scanf("%d %d",&n,&k) 

	mp := make(map[int]int)
	
	for i:=0; i<n ; i++{
		fmt.Scanf("%d",&ele)
		arr = append(arr, ele)
	}
	
	for i:= range arr {
		mp[arr[i]]++ ;
	}
    
	for _ , v := range mp {
		freq = append(freq, v)
    }
    sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})
	if freq[0] < k {
		fmt.Println(n)
	} else {

	}
	fmt.Println(freq)
}

func main() {
	var testcase int 
	fmt.Scanf("%d",&testcase) 
	for i:=0;i<testcase;i++ {
		solve()
	}
}