import "time"

type LogEntry struct {
	IP         string
	StatusCode int
	Timestamp  time.Time
}

func filterErrors(logs []LogEntry) []LogEntry {
	var result []LogEntry
	for _, l := range logs {
		if l.StatusCode >= 400 && l.StatusCode < 600 {
			result = append(result, l)
		}
	}
	return result
}