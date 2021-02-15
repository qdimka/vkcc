package vkUtils

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

type ShortenerOptions struct {
	Token string
}

type ShortenerService struct {
	options *ShortenerOptions
	api     *api.VK
}

func NewShortener(options *ShortenerOptions) *ShortenerService {
	return &ShortenerService{
		options: options,
		api:     api.NewVK(options.Token),
	}
}

func (service *ShortenerService) CreateLink(url string) (string, error) {
	response, err := service.api.UtilsGetShortLink(api.Params{
		"url":     url,
		"private": 0,
	})

	if err != nil {
		return "", err
	}

	return response.ShortURL, nil
}
