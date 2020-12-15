package main

import (
	"fmt"
	"time"
)

func main(){
	const len = 1000
	start := time.Now()

	var a [len][len]int
	for i :=0;i<len;i++{
		for j:=0;j<len;j++{
			a[j][i]++
		}
	}
	fmt.Printf("Program runs for : %v\n",time.Since(start))
}
