package mihomo

import (
	"proxy-test/httpx"
	"fmt"
	"strconv"
)

func GroupDelay(api API, group string, url string, timeout int) (map[string]float64, error) {
	baseUrl := fmt.Sprintf("%s/group/%s/delay", api.ParseBaseUrl(), group)
	paramPairs := []httpx.Pair{
		{Key: "url", Value: url},
		{Key: "timeout", Value: strconv.Itoa(timeout)},
	}
	headerPairs := api.ParseHeaderPairs()

	response, err := httpx.Get(baseUrl, paramPairs, headerPairs)
	if err != nil {
		return nil, err
	}
	if response.Status != 200 {
		return nil, fmt.Errorf("failed to fetch delays from '%s' (status %d)\n%s", baseUrl, response.Status, response.Raw)
	}

	parsed, err := response.ToJson()
	if err != nil {
		return nil, err
	}

	results := make(map[string]float64)
	for proxy, delay := range parsed {
		delay, ok := delay.(float64)
		if !ok {
			return nil, fmt.Errorf("failed to process delays")
		}
		results[proxy] = delay
	}

	return results, nil
}
