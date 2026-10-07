package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс с методами для парсинга данных и формирования информации.
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info перебирает все строки из dataset, парсит их через dp и выводит информацию.
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		// Парсим строку.
		err := dp.Parse(datastring)
		if err != nil {
			log.Println("ошибка парсинга:", err)
			continue
		}

		// Формируем информацию об активности.
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации:", err)
			continue
		}

		// Выводим информацию.
		fmt.Println(info)
	}
}
