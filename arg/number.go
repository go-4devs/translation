package arg

import "fmt"

// FormatNumber format number.
type FormatNumber int

// format argument.
const (
	NumberFormatDecimal FormatNumber = iota + 1
	NumberFormatPercent
	NumberFormatPerMille
	NumberFormatEngineering
	NumberFormatScientific
)

// NumberOption configure number argument.
type NumberOption func(*Number)

// WithNumberFormat sets format number.
func WithNumberFormat(format FormatNumber) NumberOption {
	return func(n *Number) { n.Format = format }
}

// Number argument.
type Number struct {
	Key    string
	Value  interface{}
	Format FormatNumber
}

// Configure number.
func (n Number) Configure(opts ...NumberOption) Number {
	for _, o := range opts {
		o(&n)
	}

	return n
}

// Val gets number value.
func (n Number) Val() interface{} {
	return n.Value
}

// String number to string.
func (n Number) String() string {
	return fmt.Sprintf("number key: %s, value: %v", n.Key, n.Value)
}
