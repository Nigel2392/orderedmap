package orderedmap_test

import (
	"encoding/json"
	"testing"

	"github.com/Nigel2392/orderedmap"
)

func TestMapJSON(t *testing.T) {
	var slice = make([]orderedmap.Map[string, string], 0)
	var m = orderedmap.New[string, string]()
	m.Set("a", "b")
	m.Set("c", "d")
	m.Set("e", "f")
	m.Set("g", "h")
	slice = append(slice, *m)
	m.Set("i", "j")
	m.Set("k", "l")
	m.Set("m", "n")
	m.Set("o", "p")
	slice = append(slice, *m)
	m.Set("q", "r")
	m.Set("s", "t")
	m.Set("u", "v")
	m.Set("w", "x")
	m.Set("y", "z")
	slice = append(slice, *m)

	var js, err = json.Marshal(slice)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(js))

	var slice2 = make([]orderedmap.Map[string, string], 0)
	err = json.Unmarshal(js, &slice2)
	if err != nil {
		t.Error(err)
	}

	for i, v := range slice2 {
		var aKeys, bKeys = v.Keys(), slice[i].Keys()
		if len(aKeys) != len(bKeys) {
			t.Errorf("slices are not equal, expected %v, got %v", len(bKeys), len(aKeys))
		} else if len(aKeys) == 0 {
			t.Errorf("slices are not equal, expected %v, got %v", len(bKeys), len(aKeys))
		} else if aKeys[0] != bKeys[0] {
			t.Errorf("slices are not equal, expected %v, got %v", bKeys[0], aKeys[0])
		} else if v.Get(aKeys[0]) != slice[i].Get(bKeys[0]) {
			t.Errorf("slices are not equal, expected %v, got %v", slice[i].Get(bKeys[0]), v.Get(aKeys[0]))
		}
	}
}
