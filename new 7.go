type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func payroll(employees []Employee) (float64, float64) {
	var total float64
	for _, e := range employees {
		total += e.Salary
	}
	avg := 0.0
	if len(employees) > 0 {
		avg = total / float64(len(employees))
	}
	return total, avg
}