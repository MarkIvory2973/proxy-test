package httpx

import (
	"io"
	"net/http"
)

func Get(baseUrl string, paramPairs []Pair, headerPairs []Pair) (Response, error) {
	url := ParseParams(baseUrl, paramPairs)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, err
	}
	ParseHeaders(request, headerPairs)

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return Response{}, err
	}
	defer response.Body.Close()

	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{Status: response.StatusCode, Raw: raw}, nil
}
