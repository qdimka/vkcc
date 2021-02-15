package main

import (
	"flag"
	"log"
	"time"
	"vkcc/files"
	vkUtils2 "vkcc/vkUtils"
)

func main() {
	tokenFlag := flag.String("token", "", "Сервисный токен Вконтакте")
	fromFlag := flag.String("from", "links.txt", "Путь к файлу источнику")
	toFlag := flag.String("to", "updated_links.txt", "Путь к файлу результату")
	failedFlag := flag.String("failed", "failed_links.txt", "Список ссылок, которые не удалось сократить")
	sleepFlag := flag.Duration("sleep", 1000, "Каждые 20 запросов, делаем паузу в 1000 миллисекунд")

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
	var failed []string
	for i, line := range lines {
		link, err := vkUtils.CreateLink(line)
		if err != nil {
			log.Printf("При сокращении ссылки %s возникла ошибка %s", line, err.Error())
			failed = append(failed, link)
			continue
		}
		shortUrls = append(shortUrls, link)
		log.Printf("Создание короткой ссылки для %s завершено", line)

		if i != 0 && i%20 == 0 {
			time.Sleep((*sleepFlag) * time.Millisecond)
		}
	}

	if len(failed) != 0 {
		err = files.SaveLines(*failedFlag, failed)
		if err != nil {
			log.Printf("При сохранении произошла ошибка: %s", err.Error())
		}
	}

	err = files.SaveLines(*toFlag, shortUrls)
	if err != nil {
		log.Printf("При сохранении произошла ошибка: %s", err.Error())
	}
}
