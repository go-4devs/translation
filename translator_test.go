package translation_test

import (
	"context"
	"fmt"
	"testing"

	"gitoa.ru/go-4devs/translation"
	"golang.org/x/text/language"
)

type TestProvider struct{}

func (tp *TestProvider) Translate(_ context.Context, key string, opt translation.Translate) string {
	args := make([]interface{}, 0, len(opt.Args)+1)
	args = append(args, key)
	args = append(args, opt.ArgValues()...)

	return fmt.Sprint(args...)
}

func TestTranslator(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	trans := translation.New(language.Russian, &TestProvider{})
	tr := trans.Trans(ctx, "key", translation.WithArgs("arg1", "arg2"))

	if tr != "keyarg1arg2" {
		t.Fatalf("expect: keyarg1arg2, got:%s", tr)
	}
}
