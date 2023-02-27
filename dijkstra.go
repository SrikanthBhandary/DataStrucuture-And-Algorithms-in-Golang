package main

import (
	"fmt"
	"math"
)

type VisitedNodes struct {
	stack []*City
}

func NewVisitedNodes() *VisitedNodes {
	return &VisitedNodes{}
}

func (v *VisitedNodes) Nil() bool {
	if len(v.stack) > 0 {
		return true
	}
	return false
}
func (v *VisitedNodes) Append(item *City) {
	v.stack = append(v.stack, item)
}

func (v *VisitedNodes) Print() {
	for _, item := range v.stack {
		fmt.Println(item)
	}
}

func (v *VisitedNodes) HasAny(item *City) bool {
	for _, ite := range v.stack {
		if item == ite {
			return true
		}
	}
	return false
}

type City struct {
	Name   string
	Routes map[*City]float64
}

func NewCity(name string) *City {
	city := &City{Name: name}
	city.Routes = make(map[*City]float64)
	return city
}

func (c City) AddRoute(city *City, distance float64) {
	c.Routes[city] = distance
}

func dijkstra(startingCity *City, otherCities ...*City) {

	routesFromCity := make(map[*City]Result)
	routesFromCity[startingCity] = Result{distance: 0, city: startingCity}

	for _, adjecentCity := range otherCities {
		routesFromCity[adjecentCity] = Result{distance: math.Inf(1), city: nil}
	}

	visitedNodes := NewVisitedNodes()
	currentCity := startingCity
	for currentCity != nil {
		visitedNodes.Append(currentCity)
		fmt.Println("Current city", currentCity.Name)
		for city, distance := range currentCity.Routes {
			prevresult, _ := routesFromCity[currentCity] //Current City
			if result, ok := routesFromCity[city]; ok && (result.distance > distance+prevresult.distance) {
				routesFromCity[city] = Result{distance: distance + prevresult.distance, city: currentCity}

			}
		}
		currentCity = nil
		cheapestDistanceFromCurrentCity := math.Inf(1)
		for city, result := range routesFromCity {

			if result.distance < cheapestDistanceFromCurrentCity && !visitedNodes.HasAny(city) {
				cheapestDistanceFromCurrentCity = result.distance
				currentCity = city
			}
		}
	}

	for c, result := range routesFromCity {
		fmt.Println("City : ", c.Name, "Distance :", result.distance)
	}

}

type Result struct {
	distance float64
	city     *City
}

func main() {
	atlanta := NewCity("Atlanta")
	boston := NewCity("Boston")
	chicago := NewCity("Chicago")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston, 100)
	atlanta.AddRoute(denver, 160)
	boston.AddRoute(chicago, 120)
	boston.AddRoute(denver, 180)
	chicago.AddRoute(elPaso, 80)
	denver.AddRoute(chicago, 40)
	denver.AddRoute(elPaso, 140)
	dijkstra(atlanta /*starting city*/, boston, chicago, denver, elPaso /*others*/)
}
