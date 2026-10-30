package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	text := strings.Split(datastring, ",")
	if len(text) != 3 {
		return fmt.Errorf("wrong number of args")
	}
	t.Steps, err = strconv.Atoi(text[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return fmt.Errorf("wrong args")
	}
	t.TrainingType = text[1]
	t.Duration, err = time.ParseDuration(text[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return fmt.Errorf("wrong args")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var caloriesNumber float64
	var err error
	switch t.TrainingType {
	case "Бег":
		caloriesNumber, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		caloriesNumber, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, caloriesNumber), nil
}
