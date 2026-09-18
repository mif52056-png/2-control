func main() {
	expenses := map[string]float64{
		"Еда":         15000,
		"Транспорт":   5000,
		"Развлечения": 3000,
	}

	expenses["Еда"] += 2000

	var total float64
	for _, sum := range expenses {
		total += sum
	}
	fmt.Println(total) // 25000
}