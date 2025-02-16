package main

import (
	"fmt"
	"magicmaps/magic"
	"testing"
)

func TestMaps(t *testing.T) {
	m := map[string]any{}
	m["foo"] = 123
	mm := magic.NewMap(m)

	fmt.Println(mm)
}
