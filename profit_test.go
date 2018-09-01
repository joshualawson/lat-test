package profit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestProfitable tests that the trade data
// returns a set of profitable trades
func TestProfitable(t *testing.T) {
	i := []int{
		10, 7, 5, 8, 11, 9,
	}

	profitableTrades := [][]int{
		[]int{10, 11},
		[]int{7, 8},
		[]int{7, 11},
		[]int{7, 9},
		[]int{5, 8},
		[]int{5, 11},
		[]int{5, 9},
		[]int{8, 11},
		[]int{8, 9},
	}

	assert.Equal(t, ProfitableTrades(i), profitableTrades, "The slices should match")
}

// BenchmarkProfitable1 benchmark tests for Profitable
func BenchmarkProfitable1(b *testing.B) {
	i := []int{
		10, 7, 5, 8, 11, 9,
	}

	ProfitableTrades(i)
}

// BenchmarkProfitable2 benchmark tests for Profitable with some more data
func BenchmarkProfitable2(b *testing.B) {
	i := []int{
		10, 7, 5, 8, 11, 9, 10, 7, 5, 8, 11, 9,
	}

	ProfitableTrades(i)
}

func TestMaxProfit(t *testing.T) {
	profitable := [][]int{
		[]int{10, 11},
		[]int{7, 8},
		[]int{7, 11},
		[]int{7, 9},
		[]int{5, 8},
		[]int{5, 11},
		[]int{5, 9},
		[]int{8, 11},
		[]int{8, 9},
	}

	assert.Equal(t, MaxProfit(profitable), 11-5, "Max profit should be 11-5")
}

// TestMaxProfitFromProfitable tests MaxProfit using the Profitable data
func TestMaxProfitFromProfitable(t *testing.T) {
	i := []int{
		10, 7, 5, 8, 11, 9,
	}

	assert.Equal(t, MaxProfit(ProfitableTrades(i)), 11-5, "Max profit should be 11-5")
}

// BenchmarkMaxProfit1 benchmark tests for MaxProfit
func BenchmarkMaxProfit1(b *testing.B) {
	i := [][]int{
		[]int{10, 11},
		[]int{7, 8},
		[]int{7, 11},
		[]int{7, 9},
		[]int{5, 8},
		[]int{5, 11},
		[]int{5, 9},
		[]int{8, 11},
		[]int{8, 9},
	}

	MaxProfit(i)
}

// BenchmarkMaxProfit1 benchmark tests for MaxProfit with more data
func BenchmarkMaxProfit2(b *testing.B) {
	i := [][]int{
		[]int{10, 11},
		[]int{7, 8},
		[]int{7, 11},
		[]int{7, 9},
		[]int{5, 8},
		[]int{5, 11},
		[]int{5, 9},
		[]int{8, 11},
		[]int{8, 9},
		[]int{10, 11},
		[]int{7, 8},
		[]int{7, 11},
		[]int{7, 9},
		[]int{5, 8},
		[]int{5, 11},
		[]int{5, 9},
		[]int{8, 11},
		[]int{8, 9},
	}

	MaxProfit(i)
}

// BenchmarkMaxProfitWithProfitable benchmark tests for MaxProfit using Profitable input
func BenchmarkMaxProfitWithProfitable(b *testing.B) {
	i := []int{
		10, 7, 5, 8, 11, 9,
	}

	MaxProfit(ProfitableTrades(i))
}
