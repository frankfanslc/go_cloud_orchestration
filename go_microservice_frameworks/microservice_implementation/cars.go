package main

// Car type with Manufacturer, Model, ID
type Car struct {
	ID        		string `json:"id"`
	Manufacturer    string `json:"manufacturer"`
	Model      		string `json:"model"`
	Description 	string `json:"description,omitempty"`
}

var cars = map[string]Car{
	"0345391802": Car{Manufacturer: "Ford", Model: "Galaxy", ID: "0345391802"},
	"0000000000": Car{Manufacturer: "Porsche", Model: "Carrera", ID: "0000000000"},
}

// GetAllCars returns a slice of all cars
func GetAllCars() []Car {
	values := make([]Car, len(cars))
	idx := 0
	for _, car := range cars {
		values[idx] = car
		idx++
	}
	return values
}

// GetCar returns the car for a given ID
func GetCar(id string) (Car, bool) {
	car, found := cars[id]
	return car, found
}

// CreateCar creates a new Car if it does not exist
func CreateCar(car Car) (string, bool) {
	_, exists := cars[car.ID]
	if exists {
		return "", false
	}
	cars[car.ID] = car
	return car.ID, true
}

// UpdateCar updates an existing car
func UpdateCar(id string, car Car) bool {
	_, exists := cars[id]
	if exists {
		cars[id] = car
	}
	return exists
}

// DeleteCar removes a car from the map by ID key
func DeleteCar(id string) bool {
	delete(cars, id)
	return true
}
