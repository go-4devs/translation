package icu

import (
	"context"

	"gitoa.ru/go-4devs/translation"
	"gitoa.ru/go-4devs/translation/provider/gotext"
	"golang.org/x/text/language"
)

// nolint: gochecknoglobals
var trans = translation.New(language.English, gotext.NewProvider())

// Trans translate message.
func Trans(ctx context.Context, key string, opts ...translation.Option) string {
	return trans.Trans(ctx, key, opts...)
}

// SetTranslator sets translator.
func SetTranslator(tr *translation.Translator) {
	trans = tr
}
