package daysteps

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
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, errors.New("Неверные значения")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, errors.New("Неизвестный формат количества шагов")
	}
	if steps <= 0 {
		return 0, 0, errors.New("Количество шагов равно 0")
	}
	durationWalk, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, errors.New("Неизвестный формат времени прогулки")
	}
	if durationWalk <= 0 {
		return 0, 0, errors.New("Количество времени равно 0")
	}
	return steps, durationWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, durationWalk, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println("")
		return ""
	}
	var dist float64
	dist = float64(steps) * stepLength
	var distKm float64
	distKm = dist / float64(mInKm)
	Calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, durationWalk)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distKm, Calories)
}
