package httpx

import "encoding/json"

type Pair struct {
	Key   string
	Value string
}

type Response struct {
	Status int
	Raw    []byte
}

func (response Response) ToJson() (map[string]interface{}, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal(response.Raw, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
