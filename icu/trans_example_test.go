package icu_test

import (
	"context"
	"fmt"
	"log"

	"gitoa.ru/go-4devs/translation"
	"gitoa.ru/go-4devs/translation/arg"
	"gitoa.ru/go-4devs/translation/icu"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func ExampleTrans_withLanguage() {
	err := message.Set(language.Russian, "Hello {city}", catalog.String("Привет %s"))
	if err != nil {
		log.Fatal(err)
	}

	err = message.Set(language.Russian, "It costs {cost}", catalog.String("Это стоит %.2f."))
	if err != nil {
		log.Fatal(err)
	}

	lang, err := language.Parse("ru")
	if err != nil {
		log.Fatal(err)
	}

	ctx := translation.WithLanguage(context.Background(), lang)

	tr := icu.Trans(ctx, "Hello {city}", translation.WithArgs("Москва"))
	fmt.Println(tr)
	tr = icu.Trans(ctx, "It costs {cost}", translation.WithNumber("cost", 1000.00, arg.WithNumberFormat(arg.NumberFormatDecimal)))
	fmt.Println(tr)
	// Output:
	// Привет Москва
	// Это стоит 1 000,00.
}
