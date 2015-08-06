package main

import (
	"fmt"
	"github.com/natanelia/tkpd-solr/solr"
	"net/url"
	"reflect"
)

func main() {
	params := &url.Values{}
	params.Set("spellcheck.q", "samsun")

	c, err := solr.NewConnection("http://192.168.100.129:8983", "product_spellcheck")
	if err != nil {
		fmt.Println("error")
	}
	resp, err2 := c.GetResponseMap("spell", params)
	if err2 != nil {
		fmt.Println("error")
	}

	fmt.Println(reflect.TypeOf(resp.GetString("spellcheck", "suggestions", "1", "suggestion", "0")))
	fmt.Println(reflect.TypeOf(resp.GetMapToString("responseHeader")))
	fmt.Println(reflect.TypeOf(resp.GetMapToInterface("spellcheck")))
	fmt.Println(reflect.TypeOf(resp.GetArrayOfInterface("spellcheck", "suggestions")))
}
