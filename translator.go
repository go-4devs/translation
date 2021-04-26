package translation

import (
	"context"

	"golang.org/x/text/language"
)

// Default values.
const (
	DefaultDomain = "messages"
)

// Provider for translate key.
type Provider interface {
	Translate(ctx context.Context, key string, opt Translate) string
}

// TranslatorOption options translator.
type TranslatorOption func(*Translator)

// New creates new translator.
func New(locale language.Tag, provider Provider, opts ...TranslatorOption) *Translator {
	tr := Translator{
		locale:   locale,
		provider: provider,
		domain:   DefaultDomain,
	}

	for _, o := range opts {
		o(&tr)
	}

	return &tr
}

// Translator struct.
type Translator struct {
	provider Provider
	domain   string
	locale   language.Tag
}

// Trans translates key by options.
func (t *Translator) Trans(ctx context.Context, key string, opts ...Option) string {
	opt := Translate{
		Locale: FromContext(ctx, t.locale),
		Domain: t.domain,
	}

	for _, o := range opts {
		o(&opt)
	}

	return t.provider.Translate(ctx, key, opt)
}
