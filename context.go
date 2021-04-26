package translation

import (
	"context"

	"golang.org/x/text/language"
)

type ctxkey uint8

const (
	localeKey ctxkey = iota
)

// WithLanguage sets language to context.
func WithLanguage(ctx context.Context, lang language.Tag) context.Context {
	return context.WithValue(ctx, localeKey, lang)
}

// FromContext get language from context or use default.
func FromContext(ctx context.Context, def language.Tag) language.Tag {
	if t, ok := ctx.Value(localeKey).(language.Tag); ok {
		return t
	}

	return def
}
