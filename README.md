# Profit

Profit calculates a set of historical data to find the most profitable trade for that specific day.

# Scenario

Suppose we could access yesterday's stock prices as a list, where:

* The indices are the time in minutes past trade opening time, which was 10:00am local time.
* The values are the price in dollars of the Latitude Financial stock at that time.
* So if the stock cost $5 at 11:00am, stock_prices_yesterday[60] = 5.

Write an *efficient function* that takes an array of stock prices and returns the best profit I could have made from 1 purchase and 1 sale of 1 Latitude Financial stock yesterday.

For example:
```js
var stock_prices_yesterday = [10, 7, 5, 8, 11, 9];

get_max_profit(stock_prices_yesterday)
# returns 6 (buying for $5 and selling for $11)
```

You must buy before you sell.
You may not buy and sell in the same time step (at least 1 minute must pass).

# Usage

```sh
$ go get github.com/joshualawson/profit
```


Basic Usage
```go
import (
    "fmt"
    "github.com/joshualawson/profit"
)

data := []int{
    10, 5, 2, 22, 8, 9, 12
}

profitableTrades := profit.Profitable(data)

maxProfit := profit.MaxProfit(profitableTrades)

fmt.Printf("Most profitable trade will have a profit of: %v\n", maxProfit)
```

## Dependencies

Dependency management is done using `dep` 

Dep can be found [here](https://github.com/golang/dep) with instruction on how to install `dep`

To load all of the packages dependencies run the following command from within the packages root dir

```sh
$ dep ensure
```

## Testing

To test the library simply run the following command

```sh
$ go test ./...
```

## Benchmark

To run benchmark results yourself use the following command

```sh
$ go test -bench=.
```

### Benchmark Results

Benchmark results using my MacBook Air

```
BenchmarkProfitable1-4                  2000000000               0.00 ns/op
BenchmarkProfitable2-4                  2000000000               0.00 ns/op
BenchmarkMaxProfit1-4                   2000000000               0.00 ns/op
BenchmarkMaxProfit2-4                   2000000000               0.00 ns/op
BenchmarkMaxProfitWithProfitable-4      2000000000               0.00 ns/op
```

## Coverage

Coverage is currently at 100% to see the results yourself run the following command

```sh
$ go test -cover
```

### Coverage Results

```
PASS
coverage: 100.0% of statements
```