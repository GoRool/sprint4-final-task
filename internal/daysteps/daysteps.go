package daysteps

// Пакет содержит две функции:
// одну экспортируемую
// и одну неэкспортируемую — вспомогательную.

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	stepLength = 0.65	// Длина одного шага в метрах
	mInKm = 1000// Количество метров в одном километре
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
		log.Printf("data processing error: %v",err)
	return ""
	}

	distanceMeters := float64(steps) * stepLength
  distanceKilometers := distanceMeters / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("duration processing error: %v",err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKilometers, spentCalories)
	return result
}
