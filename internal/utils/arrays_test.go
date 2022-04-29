package utils

import "testing"

func TestForEachSlice(t *testing.T) {
	entries := []string{"a", "b", "c", "d", "e"}
	var matcher []string
	ForEachSlice(entries, func(entry string) {
		matcher = append(matcher, entry)
	})

	for i := range matcher {
		if entries[i] != matcher[i] {
			t.Errorf("entry are in wrong position, Expect %s, Actual: %s", entries[i], matcher[i])
		}
	}

	if len(entries) != len(matcher) {
		t.Errorf("some entries are not a same size, Expect: %s, Actual: %s", entries, matcher)
	}
}

func TestFilterSlice(t *testing.T) {
	entries := []string{"a", "b", "c", "d", "e"}
	var matcher = FilterSlice(entries, func(entry string) bool {
		return entry == "a" || entry == "c"
	})

	if matcher[0] != "a" {
		t.Errorf("first entry is not 'a', Actual: %s", matcher[0])
	}

	if matcher[1] != "c" {
		t.Errorf("second entry is not 'c', Actual: %s", matcher[1])
	}

	if len(matcher) != 2 {
		t.Errorf("matched entry are not a requested size, Expect: 2, Actual: %d", len(matcher))
	}
}

func TestAnySlice(t *testing.T) {
	entries := []string{"a", "b", "c", "d", "e"}
	var matcher = AnySlice(entries, func(entry string) bool {
		return entry == "a" || entry == "c"
	})

	if !matcher {
		t.Error("request is incorrect, Expect: true, Actual: false")
	}

	matcher = AnySlice(entries, func(entry string) bool {
		return entry == "f"
	})

	if matcher {
		t.Error("request is incorrect, Expect: false, Actual: true")
	}
}

func TestNoneSlice(t *testing.T) {

}

func TestMapSlice(t *testing.T) {

}

func TestFlatMapSlice(t *testing.T) {

}
