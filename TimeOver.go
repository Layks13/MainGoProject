package main

import (
	"fmt"
	"time"
)

func main() {
	// Получаем текущее системное время
	currentTime := time.Now()

	// Задаем начало и конец рабочего дня
	workStart := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 8, 0, 0, 0, currentTime.Location())
	workEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 20, 0, 0, 0, currentTime.Location())

	// Вычисляем оставшееся время до конца рабочего дня
	remainingTime := workEnd.Sub(currentTime)

	// Проверяем, что текущее время находится в рабочее время и выводим оставшееся время
	if currentTime.Before(workStart) {
		fmt.Println("Рабочий день еще не начался")
	} else if currentTime.After(workEnd) {
		fmt.Println("Рабочий день уже закончился")
	} else {
		hours := int(remainingTime.Hours())
		minutes := int(remainingTime.Minutes()) % 60
		fmt.Printf("До конца рабочего времени осталось %d часов %d минут\n", hours, minutes)
	}
}