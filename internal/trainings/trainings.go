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

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	words := strings.Split(datastring, ",")

	if len(words) != 3 {
		return errors.New("the data is not full")
	}

	steps, err := strconv.Atoi(words[0])

	if err != nil {
		return errors.New("something went wrong")
	}

	if steps <= 0 {
		return errors.New("steps must be greater than 0")
	}

	t.Steps = steps

	t.TrainingType  = words[1]

	parsedTime, err := time.ParseDuration(words[2])

	if err != nil {
		return errors.New("something went wrong")
	}

	if parsedTime <= 0 {
		return errors.New("duration must be greater than 0")
	}

	t.Duration = parsedTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, averageSpeed, calories,
		), nil
}
