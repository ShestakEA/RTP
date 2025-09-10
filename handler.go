package main

import (
	"encoding/json"
	"math"
	"math/rand/v2"
	"net/http"
)

type Response struct {
	Result float64 `json:"result"`
}

func NewRandomHandler(router *http.ServeMux, rtp float64) {
	step := 50
	N := 10000
	expcoeff := rtp / 50_000
	arr := make([]float64, N/step+1)
	prob := make([]float64, N/step+1)
	arr[0] = 1
	prob[0] = 1
	sum := prob[0]
	for i := range N / step {
		arr[i+1] = float64(50 * (i + 1))
		prob[0] = 1 / math.Pow(arr[i+1], expcoeff)
		sum += prob[i+1]
	}

	for i := range len(prob) {
		prob[i] = prob[i] / sum
	}

	router.HandleFunc("GET /get", func(w http.ResponseWriter, req *http.Request) {
		randNum := rand.Float64()
		isFind := false
		cumulativeProbability := 0.0
		value := 0.0

		for i, multiplier := range arr {
			cumulativeProbability += prob[i]
			if randNum < cumulativeProbability {
				value = multiplier
				isFind = true
				break
			}
		}

		if !isFind {
			value = arr[len(arr)-1]
		}

		res := Response{value}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	})
}
