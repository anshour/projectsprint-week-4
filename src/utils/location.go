package utils

import (
	"errors"
	"math"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
)

const earthRadius = 6371                  // Earth's radius in kilometers
const deliverySpeed = 40.0                // Speed in km/h
var diameter = 2 * math.Sqrt(3.0/math.Pi) // the diameter for a circle with an area of 3 km²

// Haversine function to calculate the distance between two locations
func Haversine(lat1, lon1, lat2, lon2 float64) (float64, error) {
	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance, nil
}

func IsWithin3Km2(locations []entity.Location) (bool, error) {
	if len(locations) == 0 {
		return false, constants.ErrNoLatLongPointProvided
	}

	// Ending point is the last element
	end := locations[len(locations)-1]

	// Calculate the distance between each point and the ending point
	for _, point := range locations {
		distance, err := Haversine(point.LocationLat, point.LocationLong, end.LocationLat, end.LocationLong)
		if err != nil {
			return false, errors.New("Error calculating distance:" + err.Error())
		}
		if distance > diameter {
			return false, nil
		}
	}

	return true, nil
}

func NearestNeighborTSP(locations []entity.Location) ([]int, float64, error) {
	n := len(locations)

	// Check for edge cases
	if n < 2 {
		return nil, 0, errors.New("not enough locations to form a route")
	}

	start := 0
	end := n - 1

	visited := make([]bool, n)
	route := make([]int, 0)
	totalDistance := 0.0

	current := start
	route = append(route, current)
	visited[current] = true

	for len(route) < n-1 {
		nearest := -1
		minDistance := math.MaxFloat64

		for i := range locations {
			if !visited[i] && i != end {
				dist, err := Haversine(locations[current].LocationLat, locations[current].LocationLong, locations[i].LocationLat, locations[i].LocationLong)
				if err != nil {
					return nil, 0, err
				}
				if dist < minDistance {
					minDistance = dist
					nearest = i
				}
			}
		}

		if nearest != -1 {
			route = append(route, nearest)
			visited[nearest] = true
			totalDistance += minDistance
			current = nearest
		}
	}

	// Add the end location
	route = append(route, end)
	finalLegDistance, err := Haversine(locations[current].LocationLat, locations[current].LocationLong, locations[end].LocationLat, locations[end].LocationLong)
	if err != nil {
		return nil, 0, err
	}
	totalDistance += finalLegDistance

	return route, totalDistance, nil
}

func EstimatedDeliveryTimeInMinutes(totalDistance float64) float64 {
	// Speed is given in km/h, converting to minutes
	timeInHours := totalDistance / deliverySpeed
	timeInMinutes := timeInHours * 60
	return timeInMinutes
}
