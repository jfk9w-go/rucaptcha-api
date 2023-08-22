package example

import (
	"context"
	"fmt"

	"github.com/caarlos0/env"

	"github.com/jfk9w-go/rucaptcha-api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var config struct {
		Key     string `env:"KEY,required"`
		SiteKey string `env:"SITE_KEY,required"`
		PageURL string `env:"PAGE_URL,required"`
	}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	client, err := rucaptcha.ClientBuilder{
		Config: rucaptcha.Config{
			Key: config.Key,
		},
	}.Build()
	if err != nil {
		panic(err)
	}

	result, err := client.Solve(ctx, &rucaptcha.YandexSmartCaptchaIn{
		SiteKey: config.SiteKey,
		PageURL: config.PageURL,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", result.Answer)
}
