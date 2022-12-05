package main

import "fmt"
import "strconv"

func createAdjacencyList() [][]string {
	adjacencyLis := make([][]int, 2)
	adjacencyLis[0] = []int{1,2,3,4}
	adjacencyLis[1] = []int{5,6,7,8}

	strSum := make([][]string, 0)
    for _, v := range adjacencyLis {
        // intSlの値を文字列にしてstrSlに突っ込む
		strSl := make([]string, 0)
		for _, X := range v {
			strSl = append(strSl, strconv.Itoa(X))
		}
		strSum = append(strSum, strSl)
        
    }
	fmt.Println(strSum)

	return strSum

}