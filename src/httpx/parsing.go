package httpx

import (
	"fmt"
	"net/http"
	"net/url"
)

func ParseParams(baseUrl string, paramPairs []Pair) string {
	if len(paramPairs) == 0 {
		return baseUrl
	}

	params := url.Values{}
	for _, paramPair := range paramPairs {
		params.Add(paramPair.Key, paramPair.Value)
	}

	return fmt.Sprintf("%s?%s", baseUrl, params.Encode())
}

func ParseHeaders(request *http.Request, headerPairs []Pair) {
	if len(headerPairs) == 0 {
		return
	}

	for _, headerPair := range headerPairs {
		request.Header.Add(headerPair.Key, headerPair.Value)
	}
}
