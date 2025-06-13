package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Onlyrr/tracker/internal/spentcalories"
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
	return steps, durationWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, durationWalk, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("Ошибка преобразования %v", err)
	}
	if steps <= 0 {
		return ""
	}
	var dist float64
	dist = float64(steps) * stepLength
	var distKm float64
	distKm = dist / float64(mInKm)
	Calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, durationWalk)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал .\n", steps, distKm, Calories)
}
