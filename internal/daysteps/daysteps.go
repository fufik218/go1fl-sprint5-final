package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps — экспортируемая структура с данными о дневной прогулке.
type DaySteps struct {
	Steps                 int           // количество шагов за прогулку
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // встроенная структура с именем, весом и ростом пользователя
}

// Parse разбирает строку формата "678,0h50m" и заполняет поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разбиваем строку по запятой.
	parts := strings.Split(datastring, ",")

	// Должно быть ровно 2 части: шаги и длительность.
	if len(parts) != 2 {
		return errors.New("неверный формат строки: ожидается 2 значения, разделённые запятой")
	}

	// Преобразуем шаги из строки в int.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	// Шаги должны быть положительными.
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	ds.Steps = steps

	// Преобразуем длительность из строки в time.Duration.
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	// Длительность должна быть положительной.
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию в километрах.
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Считаем калории при ходьбе.
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем итоговую строку с завершающим переводом строки.
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
