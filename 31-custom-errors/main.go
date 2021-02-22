package main

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}

	return math.Pi * radius * radius, nil
}

func calculateArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}

	return math.Pi * radius * radius, nil
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func getCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func rectArea(length float64, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is negative"
	}
	if width < 0 {
		if err == "" {
			err += "width is negative"
		} else {
			err += ", width is negative"
		}
	}

	if err != "" {
		return 0, &rectAreaError{err, length, width}
	}

	return length * width, nil
}

type rectAreaError struct {
	err    string
	length float64
	width  float64
}

func (e *rectAreaError) Error() string {
	return e.err
}

func (e *rectAreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectAreaError) widthNegative() bool {
	return e.width < 0
}

func main() {
	// creating error
	radius := -13.4
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of circle %0.2f\n", area)
	}

	// add more information to error
	area1, err := calculateArea(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of circle %0.2f\n", area1)
	}

	// add more information using struct type and fields
	area2, err := getCircleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero.\n", err.radius)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Area of circle %0.2f\n", area2)
	}

	// add more information using methods on structs
	length, width := 12.2, -8.4
	area3, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			fmt.Println("Rect area error:", err)
			if err.lengthNegative() {
				fmt.Println("with length:", err.length)
			}
			if err.widthNegative() {
				fmt.Println("with width:", err.width)
			}
		} else {
			fmt.Println("Generic error:", err)
		}
	} else {
		fmt.Println("Rectangle area:", area3)
	}
}
