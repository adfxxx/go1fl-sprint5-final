package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	text := strings.Split(datastring, ",")
	if len(text) != 2 {
		return fmt.Errorf("wrong number of args")
	}
	ds.Steps, err = strconv.Atoi(text[0])
	if ds.Steps <= 0 {
		return fmt.Errorf("wrong args")
	}
	if err != nil {
		return nil
	}
	ds.Duration, err = time.ParseDuration(text[1])
	if ds.Duration <= 0 {
		return fmt.Errorf("wrong args")
	}
	if err != nil {
		return nil
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	caloriesNumber, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, caloriesNumber), nil
}
