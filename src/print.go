package main

import (
	"fmt"

	"github.com/mkmik/argsort"
	"github.com/pterm/pterm"
)

func printResult(proxies []string, scores []float64, delays []float64, stds []float64) {
	indexs := argsort.SortSlice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	table := [][]string{
		{"Name", "Score", "Delay", "Std"},
	}
	for _, index := range indexs {
		proxy := proxies[index]
		score := fmt.Sprintf("%.2f", scores[index])
		delay := fmt.Sprintf("%.2f", delays[index])
		std := fmt.Sprintf("%.2f", stds[index])

		table = append(table, []string{proxy, score, delay, std})
	}

	pterm.DefaultTable.WithData(table).Render()
}
