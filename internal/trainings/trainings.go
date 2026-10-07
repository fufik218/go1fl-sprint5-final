package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training — экспортируемая структура с данными о тренировке.
type Training struct {
	Steps                 int           // количество шагов за тренировку
	TrainingType          string        // тип тренировки (бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенная структура с именем, весом и ростом пользователя
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и заполняет поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	// Разбиваем строку по запятой.
	parts := strings.Split(datastring, ",")

	// Должно быть ровно 3 части: шаги, тип тренировки, длительность.
	if len(parts) != 3 {
		return errors.New("неверный формат строки: ожидается 3 значения, разделённые запятой")
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
	t.Steps = steps

	// Сохраняем тип тренировки.
	t.TrainingType = parts[1]

	// Преобразуем длительность из строки в time.Duration.
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	// Длительность должна быть положительной.
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// Вычисляем дистанцию в километрах.
	distance := spentenergy.Distance(t.Steps, t.Height)

	// Вычисляем среднюю скорость в км/ч.
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// Переменная для количества сожжённых калорий.
	var calories float64

	// Выбираем функцию расчёта калорий в зависимости от типа тренировки.
	switch t.TrainingType {
	case "Бег":
		var err error
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		var err error
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// Формируем итоговую строку с завершающим переводом строки.
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)

	return result, nil
}
