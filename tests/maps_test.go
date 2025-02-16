package main

import (
	"encoding/json"
	"fmt"
	"magicmaps/magic"
	"testing"
)

var sampleJson = []byte(`{"foo": 123}`)

func TestMaps(t *testing.T) {
	var m map[string]any
	json.Unmarshal(sampleJson, &m)
	mm := magic.NewMap(m)
	val := mm.GetFloatAsInt("foo")

	fmt.Println(val)
}
