package arg

import "fmt"

// FormatCurrency types format.
type FormatCurrency int

// Currency format.
const (
	CurrencyFormatSymbol FormatCurrency = iota + 1
	CurrencyFormatISO
	CurrencyFormatNarrowSymbol
)

// CurrencyOption configures option.
type CurrencyOption func(*Currency)

// WithCurrencyFormat sets format currency.
func WithCurrencyFormat(format FormatCurrency) CurrencyOption {
	return func(c *Currency) { c.Format = format }
}

// WithCurrencyISO sets ISO 4217 code currecy.
func WithCurrencyISO(iso string) CurrencyOption {
	return func(c *Currency) { c.ISO = iso }
}

// Currency argument.
type Currency struct {
	Key    string
	Value  interface{}
	Format FormatCurrency
	// ISO 3-letter ISO 4217
	ISO string
}

// String gets string from currency.
func (a Currency) String() string {
	return fmt.Sprintf("currency key:%s, value:%v", a.Key, a.Value)
}

// Val gets value currency.
func (a Currency) Val() interface{} {
	return a.Value
}
