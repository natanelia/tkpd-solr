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

	c, err := solr.NewConnection("http://192.168.100.16:8983", "product_spellcheck")
	if err != nil {
		fmt.Println("error")
	}
	resp, err2 := c.GetResponseMap("spell", params)
	if err2 != nil {
		fmt.Println("error")
	}

	fmt.Println(reflect.TypeOf(resp.GetString("spellcheck", "suggestions", "1", "suggestion", "0")))
	fmt.Println(reflect.TypeOf(resp.GetMapToString("responseHeader", "0")))
	fmt.Println((resp.GetMapToString("responseHeader", "0")))
	fmt.Println(reflect.TypeOf(resp.GetMapToInterface("spellcheck")))
	fmt.Println((resp.GetMapToInterface("spellcheck", "0")))
	fmt.Println(reflect.TypeOf(resp.GetArrayOfInterface("spellcheck", "suggestions")))
	fmt.Println((resp.GetArrayOfInterface("spellcheck", "0")))
}
