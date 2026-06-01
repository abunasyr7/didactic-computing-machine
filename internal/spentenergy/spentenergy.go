package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}

	if duration <=0 {
		return 0, errors.New("duration must be greather than 0")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	return ((weight * averageSpeed * duration.Minutes())/minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}

	if duration <=0 {
		return 0, errors.New("duration must be greather than 0")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	return (weight * averageSpeed * duration.Minutes())/minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	if steps <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return stepLength * float64(steps) / mInKm
}
