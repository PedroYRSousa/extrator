package auth

import (
	"encoding/json"
	"errors"
	"extrator/internal/http_request"
	"extrator/internal/utils"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/PaesslerAG/jsonpath"
)

func (dd *S_DynamicDetails) getToken() (*string, error) {
	request := http.Request{}

	uri, err := url.Parse(dd.Endpoint)
	if err != nil {
		return nil, err
	}
	request.URL = uri

	err = dd.EndpointConfig.Mount(&request)
	if err != nil {
		return nil, err
	}

	token := ""
	err = utils.ExecWithAttempts(int(*dd.EndpointConfig.Retry.Attempts), int(*dd.EndpointConfig.Retry.DelayInSeconds), func() error {
		res, err := http_request.MakeRequest(&request, *dd.EndpointConfig)
		if err != nil {
			if errors.Is(err, http_request.ErrStatusCodeWaiting) {
				time.Sleep(time.Duration(*dd.EndpointConfig.WaitingTimeInSeconds) * time.Second)
				return err
			} else if errors.Is(err, http_request.ErrStatusCodeSkip) {
				return nil
			} else if errors.Is(err, http_request.ErrStatusCodeError) {
				time.Sleep(time.Duration(*dd.EndpointConfig.WaitingTimeErrorInSeconds) * time.Second)
				return err
			}
		}
		defer res.Body.Close()

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// TODO, lidar com outros tipos de resposta
		var body interface{}
		err = json.Unmarshal(bodyData, &body)
		if err != nil {
			return err
		}

		tokenInterface, err := jsonpath.Get(dd.Extract, body)
		if err != nil {
			return err
		}

		tokenStr, ok := tokenInterface.(string)
		if !ok {
			return fmt.Errorf("failed to assert token to string")
		}
		token = tokenStr

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}
