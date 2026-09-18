func uniqueTags(posts [][]string) map[string]bool {
	tags := make(map[string]bool)
	for _, post := range posts {
		for _, tag := range post {
			tags[tag] = true
		}
	}
	return tags
}