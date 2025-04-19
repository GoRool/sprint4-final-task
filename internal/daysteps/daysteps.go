package daysteps

// Пакет содержит две функции:
// одну экспортируемую
// и одну неэкспортируемую — вспомогательную.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",") 
	if len(parts) != 2 {
		return 0, 0, errors.New("bad data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <=0 {
		return	0, 0 ,errors.New("bad data in steps")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return 0, 0, errors.New("error conversion of time")
	}

	return steps, duration, nil
}


func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err.Error())
	return ""
	}

	distanceMeters := float64(steps) * stepLength
  distanceKilometers := distanceMeters / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKilometers, spentCalories)
	return result
}
