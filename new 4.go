func countVotes(votes []string) {
	candidates := []string{"Анна", "Борис", "Виктор"}
	counts := make(map[string]int)

	for _, v := range votes {
		counts[v]++
	}

	total := len(votes)
	for _, c := range candidates {
		percent := 0.0
		if total > 0 {
			percent = float64(counts[c]) / float64(total) * 100
		}
		fmt.Printf("%s: %d голосов (%.2f%%)\n", c, counts[c], percent)
	}
}