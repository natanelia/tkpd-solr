package solr

import (
	"fmt"
	"net/url"
	"testing"
)

func TestConnection(t *testing.T) {
	_, err1 := NewConnection("http://www.google.com", "core1")
	if err1 != nil {
		t.Errorf("ERROR: %s", err1.Error())
	}
}

func TestGetResponseMap(t *testing.T) {
	params := &url.Values{}
	params.Set("spellcheck.q", "samsun")

	c, err := NewConnection("http://192.168.100.129:8983/solr", "product_spellcheck")
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}
	resp, err2 := c.GetResponseMap("spell", params)
	if err2 != nil {
		t.Errorf("ERROR: %s", err2.Error())
	}

	resp = ConvertMapValueTypesToString(resp).(map[string]interface{})
	res := NewMapOfStrings(resp)
	fmt.Println(res.Get("spellcheck", "suggestions", "1", "suggestion", "0"))
}
