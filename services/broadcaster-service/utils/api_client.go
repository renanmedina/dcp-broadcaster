package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/renanmedina/dcp-broadcaster/internal/exceptions"
)

type ApiClient[T any] struct {
	httpClient http.Client
	baseUrl    string
	authToken  string
	logger     *ApplicationLogger
}

type ApiConfig struct {
	ApiUrl     string
	AuthToken  string
	LogEnabled bool
}

func NewApiClient[T any](config ApiConfig) ApiClient[T] {
	var logger *ApplicationLogger

	if config.LogEnabled {
		logger = GetApplicationLogger()
	}

	return ApiClient[T]{
		httpClient: http.Client{},
		baseUrl:    config.ApiUrl,
		authToken:  config.AuthToken,
		logger:     logger,
	}
}

func (c *ApiClient[T]) BuildUrl(path string, params map[string]interface{}, requestMethod string) string {
	url := fmt.Sprintf("%s%s", c.baseUrl, path)
	if requestMethod != http.MethodGet {
		return url
	}

	var paramsStrings []string

	for pname, pvalue := range params {
		paramsStrings = append(paramsStrings, fmt.Sprintf("%s=%s", pname, pvalue))
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(paramsStrings, "&"))
}

func parseResult[T any](data []byte) (*T, error) {
	var resultData T
	err := json.Unmarshal(data, &resultData)

	if err != nil {
		return nil, err
	}

	return &resultData, nil
}

func (client *ApiClient[T]) Get(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodGet, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) Post(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodPost, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) Put(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodPut, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) parseResponse(response *http.Response, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	parsed, err := parseResult[T](bodyData)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (client *ApiClient[T]) performRequest(requestMethod string, path string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	url := client.BuildUrl(path, params, requestMethod)
	paramsBuffer := bytes.NewBuffer(make([]byte, 0))

	if len(params) > 0 {
		bodyParams, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		paramsBuffer = bytes.NewBuffer(bodyParams)
	}

	request, err := http.NewRequest(requestMethod, url, paramsBuffer)

	if err != nil {
		return nil, err
	}

	headers["Accept"] = "*/*"
	headers["Content-Type"] = "application/json"

	if client.authToken != "" {
		headers["Authorization"] = client.authToken
	}

	for headerKey, headerValue := range headers {
		request.Header.Add(headerKey, headerValue)
	}

	client.log(fmt.Sprintf("[%s] Sending http request to %s", requestMethod, url))
	response, err := client.httpClient.Do(request)
	if response.StatusCode == http.StatusUnprocessableEntity {
		return response, exceptions.NewHttpResponseError(response.StatusCode, response.Status)
	}

	client.log(fmt.Sprintf("Response Status: %s", response.Status))
	client.log(fmt.Sprintf("Response StatusCode: %d", response.StatusCode))

	return response, err
}

func (client *ApiClient[T]) log(message string) {
	if client.logger != nil {
		typeName := reflect.TypeFor[T]().Name()
		formattedLog := fmt.Sprintf("[%s] %s", typeName, message)
		client.logger.Info(formattedLog)
	}
}
