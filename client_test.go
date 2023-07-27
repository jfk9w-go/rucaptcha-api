package rucaptcha_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/require"

	"github.com/jfk9w-go/rucaptcha-api"
)

func TestClient_Solve_Yandex(t *testing.T) {
	var config struct {
		Key     string `env:"KEY,required"`
		SiteKey string `env:"SITE_KEY,required"`
		PageURL string `env:"PAGE_URL,required"`
	}

	if err := env.Parse(&config); err != nil {
		t.Skipf(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := rucaptcha.ClientBuilder{
		Config: rucaptcha.Config{
			Key: config.Key,
		},
	}.Build(ctx)
	require.NoError(t, err)

	result, err := client.Solve(ctx, &rucaptcha.YandexSmartCaptchaIn{
		SiteKey: config.SiteKey,
		PageURL: config.PageURL,
	})

	require.NoError(t, err)

	t.Logf("Result: %s", result.Answer)
}
