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

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")

	if len(data) != 2 {
		return errors.New("the data is not full")
	}

	steps, err := strconv.Atoi(data[0])

	if err != nil {
		return errors.New("something went wrong in steps")
	}

	if steps <= 0 {
		return errors.New("steps must be greater than 0")
	}

	ds.Steps = steps
	
	parsedTime, err := time.ParseDuration(data[1])

	if err != nil {
		return errors.New("something went wrong")
	}

	if parsedTime <= 0 {
		return errors.New("duration must be greater than 0")
	}

	ds.Duration = parsedTime

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", 
		ds.Steps, distance, calories), nil
	
}
