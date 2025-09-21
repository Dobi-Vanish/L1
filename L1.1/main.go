package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name    string
	Age     int
	Country string
	Sex     string
}

type Action struct {
	Human      // embedded struct
	ActionType string
	Timestamp  time.Time
}

func main() {
	action := Action{
		Human: Human{
			Name:    "Alex",
			Age:     25,
			Country: "Australia",
			Sex:     "Male",
		},
		ActionType: "creating example",
	}

	action.Introduce()
	action.Birthday()
	action.Perform()

	fmt.Printf("Age: %d\n", action.GetAge())
	fmt.Printf("Adult: %t\n", action.IsAdult())
	fmt.Printf("Name: %s, Country: %s\n", action.Name, action.Country)
	fmt.Println(action.GetActionInfo())
}

func (a *Action) Perform() {
	fmt.Printf("Performing action: %s\n", a.ActionType)
	a.Timestamp = time.Now()
}

func (a *Action) GetActionInfo() string {
	return fmt.Sprintf("Action '%s' completed %s by %s",
		a.ActionType, a.Timestamp.Format("2006-01-02 15:04:05"), a.Name)
}

func (h *Human) Introduce() {
	fmt.Printf("Hola, my name is %s, I'm %d years old. I'm from %s., I'm %s.\n", h.Name, h.Age, h.Country, h.Sex)
}

func (h *Human) Birthday() {
	h.Age++
	fmt.Printf("Happy Birthday! Nom I'm %d.\n", h.Age)
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) IsAdult() bool {
	return h.Age >= 18
}
