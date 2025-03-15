package strategy

type Payment interface {
	Pay(amount float32) string
}
