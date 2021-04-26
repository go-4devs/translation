package gotext

import (
	"context"

	"gitoa.ru/go-4devs/translation"
	"gitoa.ru/go-4devs/translation/arg"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"golang.org/x/text/number"
)

var _ translation.Provider = (*Provider)(nil)

// Option confires message provider.
type Option func(*Provider)

// WithCatalog set coatalog bu domain name.
func WithCatalog(domain string, cat catalog.Catalog) Option {
	return func(mp *Provider) {
		mp.catalog[domain] = cat
	}
}

// NewProvider creates new messgae provider.
func NewProvider(opts ...Option) *Provider {
	mp := &Provider{
		catalog: map[string]catalog.Catalog{
			translation.DefaultDomain: message.DefaultCatalog,
		},
	}

	for _, o := range opts {
		o(mp)
	}

	return mp
}

// Provider provider messages.
type Provider struct {
	catalog map[string]catalog.Catalog
}

// Translate by key and args.
func (mp *Provider) Translate(ctx context.Context, key string, opt translation.Translate) string {
	return message.NewPrinter(opt.Locale, message.Catalog(mp.catalog[opt.Domain])).
		Sprintf(key, messageArgs(opt.Locale, opt.Args)...)
}

func currencyMessage(lang language.Tag) func(v arg.Currency) interface{} {
	return func(v arg.Currency) interface{} {
		var u currency.Unit

		if v.ISO != "" {
			u, _ = currency.ParseISO(v.ISO)
		} else {
			u, _ = currency.FromTag(lang)
		}

		switch v.Format {
		case arg.CurrencyFormatISO:
			return currency.ISO(u.Amount(v.Val()))
		case arg.CurrencyFormatSymbol:
			return currency.Symbol(u.Amount(v.Val()))
		case arg.CurrencyFormatNarrowSymbol:
			return currency.NarrowSymbol(u.Amount(v.Val()))
		}

		return u.Amount(v.Val())
	}
}

func numberMessage(v arg.Number) interface{} {
	switch v.Format {
	case arg.NumberFormatPercent:
		return number.Percent(v.Val())
	case arg.NumberFormatEngineering:
		return number.Engineering(v.Val())
	case arg.NumberFormatPerMille:
		return number.PerMille(v.Val())
	case arg.NumberFormatDecimal:
		return number.Decimal(v.Val())
	case arg.NumberFormatScientific:
		return number.Scientific(v.Val())
	}

	return v.Val()
}

func messageArgs(lang language.Tag, in []translation.Arg) []interface{} {
	out := make([]interface{}, 0, len(in))

	for _, a := range in {
		switch v := a.(type) {
		case arg.Currency:
			out = append(out, currencyMessage(lang)(v))
		case arg.Number:
			out = append(out, numberMessage(v))
		default:
			out = append(out, v.Val())
		}
	}

	return out
}
