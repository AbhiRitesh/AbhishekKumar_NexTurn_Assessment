package main

import (
	"fmt"
	"strings"
)

// store climate data for each city
type City struct {
	Name         string
	Temperature  float64 
	Rainfall     float64 
}

// find the city with the highest temperature
func HighestTemp(cities []City) City {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// find the city with the lowest temperature
func LowestTemp(cities []City) City {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// calculate the average rainfall
func AvgRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// filter cities by a rainfall threshold
func filterByRainfall(cities []City, threshold float64) []City {
	filteredCities := []City{}
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

// search for a city by name
func searchByName(cities []City, name string) (City, bool) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, true
		}
	}
	return City{}, false
}

func main() {
	cities := []City{
		{"New York", 22.5, 120.0},
		{"London", 18.0, 85.0},
		{"Mumbai", 30.5, 220.0},
		{"Tokyo", 25.0, 150.0},
		{"Sydney", 20.0, 95.0},
	}

	fmt.Println("Climate Data Analysis")

	
	highestTempCity := HighestTemp(cities)
	fmt.Printf("\nCity with the highest temp: %s (%.2f°C)\n", highestTempCity.Name, highestTempCity.Temperature)

	lowestTempCity := LowestTemp(cities)
	fmt.Printf("City with the lowest temp: %s (%.2f°C)\n", lowestTempCity.Name, lowestTempCity.Temperature)

	
	
	averageRainfall := AvgRainfall(cities)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	var threshold float64
	fmt.Print("\nEnter a rainfall threshold (mm): ")
	_, err := fmt.Scan(&threshold)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid value.")
		return
	}

	filteredCities := filterByRainfall(cities, threshold)
	fmt.Printf("\nCities with rainfall above %.2f mm:\n", threshold)
	if len(filteredCities) == 0 {
		fmt.Println("No cities found.")
	} else {
		for _, city := range filteredCities {
			fmt.Printf("- %s (%.2f mm)\n", city.Name, city.Rainfall)
		}
	}

	
	var cityName string
	fmt.Print("\nEnter a city name: ")
	fmt.Scan(&cityName)

	foundCity, found := searchByName(cities, cityName)
	if found {
		fmt.Printf("\nCity found: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n",
			foundCity.Name, foundCity.Temperature, foundCity.Rainfall)
	} else {
		fmt.Println("\nCity not found.")
	}
}
