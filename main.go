package main

import (
	"flag"
	"log"
	"vkcc/files"
	vkUtils2 "vkcc/vkUtils"
)

func main() {
	tokenFlag := flag.String("token", "", "Сервисный токен Вконтакте")
	fromFlag := flag.String("from", "links.txt", "Путь к файлу источнику")
	toFlag := flag.String("to", "updated_links.txt", "Путь к файлу результату")

	flag.Parse()

	vkUtils := vkUtils2.NewShortener(&vkUtils2.ShortenerOptions{
		Token: *tokenFlag,
	})

	lines, err := files.ReadAllLines(*fromFlag)
	if err != nil {
		log.Printf("Ошибка при открытии файла: %s", err.Error())
		return
	}

	var shortUrls []string
	for _, line := range lines {
		link, err := vkUtils.CreateLink(line)
		if err != nil {
			log.Printf("При сокращении ссылки %s возникла ошибка %s", line, err.Error())
			continue
		}
		shortUrls = append(shortUrls, link)
	}

	_ = files.SaveLines(*toFlag, shortUrls)
}
