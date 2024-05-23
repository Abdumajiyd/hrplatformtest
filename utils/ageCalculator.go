package utils

import (
    "time"
)

func CalculateAge(birthdate time.Time) int {
    now := time.Now()
    years := now.Year() - birthdate.Year()
    if now.YearDay() < birthdate.YearDay() {
        years--
    }
    return years
}
