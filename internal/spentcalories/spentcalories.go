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
		return 0, "", 0, fmt.Errorf("unknown step count format: %v", err)
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("the number of steps is 0")
	}
	activ := dataSlice[1]
	durationWalk, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("unknown walk time format: %v", err)
	}
	if durationWalk <= 0 {
		return 0, "", 0, errors.New("wrong time")
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
	time := duration.Hours()
	midSpeed := dist / time
	return midSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activ, time, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("data entry error: %v", err)
	}
	dist := distance(steps, height)
	duration := time.Hours()
	speed := meanSpeed(steps, height, time)
	switch activ {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, time)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", duration, dist, speed, calories), nil
	case "Ходьба":
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
	if steps <= 0 {
		return 0, errors.New("incorrect values: steps")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect values: weight")
	}
	if height <= 0 {
		return 0, errors.New("incorrect values: height")
	}
	if duration <= 0 {
		return 0, errors.New("нincorrect values: duration")
	}
	midSpeed := meanSpeed(steps, height, duration)
	timeMinutes := duration.Minutes()
	runningCalories := (weight * midSpeed * timeMinutes) / minInH
	return runningCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("incorrect values: steps")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect values: weight")
	}
	if height <= 0 {
		return 0, errors.New("incorrect values: height")
	}
	if duration <= 0 {
		return 0, errors.New("нincorrect values: duration")
	}
	midSpeed := meanSpeed(steps, height, duration)
	timeMinutes := duration.Minutes()
	walkingCalories := ((weight * midSpeed * timeMinutes) / minInH) * walkingCaloriesCoefficient
	return walkingCalories, nil
}
