package rucaptcha_test

import (
	"context"
	"os"
	"testing"

	"github.com/jfk9w-go/based"
	"github.com/jfk9w-go/rucaptcha-api"

	"github.com/stretchr/testify/require"
)

func TestClient_Solve_Yandex(t *testing.T) {
	key := os.Getenv("KEY")
	siteKey := os.Getenv("SITE_KEY")
	pageURL := os.Getenv("PAGE_URL")

	if key == "" || siteKey == "" || pageURL == "" {
		t.Skipf("KEY, SITE_KEY and PAGE_URL environment variables must be specified")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &rucaptcha.Config{
		Key:      key,
		Pingback: os.Getenv("PINGBACK"),
	}

	client, err := rucaptcha.NewClient(based.StandardClock, config)
	require.NoError(t, err)

	result, err := client.Solve(ctx, &rucaptcha.YandexSmartCaptchaIn{
		SiteKey: siteKey,
		PageURL: pageURL,
	})

	require.NoError(t, err)

	t.Logf("Result: %s", result)
}
