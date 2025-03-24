package main

import "fmt"

// empl interface
type Employee interface {
	CalculateSalary() int
}

// fulltime struct
type Fulltime struct {
	Daysworked int
}

// contractor
type Contractor struct {
	Daysworked int
}

// Frreelance
type Freelance struct {
	HoursWorked int
}

// meths to calcu salary
func (f Fulltime) CalculateSalary() int {
	return f.Daysworked * 500
}

func (c Contractor) CalculateSalary() int {
	return c.Daysworked * 100
}

func (fl Freelance) CalculateSalary() int {
	return fl.HoursWorked * 100
}

func main() {
	var choice int
	fmt.Println("Choose employee type: ")
	fmt.Println("1. Fulltime ")
	fmt.Println("2. Contractor ")
	fmt.Println("3. Freelance ")
	fmt.Scan(&choice)

	var salary int

	if choice == 1 {
		var days int
		fmt.Print("Enter days worked: ")
		fmt.Scan(&days)
		ft := Fulltime{Daysworked: days}
		salary = ft.CalculateSalary()
	} else if choice == 2 {
		var days int
		fmt.Print("Enter days worked: ")
		fmt.Scan(&days)
		ct := Contractor{Daysworked: days}
		salary = ct.CalculateSalary()
	} else if choice == 3 {
		var hours int
		fmt.Print("Enter hours worked: ")
		fmt.Scan(&hours)
		fl := Freelance{HoursWorked: hours}
		salary = fl.CalculateSalary()
	} else {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("Total salary: ", salary)

}
