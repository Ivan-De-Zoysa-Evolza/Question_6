package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		temp := getValidTemperature()
		unit := getValidUnit()

		if unit == "C" {
			result := celsiusToFahrenheit(temp)
			fmt.Printf("%.2f°C is %.2f°F\n", temp, result)
		} else {
			result := fahrenheitToCelsius(temp)
			fmt.Printf("%.2f°F is %.2f°C\n", temp, result)
		}

		fmt.Print("\nDo you want to convert another temperature? (yes/no): ")
		var again string
		fmt.Scanln(&again)
		if strings.ToLower(again) != "yes" {
			fmt.Println("\nGoodbye!")
			break
		}
	}
}

// Function to validate temperature input (any number)
func getValidTemperature() float64 {
	for {
		fmt.Print("\nEnter the temperature value: ")
		var input string
		fmt.Scanln(&input)

		temp, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("\nInvalid input. Please enter a valid number.")
			continue
		}
		return temp
	}
}

// Function to validate unit selection
func getValidUnit() string {
	for {
		fmt.Print("Convert to (C)elsius or (F)ahrenheit? Enter C or F: ")
		var unit string
		fmt.Scanln(&unit)
		upper := strings.ToUpper(unit)
		if upper == "C" || upper == "F" {
			return upper
		}
		fmt.Println("\nInvalid unit. Please enter C or F.")
	}
}

// Conversion functions
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
