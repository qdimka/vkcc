package vkUtils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test(t *testing.T) {
	shortener := NewShortener(&ShortenerOptions{
		Token: os.Getenv("ServiceKey"),
	})

	shortUrl, err := shortener.CreateLink("https://www.google.com")

	assert.Nil(t, err)
	assert.NotEmpty(t, shortUrl)
}
