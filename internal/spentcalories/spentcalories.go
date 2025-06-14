package spentcalories

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

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 3 {
		return 0, "", 0, errors.New("")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, "", 0, errors.New("Неизвестный формат количества шагов")
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("Количество шагов равно 0")
	}
	activ := dataSlice[1]
	durationWalk, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, "", 0, errors.New("Неизвестный формат времени прогулки")
	}
	if durationWalk <= 0 {
		return 0, "", 0, errors.New("Невероное время")
	}
	return steps, activ, durationWalk, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenghtStep := height * stepLengthCoefficient
	dist := lenghtStep * float64(steps) / mInKm
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	time := duration.Minutes()
	midSpeed := (dist / time) * 60
	return midSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activ, time, err := parseTraining(data)
	if err != nil {
		return "", errors.New("Ошибка ввода данных")
	}
	switch activ {
	case "Бег":
		dist := distance(steps, height)
		duration := time.Hours()
		speed := meanSpeed(steps, height, time)
		calories, err := RunningSpentCalories(steps, weight, height, time)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", duration, dist, speed, calories), nil
	case "Ходьба":
		dist := distance(steps, height)
		duration := time.Hours()
		speed := meanSpeed(steps, height, time)
		calories, err := WalkingSpentCalories(steps, weight, height, time)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", duration, dist, speed, calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Некорректные значения")
	}
	midSpeed := meanSpeed(steps, height, duration)
	timeMinutes := duration.Minutes()
	runningCalories := (weight * midSpeed * timeMinutes) / minInH
	return runningCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Некорректные значения")
	}
	midSpeed := meanSpeed(steps, height, duration)
	timeMinutes := duration.Minutes()
	walkingCalories := ((weight * midSpeed * timeMinutes) / minInH) * walkingCaloriesCoefficient
	return walkingCalories, nil
}
