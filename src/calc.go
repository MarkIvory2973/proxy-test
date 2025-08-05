package main

import (
	"proxy-test/mathx"
)

func calcScores(results map[string][]float64, weight float64) ([]string, []float64, []float64, []float64) {
	var proxies []string
	var delays [][]float64
	for proxy, delay := range results {
		proxies = append(proxies, proxy)
		delays = append(delays, delay)
	}

	var percentiles, stds []float64
	for _, delay := range delays {
		percentile := mathx.Percentile(delay, 0.75)
		std := mathx.Std(delay, 0)

		percentiles = append(percentiles, percentile)
		stds = append(stds, std)
	}

	xPercentiles, xStds := mathx.Norm(percentiles, 1e-2), mathx.Norm(stds, 1e-2)

	var scores []float64
	for index := range len(delays) {
		score := weight*xPercentiles[index] + (1-weight)*xStds[index]

		scores = append(scores, score)
	}

	return proxies, scores, percentiles, stds
}
