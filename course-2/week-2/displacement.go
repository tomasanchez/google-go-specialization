package main

// Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo,
// and initial displacement so.
// s = 1/2 a t^2 + vo t + so

// Write a program which first prompts the user
// to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the program should compute the displacement
// after the entered time.

// You will need to define and use a function called GenDisplaceFn() which takes three float64
// arguments, acceleration a, initial velocity vo, and initial displacement so.

// GenDisplaceFn() should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
// float64 argument which is the displacement travelled after time t.

import "fmt"

func EnterValue(value *float64, name string) {
	fmt.Printf("Enter value for %s:\n", name)
	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {

	var a, vo, so, t float64

	EnterValue(&a, "acceleration")
	EnterValue(&vo, "initial velocity")
	EnterValue(&so, "initial displacement")
	EnterValue(&t, "time")

	fn := GenDisplaceFn(a, vo, so)
	fmt.Printf("s(t=%.2f)= %.2f\n", t, fn(t))
}
