package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/gommon/log"
)

func DoHttpGet(apiUrl string) (map[string]interface{}, error) {

	hc := http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	stringData, _ := ParseResponse(resp)
	log.Info(apiUrl, "：返回结果为：", stringData)
	return stringData, nil
}

func DoGet(apiUrl string) (map[string]interface{}, error) {

	hc := http.Client{}
	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	stringData, _ := ParseResponse(resp)
	log.Info(apiUrl, "：返回结果为：", stringData)
	return stringData, nil
}

func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	if response == nil {
		var result map[string]interface{}
		return result, nil
	}
	log.Info(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	stringBody := string(body)
	log.Info("response body=", stringBody)
	// Log.Info(stringBody)
	var result map[string]interface{}
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	return result, err
}

// httpPostForm
func HttpPostForm(apiUrl string, params url.Values) (map[string]interface{}, error) {
	resp, err := http.PostForm(apiUrl, params)
	// url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		// handle error
		return nil, err
	}

	defer resp.Body.Close()

	stringData, _ := ParseResponse(resp)
	log.Info(apiUrl, "：返回结果为：", stringData)
	return stringData, nil

}

func PostJsonWithHeader(httpposturl string, jsonStr interface{}, header map[string]string) (string, error) {

	fmt.Println("HTTP JSON POST URL:", httpposturl)

	jsonData, err := json.Marshal(jsonStr)
	if err != nil {
		return "", err
	}
	log.Info("json params========", string(jsonData))
	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	for k, v := range header {
		request.Header.Set(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do error ===", err.Error())
		return "", err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body==", string(body))

	stringBody := string(body)
	log.Info("response body=", stringBody)
	// Log.Info(stringBody)

	return stringBody, err
}

func PostJson(httpposturl string, jsonStr interface{}) (map[string]interface{}, error) {

	fmt.Println("HTTP JSON POST URL:", httpposturl)

	jsonData, err := json.Marshal(jsonStr)
	if err != nil {
		return nil, err
	}
	log.Info("json params========", string(jsonData))
	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do error ===", err.Error())
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body==", string(body))

	stringBody := string(body)
	log.Info("response body=", stringBody)
	// Log.Info(stringBody)
	var result map[string]interface{}
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}

func HttpPostForm301(apiUrl string, params url.Values) (string, error) {
	resp, err := http.PostForm(apiUrl, params)
	// url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	stringBody := string(body)
	log.Info(apiUrl, "：返回结果为：", stringBody)
	return stringBody, nil

}
