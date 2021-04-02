package types

// Money

type Money int64

// Category

type Category string

// Status представляет статуч плятежей.
type Status string

// Предопределеные статусы платежа.
const (
	StatusOk       Status = "OK"
	StatusFail     Status = "FAIL"
	StatusProgress Status = "INPROGRESS"
)

// Payment

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
