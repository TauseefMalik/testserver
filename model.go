package main

type Info struct {
	Url            string `json:"url,omitempty"`
	Views          float64    `json:"views,omitempty"`
	RelevanceScore float64    `json:"relevance_score,omitempty"`
}

type Response struct {
	Data  []Info `json:"data,omitempty"`
	Count int    `json:"count,omitempty"`
}
