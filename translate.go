package translation

import (
	"fmt"

	"gitoa.ru/go-4devs/translation/arg"
	"golang.org/x/text/language"
)

// Option configures translate.
type Option func(*Translate)

// Translate configure translate.
type Translate struct {
	Domain string
	Locale language.Tag
	Args   []Arg
}

// ArgValues gets value arguments.
func (t *Translate) ArgValues() []interface{} {
	a := make([]interface{}, len(t.Args))
	for i, v := range t.Args {
		a[i] = v.Val()
	}

	return a
}

// Arg arg translate.
type Arg interface {
	Val() interface{}
	fmt.Stringer
}

// WithNumber sets number with options.
func WithNumber(key string, val interface{}, opts ...arg.NumberOption) Option {
	return func(t *Translate) {
		t.Args = append(t.Args, arg.Number{Key: key, Value: val}.Configure(opts...))
	}
}

// WithCurrency sets date argument.
func WithCurrency(key string, val interface{}, opts ...arg.CurrencyOption) Option {
	return func(t *Translate) {
		c := arg.Currency{Value: val, Key: key}

		for _, o := range opts {
			o(&c)
		}

		t.Args = append(t.Args, c)
	}
}

// WithDomain sets domain translate.
func WithDomain(domain string) Option {
	return func(o *Translate) {
		o.Domain = domain
	}
}

// WithLocale sets locale translate.
func WithLocale(locale string) Option {
	return func(o *Translate) {
		o.Locale = language.Make(locale)
	}
}

// WithArgs sets arguments value.
func WithArgs(vals ...interface{}) Option {
	return func(t *Translate) {
		args := make([]Arg, len(vals))
		for i, val := range vals {
			args[i] = arg.Arg{Value: val}
		}

		t.Args = append(t.Args, args...)
	}
}
