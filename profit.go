// Package profit calculates historical stock data to find the best
// profitable trade and trades
package profit

// MaxProfit takes a map of ints that represent
// stock values starting from a past opening time that
// is always 10:00am local time, it then traverses over the
// historical data to find the best buy and sell time and return the total profit
// Example Usage:
// 	profit.MaxProfit(Profitable(data))
func MaxProfit(i [][]int) int {
	var profit int
	for _, v := range i {
		if v[1]-v[0] > profit {
			profit = v[1] - v[0]
		}
	}

	return profit
}

// TotalMaxProfit returns the total maximum profit that
// could of been made if all profitable trades were executed
func TotalMaxProfit(i [][]int) int {
	var profit int
	for _, v := range i {
		profit = profit + (v[1] - v[0])
	}

	return profit
}

// ProfitableTrades calulates all the trades
// that could be profitable
// Example Usage:
// 	data := profit.Profitable(data)
func ProfitableTrades(i []int) [][]int {
	var p [][]int

	for k, v := range i {
		// We could do this in a go routine each loop
		// is done in a seperate thread to speed things up
		// but I think that would be a bit of over kill and less idomatic
		for k2, v2 := range i {
			if v < v2 && k < k2 {
				p = append(p, []int{v, v2})
			}
		}
	}

	return p
}
