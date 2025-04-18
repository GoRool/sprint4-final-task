package spentcalories

// Пакет spentcalories
// В этом пакете вы будете реализовывать шесть функций: три экспортируемые и три неэкспортируемые, вспомогательные.
// Начнём со вспомогательных функций.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)



// Функция parseTraining()
// Сигнатура функции:
// func parseTraining(data string) (int, string, time.Duration, error) 
// Функция принимает строку с данными формата "3456,Ходьба,3h00m", которая содержит количество шагов, вид активности и продолжительность активности. Функция возвращает четыре значения:
// int — количество шагов.
// string — вид активности.
// time.Duration — продолжительность активности.
// error — ошибку, если что-то пошло не так.
// Функция парсит строку, переводит данные из строки в соответствующие типы и возвращает эти значения.
// Алгоритм реализации функции:

func parseTraining(data string) (int, string, time.Duration, error) {

	// TODO: реализовать функцию
	// Разделить строку на слайс строк.
// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас 
// количество шагов, вид активности и продолжительность.
	parts := strings.Split(data, ",")
  if len(parts) != 3 {
  	return 0, "", 0, errors.New("bad data. need 3 elements")
	}
// Преобразовать первый элемент слайса (количество шагов) в тип int. 
    stepsStr := parts[0]
    actionType := parts[1]
    durationStr := parts[2]

	steps, err := strconv.Atoi(stepsStr)
// Обработать возможные ошибки. 
    if err != nil || steps <= 0 {
			// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
        return 0, "", 0, errors.New("bad data. bad steps")
    }



// Преобразовать третий элемент слайса в time.Duration. 
	duration, err := time.ParseDuration(durationStr)
// Обработать возможные ошибки. 
  if err != nil {

// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
    return 0, "", 0, errors.New("bad data. bad duration")
  }
// Если всё прошло без ошибок, 
// верните количество шагов, 
// вид активности, 
// продолжительность 
// и nil (для ошибки).
	return steps, actionType, duration, nil

  
}


// Функция distance()
// Сигнатура функции:

// func distance(steps int, height float64) float64 
// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
// Для вычисления дистанции:
// рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient. Соответствующая константа уже определена в пакете.
// умножьте пройденное количество шагов на длину шага.
// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
// Обратите внимание, что целочисленную переменную steps необходимо будет привести к другому числовому типу.
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthOfStep := height * stepLengthCoefficient
  totalDistance := lengthOfStep * float64(steps)
  return totalDistance / mInKm
}



// Функция meanSpeed()
// Сигнатура функции:
// func meanSpeed(steps int, height float64, duration time.Duration) float64 
// Функция принимает количество шагов steps, рост пользователя height и продолжительность активности duration  и возвращает среднюю скорость.
// Алгоритм реализации функции:
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
  if duration.Hours() == 0 {
    return 0
  }
// Вычислить дистанцию с помощью distance().
  kms := distance(steps, height)
// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах. 
// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	hours := duration.Hours()
  return kms / hours
}

//  Функция TrainingInfo()
// Сигнатура функции:
// func TrainingInfo(data string, weight, height float64) (string, error) 
// Функция принимает:
// data string — строку с данными формата "3456,Ходьба,3h00m", которая содержит количество шагов, вид активности и продолжительность активности.
// weight, height float64 — вес (кг.) и рост (м.) пользователя.
// И возвращает два значения:
// string — строка с информацией о тренировке в формате, приведенном ниже.
// error — ошибку, при ее возникновении внутри функции.
// Пример возвращаемой строки:
// Тип тренировки: Бег
// Длительность: 0.75 ч.
// Дистанция: 10.00 км.
// Скорость: 13.34 км/ч
// Сожгли калорий: 18621.75 

func TrainingInfo(data string, weight, height float64) (string, error) {
// TODO: реализовать функцию
// Получить значения из строки данных с помощью функции parseTraining(),
	steps, activityType, duration, err := parseTraining(data)
//  обработать возможные ошибки.
  if err != nil {
    return "", err
  }
// Проверить, какой вид тренировки был передан в строке, 
// которую парсили (лучше использовать switch). 
  var calories float64
  switch activityType {
  	case "Ходьба":
  	  calories, _ = WalkingSpentCalories(steps, weight, height, duration)
  	case "Бег":
  	  calories, _ = RunningSpentCalories(steps, weight, height, duration)
// Если был передан неизвестный тип тренировки, 
// вернуть ошибку с текстом неизвестный тип тренировки.
  	default:
  	  return "", errors.New("unknown type of training")
  }
// Для каждого из видов тренировки рассчитать дистанцию, среднюю скорость и калории.
// Для каждого вида тренировки сформировать и вернуть строку,
  speed := meanSpeed(steps, height, duration)

  report := fmt.Sprintf(
    "Тип тренировки: %s\n"+
    "Длительность: %.2f ч.\n"+
    "Дистанция: %.2f км.\n"+
    "Скорость: %.2f км/ч\n"+
    "Сожгли калорий: %.2f",
    activityType, duration.Hours(), distance(steps, height), speed, calories,
  )

  return report, nil
}



// Функция RunningSpentCalories()
// Сигнатура функции:
// func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) 
// Функция принимает:
// steps int — количество шагов.
// weight, height float64 — вес(кг.) и рост(м.) пользователя.
// duration time.Duration — продолжительность бега.
// И возвращает два значения:
// float64 — количество калорий, потраченных при беге.
// error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
// Алгоритм реализации функции:

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
// TODO: реализовать функцию

  if steps <= 0 || weight <= 0 || height <= 0 || duration.Minutes() <= 0 {
    return 0, errors.New("invalid data for calculating calories burned while running")
  }
// Рассчитать среднюю скорость с помощью meanSpeed().
  speed := meanSpeed(steps, height, duration)

	// Рассчитать и вернуть количество калорий. 
  caloriesPerMinute := speed * weight
  totalCalories := caloriesPerMinute * duration.Minutes() / minInH

  return totalCalories, nil
}



// Функция WalkingSpentCalories()
// Сигнатура функции:
// func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) 
// Она совпадает с сигнатурой RunningSpentCalories(). Функция принимает:
// steps int — количество шагов.
// weight, height float64 — вес(кг.) и рост(м.) пользователя.
// duration time.Duration — продолжительность бега.
// И возвращает два значения:
// float64 — количество калорий, потраченных при ходьбе.
// error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
// Алгоритм реализации функции:
// Проверить входные параметры на корректность. 
// Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
// Рассчитать среднюю скорость с помощью meanSpeed().
// Рассчитать и количество калорий. Для этого:
// Переведите продолжительность в минуты с помощью функции из пакета time.
// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
// Разделите результат на число минут в часе для получения количества потраченных калорий.
// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient. 
// Соответствующая константа объявляена в пакете. Вернуть полученное значение.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration.Minutes() <= 0 {
    return 0, errors.New("недопустимые входные параметры")
  }

	speed := meanSpeed(steps, height, duration)
	caloriesPerMinute := speed * weight
	totalCalories := caloriesPerMinute * duration.Minutes() / minInH * walkingCaloriesCoefficient

	return totalCalories, nil
}
