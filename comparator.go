package main

// By relevance
type InfoByScore []Info

func (u InfoByScore) Len() int {
	return len(u)
}
func (u InfoByScore) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u InfoByScore) Less(i, j int) bool {
	// reversed
	return u[i].RelevanceScore > u[j].RelevanceScore
}


// By views
type InfoByViews []Info

func (u InfoByViews) Len() int {
	return len(u)
}
func (u InfoByViews) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u InfoByViews) Less(i, j int) bool {
	// reversed
	return u[i].Views > u[j].Views
}