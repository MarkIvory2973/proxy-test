package main

import (
	"proxy-test/mihomo"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var api mihomo.API

func test() {
	results := make(map[string][]float64)
	for index := range num {
		logger.Infof("testing delays of '%s' (%d/%d)", group, index+1, num)

		result, err := mihomo.GroupDelay(api, group, url, timeout)
		if err != nil {
			logger.Fatal(err)
		}
		for proxy, delay := range result {
			results[proxy] = append(results[proxy], delay)
		}

		logger.Infof("waiting for %ds", interval)

		time.Sleep(time.Duration(interval) * time.Second)
	}

	logger.Infof("tested delays of '%s'", group)
	logger.Infof("calculating scores of '%s'", group)

	proxies, scores, delays, stds := calcScores(results, weight)

	logger.Infof("calculated scores of '%s'", group)

	printResult(proxies, scores, delays, stds)
}

func main() {
	logger = initLogger()
	initCliCmd()
}
