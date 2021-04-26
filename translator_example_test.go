package translation_test

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/translation"
	"gitoa.ru/go-4devs/translation/arg"
	"gitoa.ru/go-4devs/translation/provider/gotext"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

func ExampleTranslator_Trans() {
	cat := catalog.NewBuilder()

	_ = cat.Set(language.Russian, "Hello World!", catalog.String("Привет Мир!"))
	_ = cat.Set(language.Slovak, "Hello World!", catalog.String("Ahoj Svet!"))
	_ = cat.Set(language.Hebrew, "Hello World!", catalog.String("שלום עולם"))
	_ = cat.Set(language.Russian, "Hello {name}!", catalog.String("Привет %[1]s!"))
	_ = cat.Set(language.English, "Hello {name}!", catalog.String("Hello %[1]s!"))
	_ = cat.Set(language.English, "There are {bikes} bikes per household.", catalog.String("There are %v bikes per household."))
	_ = cat.Set(language.Russian, "There are {bikes} bikes per household.", catalog.String("В каждом доме есть %v велосипеда."))
	_ = cat.Set(language.English, "You are {minute} minute(s) late.",
		plural.Selectf(1, "",
			"=0", "You're on time.",
			plural.One, "You are %v minute late.",
			plural.Other, "You are %v minutes late."))
	_ = cat.Set(language.Russian, "You are {minute} minute(s) late.",
		plural.Selectf(1, "",
			"=1", "Вы опоздали на одну минуту.",
			"=0", "Вы вовремя.",
			plural.One, "Вы опоздали на %v минуту.",
			plural.Few, "Вы опоздали на %v минуты.",
			plural.Other, "Вы опоздали на %v минут.",
		),
	)
	_ = cat.Set(language.Russian, "It costs {cost}.",
		catalog.String("Это стоит %.2f."))
	_ = cat.Set(language.English, "It costs {cost}.",
		catalog.String("It costs %.2f."))

	provider := gotext.NewProvider(gotext.WithCatalog(translation.DefaultDomain, cat))

	trans := translation.New(language.Russian, provider)
	ctx := context.Background()
	// context with language
	heCtx := translation.WithLanguage(ctx, language.Make("he"))

	fmt.Println(trans.Trans(ctx, "Hello World!"))
	fmt.Println(trans.Trans(ctx, "Hello World!", translation.WithLocale("en")))
	fmt.Println(trans.Trans(ctx, "Hello World!", translation.WithLocale("sk")))
	fmt.Println(trans.Trans(heCtx, "Hello World!"))

	fmt.Println(trans.Trans(ctx, "Hello {name}!", translation.WithArgs("Andrey")))

	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithNumber("minute", 1)))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithNumber("minute", 0)))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithNumber("minute", 101)))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithArgs(123456.78)))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithArgs(50)))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithArgs(1), translation.WithLocale("en")))
	fmt.Println(trans.Trans(ctx, "You are {minute} minute(s) late.", translation.WithArgs(33), translation.WithLocale("en")))

	fmt.Println(trans.Trans(ctx, "There are {bikes} bikes per household.", translation.WithNumber("bikes", 1.2)))
	fmt.Println(trans.Trans(ctx, "There are {bikes} bikes per household.", translation.WithNumber("bikes", 1.2), translation.WithLocale("en")))

	fmt.Println(trans.Trans(ctx, "It costs {cost}.",
		translation.WithCurrency("cost", 12.0101, arg.WithCurrencyISO("rub"), arg.WithCurrencyFormat(arg.CurrencyFormatSymbol))))
	fmt.Println(trans.Trans(ctx, "It costs {cost}.",
		translation.WithCurrency("cost", 15.0, arg.WithCurrencyISO("HKD"), arg.WithCurrencyFormat(arg.CurrencyFormatSymbol)),
		translation.WithLocale("en")))
	// Output:
	// Привет Мир!
	// Hello World!
	// Ahoj Svet!
	// שלום עולם
	// Привет Andrey!
	// Вы опоздали на одну минуту.
	// Вы вовремя.
	// Вы опоздали на 101 минуту.
	// Вы опоздали на 123 456,78 минут.
	// Вы опоздали на 50 минут.
	// You are 1 minute late.
	// You are 33 minutes late.
	// В каждом доме есть 1,2 велосипеда.
	// There are 1.2 bikes per household.
	// Это стоит ₽ 12.0101.
	// It costs HK$ 15.
}
