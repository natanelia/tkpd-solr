package solr

import (
	"net/url"
	"testing"
)

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

	if len(resp.GetMapToString("responseHeader")) == 0 {
		t.Errorf("ERROR: GetMapToString map[string]string failed")
	}

	if len(resp.GetMapToInterface("spellcheck")) == 0 {
		t.Errorf("ERROR: GetMapToInterface map[string]interface{} failed")
	}

	if len(resp.GetArrayOfInterface("spellcheck", "suggestions")) == 0 {
		t.Errorf("ERROR: Get []interface{} failed")
	}

	resp.GetString()
	resp.GetMapToString()
	resp.GetMapToInterface()
	resp.GetArrayOfInterface()
}
