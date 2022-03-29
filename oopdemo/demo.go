package oopdemo

type employee struct {
	Name          string
	Age           int
	StreetAddress string
	PostalCode    int
	City          string
}

type IPayable interface {
	CalculateSalary() int
}

type hourlyEmployee struct {
	employee
	HourlySalary int // 85
}

// OOP Methods
func (emp hourlyEmployee) CalculateSalary() int {
	return emp.HourlySalary * 20
}

type monthlyEmployee struct {
	employee
	MonthlySalary int // 10000
}

func (emp monthlyEmployee) CalculateSalary() int {
	return emp.MonthlySalary
}

func New(name string, age int, street string, postal int, city string) employee {
	return employee{Name: name, Age: age, StreetAddress: street, PostalCode: postal,
		City: city}
}

func NewHourly(name string, age int, street string, postal int, city string, hourlySalary int) hourlyEmployee {
	return hourlyEmployee{employee: employee{Name: name, Age: age, StreetAddress: street, PostalCode: postal,
		City: city}, HourlySalary: hourlySalary}
}

func NewMonthly(name string, age int, street string, postal int, city string, monthlySalary int) monthlyEmployee {
	return monthlyEmployee{employee: employee{Name: name, Age: age, StreetAddress: street, PostalCode: postal,
		City: city}, MonthlySalary: monthlySalary}
}
