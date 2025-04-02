package daysteps

import (
	"time"
	"strings"
	"strconv"
	"fmt"
	"go1fl-4-sprint-final/internal"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
dataSlise := strings.Split(data, ",")
if len(dataSlise) != 2 {
	return 0, 0, fmt.Errorf("Ошибка данных")
}
// выделяем шаги
steps, err :=  strconv.Atoi(dataSlise[0])
if err != nil {
	return 0, 0, fmt.Errorf("Ошибка формата шагов")
}
if steps <= 0 {
	return 0, 0, fmt.Errorf("Нет шагов")
}
// выделяем время
duration, err := time.ParseDuration(strings.Replace(dataSlise[1], "h", "h", 1))
    if err != nil {
        return 0, 0, fmt.Errorf("Ошибка данных времени")
    }

    return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return ""
    }
    if steps <= 0 {
        return ""
    }
	// вычисляем дистанцию
    distance := (float64(steps) * StepLength) / 1000.0
    // надо ещё вычислить калории
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал", steps, distance, calories)
}
