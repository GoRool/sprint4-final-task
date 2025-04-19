package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)


const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)


func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
  if len(parts) != 3 {
  	return 0, "", 0, errors.New("need 3 elements in data")
	}

  stepsStr := parts[0]
  actionType := parts[1]
  durationStr := parts[2]

	steps, err := strconv.Atoi(stepsStr) 
  if err != nil || steps <= 0 {
    return 0, "", 0, errors.New("invalid data of steps")
  }

	duration, err := time.ParseDuration(durationStr)
  if err != nil || duration <= 0 {
    return 0, "", 0, errors.New("invalid or negativ duration")
  }
	return steps, actionType, duration, nil
}


func distance(steps int, height float64) float64 {
	lengthOfStep := height * stepLengthCoefficient
  totalDistance := lengthOfStep * float64(steps)
  return totalDistance / mInKm
}


func meanSpeed(steps int, height float64, duration time.Duration) float64 {
  if duration.Hours() <= 0 {
    return 0
  }

  kms := distance(steps, height)
	hours := duration.Hours()
  return kms / hours
}


func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
  if err != nil {
    return "", err
  }
  var calories float64
  switch activityType {
  	case "Ходьба":
  	  calories, _ = WalkingSpentCalories(steps, weight, height, duration)
  	case "Бег":
  	  calories, _ = RunningSpentCalories(steps, weight, height, duration)
  	default:
  	  return "", errors.New("неизвестный тип тренировки")
  }

	if err != nil {
		log.Printf("error processing colories: %v", err)
	}
  speed := meanSpeed(steps, height, duration)
  report := fmt.Sprintf(
    "Тип тренировки: %s\n"+
    "Длительность: %.2f ч.\n"+
    "Дистанция: %.2f км.\n"+
    "Скорость: %.2f км/ч\n"+
    "Сожгли калорий: %.2f\n",
    activityType, duration.Hours(), distance(steps, height), speed, calories,
  )
  return report, nil
}


func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
  if steps <= 0 || weight <= 0 || height <= 0 || duration.Minutes() <= 0 {
    return 0, errors.New("invalid input params")
  }

  speed := meanSpeed(steps, height, duration)
  caloriesPerMinute := speed * weight
  totalCalories := caloriesPerMinute * duration.Minutes() / minInH	
  return totalCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration.Minutes() <= 0 {
    return 0, errors.New("invalid input params")
  }

	speed := meanSpeed(steps, height, duration)
	caloriesPerMinute := speed * weight
	totalCalories := caloriesPerMinute * duration.Minutes() / minInH * walkingCaloriesCoefficient
	return totalCalories, nil
}
