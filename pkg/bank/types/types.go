package types

// Money

type Money int64

// Category

type Category string

// Payment

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
