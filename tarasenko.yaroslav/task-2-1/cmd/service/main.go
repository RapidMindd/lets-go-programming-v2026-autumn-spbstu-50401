package main

import (
	"fmt"
	"os"
)

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for i := 0; i < N; i++ {
		var K int
		_, err := fmt.Scan(&K)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		var minTemp int = 15
		var maxTemp int = 30
		var clientComparator string
		var clientWish int
		for j := 0; j < K; j++ {
			_, err := fmt.Scan(&clientComparator, &clientWish)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			switch clientComparator {
			case ">=":
				minTemp = max(minTemp, clientWish)
			case "<=":
				maxTemp = min(maxTemp, clientWish)
			default:
				fmt.Fprintln(os.Stderr, "Incorrect comparator sign")
			}
			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
