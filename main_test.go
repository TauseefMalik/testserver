package main

import (
	_ "fmt"
	"reflect"
	_"strings"
	"testing"

)


func TestSortWebSitesViews(t *testing.T) {
	var a = []Info{
		{"www.yahoo.com/abc7", 7000, 0.7},
		{"www.yahoo.com/abc6", 6000, 0.6},
		{"www.wikipedia.com/abc1", 11000, 0.1},
	}
	var want = []Info{
		{"www.wikipedia.com/abc1", 11000, 0.1},
		{"www.yahoo.com/abc7", 7000, 0.7},
		{"www.yahoo.com/abc6", 6000, 0.6},
	}

	key := "views"
	t.Run(key, func(t *testing.T) {
		SortWebSites(key, a)
		if !reflect.DeepEqual(a, want) {
			t.Errorf("got %v, want %v", a, want)
		}
	})
}

func TestSortWebSitesScore(t *testing.T) {
	var a = []Info{
		{"www.wikipedia.com/abc1", 11000, 0.1},
		{"www.yahoo.com/abc7", 7000, 0.7},
		{"www.yahoo.com/abc6", 6000, 0.6},
	}
	var want = []Info{
		{"www.yahoo.com/abc7", 7000, 0.7},
		{"www.yahoo.com/abc6", 6000, 0.6},
		{"www.wikipedia.com/abc1", 11000, 0.1},
	}

	key := "relevanceScore"
	t.Run(key, func(t *testing.T) {
		SortWebSites(key, a)
		if !reflect.DeepEqual(a, want) {
			t.Errorf("got %v, want %v", a, want)
		}
	})
}
