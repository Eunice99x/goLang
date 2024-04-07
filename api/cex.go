package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"younes.dev/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurr := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl,upCurr)) 

	if err != nil {
		return nil, err
	}
	
	var response CexRes
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	}else {
		return nil, fmt.Errorf("status code received: %v",res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}