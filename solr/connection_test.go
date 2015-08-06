package solr

import (
	"testing"
)

func TestConnection(t *testing.T) {
	_, err1 := NewConnection("http://www.google.com", "core1")
	if err1 != nil {
		t.Errorf("ERROR: %s", err1.Error())
	}
}
