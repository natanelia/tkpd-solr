package solr

import (
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

	c, err := NewConnection("http://192.168.100.129:8983", "product_spellcheck")
	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}
	resp, err2 := c.GetResponseMap("spell", params)
	if err2 != nil {
		t.Errorf("ERROR: %s", err2.Error())
	}

	if resp.GetString("spellcheck", "suggestions", "1", "suggestion", "0") == "" {
		t.Errorf("ERROR: GetString failed")
	}

	if len(resp.GetMapToString("spellcheck")) == 0 {
		t.Errorf("ERROR: GetMapToString failed")
	}

	if len(resp.GetMapToInterface("spellcheck")) == 0 {
		t.Errorf("ERROR: GetMapToInterface failed")
	}
}
