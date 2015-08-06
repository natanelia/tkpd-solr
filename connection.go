package solr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const VERSION = "0.1"

var userAgent = fmt.Sprintf("Tokopedia/%s (+https://github.com/vanng822/go-solr)", VERSION)

// HTTPPost make a POST request to path which also includes domain, headers are optional
func HTTPPost(path string, data *[]byte, headers [][]string, username, password string) ([]byte, error) {
	var (
		req *http.Request
		err error
	)

	client := &http.Client{}
	if data == nil {
		req, err = http.NewRequest("POST", path, nil)
	} else {
		req, err = http.NewRequest("POST", path, bytes.NewReader(*data))
	}

	if err != nil {
		return nil, err
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	if len(headers) > 0 {
		for i := range headers {
			req.Header.Add(headers[i][0], headers[i][1])
		}
	}
	return makeRequest(client, req)
}

// HTTPGet make a GET request to url, headers are optional
func HTTPGet(url string, headers [][]string, username, password string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	if len(headers) > 0 {
		for i := range headers {
			req.Header.Add(headers[i][0], headers[i][1])
		}
	}
	return makeRequest(client, req)
}

func makeRequest(client *http.Client, req *http.Request) ([]byte, error) {
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func bytes2json(data *[]byte) (map[string]interface{}, error) {
	var jsonData interface{}

	err := json.Unmarshal(*data, &jsonData)

	if err != nil {
		return nil, err
	}

	return jsonData.(map[string]interface{}), nil
}

func json2bytes(data interface{}) (*[]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (c *Connection) getResponseMap(handler string, params *url.Values) (map[string]interface{}, error) {
	params.Set("wt", "json")
	fmt.Println(fmt.Sprintf("%s/%s/%s?%s", c.url.String(), c.core, handler, params.Encode()))
	r, err := HTTPGet(fmt.Sprintf("%s/%s/%s?%s", c.url.String(), c.core, handler, params.Encode()), nil, c.username, c.password)
	if err != nil {
		return nil, err
	}
	resp, err := bytes2json(&r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type mapOfStrings map[string]interface{}

func NewMapOfStrings(m map[string]interface{}) mapOfStrings {
	mos := make(mapOfStrings)
	for k, _ := range m {
		mos[k] = m[k]
	}
	return mos
}

func (m mapOfStrings) Get(params ...string) string {
	for _, v := range params {
		switch vv := m[v].(type) {
		case string:
			//fmt.Println("STRING")
			return vv
		case map[string]interface{}:
			//fmt.Println("MAP")
			a := NewMapOfStrings(vv)
			return a.Get(params[1:]...)
		case int:
			//fmt.Println("INT")
			return strconv.Itoa(vv)
		case []interface{}:
			a, _ := strconv.Atoi(params[1])
			switch xx := vv[a].(type) {
			case string:
				return xx
			case map[string]interface{}:
				b := NewMapOfStrings(xx)
				return b.Get(params[2:]...)
			case int:
				return strconv.Itoa(xx)

			}
		case interface{}:
			//return "UNKNOWN"
		}
	}
	return ""
}

func convertMapValueTypesToString(d interface{}) interface{} {
	switch dd := d.(type) {
	case string:
		// fmt.Println("STRING", dd)
		return dd

	case int:
		// fmt.Println("INT", strconv.Itoa(dd))
		return strconv.Itoa(dd)

	case float64:
		// fmt.Println("FLOAT", strconv.FormatFloat(dd, 'f', 6, 64))
		return strconv.FormatFloat(dd, 'f', 6, 64)

	case map[string]interface{}:
		// fmt.Println("map string interface")
		for k, v := range dd {
			dd[k] = convertMapValueTypesToString(v)
		}
		return dd

	case []interface{}:
		// fmt.Println("[]interface{}")
		for i, v := range dd {
			dd[i] = convertMapValueTypesToString(v)
		}
		return dd
	}

	return d
}

func PrintMapInterface(d map[string]interface{}) {
	for k, v := range d {
		switch vv := v.(type) {
		case string:
			fmt.Println(fmt.Sprintf("STRING %s: %s", k, v))
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float", vv)
		case map[string]interface{}:
			fmt.Print(k, ": {")
			PrintMapInterface(vv)
			fmt.Println("}")
		case []interface{}:
			fmt.Print(k, ": [")
			for i, u := range vv {
				switch uu := u.(type) {
				case map[string]interface{}:
					PrintMapInterface(uu)
				default:
					fmt.Println(i, u)
				}
			}
			fmt.Println("]")
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv)
		}
	}
}

type UrlParams url.Values

/*func hasError(response map[string]interface{}) (string, bool) {
	errorMap, ok := response["error"].(map[string]interface{})
	if !ok {
		return "", ok
	}

	errorMessage, ok := errorMap["msg"].(string)
	if !ok {
		return "Unknown error", ok
	}

	return errorMessage, ok
}

func successStatus(response map[string]interface{}) bool {
	responseHeader, ok := response["responseHeader"].(map[string]interface{})
	if !ok {
		return false
	}

	if status, ok := responseHeader["status"].(float64); ok {
		return 0 == int(status)
	}

	return false
}*/

type Connection struct {
	url      *url.URL
	core     string
	username string
	password string
}

// NewConnection will parse solrUrl and return a connection object, solrUrl must be a absolute url or path
func NewConnection(solrUrl, core string) (*Connection, error) {
	u, err := url.ParseRequestURI(strings.TrimRight(solrUrl, "/"))
	if err != nil {
		return nil, err
	}

	return &Connection{url: u, core: core}, nil
}

// Set to a new core
func (c *Connection) SetCore(core string) {
	c.core = core
}

func (c *Connection) SetBasicAuth(username, password string) {
	c.username = username
	c.password = password
}
