package data

import (
	"math/rand"
	"strconv"
)

type DummyData struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	Age     int               `json:"age"`
	Address map[string]string `json:"address"`
	Tags    []string          `json:"tags"`
}

func GenerateDummyData(count int) []DummyData {
	var data []DummyData
	for i := 1; i <= count; i++ {
		item := DummyData{
			ID:   i,
			Name: "User " + strconv.Itoa(i),
			Age:  rand.Intn(50) + 20, // age between 20 and 70
			Address: map[string]string{
				"city": "City" + strconv.Itoa(rand.Intn(5)),
				"zip":  strconv.Itoa(380000 + i),
			},
			Tags: []string{"golang", "benchmark", "json"},
		}
		data = append(data, item)
	}
	return data
}
