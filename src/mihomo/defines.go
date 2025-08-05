package mihomo

import (
	"proxy-test/httpx"
	"fmt"
	"strconv"
)

type API struct {
	TLS    bool
	Addr   string
	Port   int
	Secret string
}

func (api API) ParseBaseUrl() string {
	var baseUrl string
	if api.TLS {
		baseUrl = fmt.Sprintf("https://%s:%s", api.Addr, strconv.Itoa(api.Port))
	} else {
		baseUrl = fmt.Sprintf("http://%s:%s", api.Addr, strconv.Itoa(api.Port))
	}

	return baseUrl
}

func (api API) ParseHeaderPairs() []httpx.Pair {
	if api.Secret == "" {
		return []httpx.Pair{}
	}

	headerPairs := []httpx.Pair{
		{Key: "Authorization", Value: fmt.Sprintf("Bearer %s", api.Secret)},
	}

	return headerPairs
}
