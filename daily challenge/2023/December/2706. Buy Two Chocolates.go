package main

import "fmt"

func buyChoco(prices []int, money int) int {
	firstPrice := 100
	secondPrice := 100
	for i := 0; i < len(prices); i++ {
		if prices[i] < firstPrice {
			secondPrice = firstPrice
			firstPrice = prices[i]
		} else if prices[i] < secondPrice {
			secondPrice = prices[i]
		}
	}
	if firstPrice+secondPrice > money {
		return money
	}
	return money - firstPrice - secondPrice
}

func main() {
	prices1 := []int{1, 2, 2}
	money1 := 3

	prices2 := []int{3, 2, 3}
	money2 := 3
	fmt.Println("TEST CASE 1", buyChoco(prices1, money1) == 0)
	fmt.Println("TEST CASE 2", buyChoco(prices2, money2) == 3)

}
