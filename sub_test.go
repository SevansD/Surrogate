package surrogate

import (
	"fmt"
	"strconv"
	"testing"
)

var testSub = []struct {
	s string
	m map[string]string
	r []string
}{
	{
		"a",
		map[string]string{"a": "b"},
		[]string{"a", "b"},
	},
	{
		"ab-cd",
		map[string]string{"a": "x", "c": "y"},
		[]string{"ab-cd", "ab-yd", "xb-cd", "xb-yd"},
	},
	{
		"ab-cd-ef-gh",
		map[string]string{"a": "x", "c": "y", "e": "b"},
		[]string{"ab-cd-ef-gh", "ab-cd-bf-gh", "ab-yd-ef-gh", "ab-yd-bf-gh", "xb-cd-ef-gh", "xb-cd-bf-gh", "xb-yd-ef-gh", "xb-yd-bf-gh"},
	},
	{
		"aA-bB",
		map[string]string{"a": "x"},
		[]string{"aa-bb", "ax-bb", "xa-bb", "xx-bb"},
	},
}

var testSubMM = []struct {
	s string
	m map[string][]string
	r []string
}{
	{
		"a",
		map[string][]string{
			"a": {"b", "c"},
		},
		[]string{"a", "b", "c"},
	},
	{
		"ab-cd",
		map[string][]string{
			"a": {"x", "z", "o", "m"}, "c": {"y"},
		},
		[]string{"ab-cd", "ab-yd", "xb-cd", "xb-yd", "zb-cd", "zb-yd", "ob-cd", "ob-yd", "mb-cd", "mb-yd"},
	},
	{
		"ab-cd-ef-gh",
		map[string][]string{"a": {"x"}, "c": {"y"}, "e": {"b"}},
		[]string{"ab-cd-ef-gh", "ab-cd-bf-gh", "ab-yd-ef-gh", "ab-yd-bf-gh", "xb-cd-ef-gh", "xb-cd-bf-gh", "xb-yd-ef-gh", "xb-yd-bf-gh"},
	},
	{
		"aA-bB",
		map[string][]string{"a": {"x", "y"}},
		[]string{"aa-bb", "ax-bb", "xa-bb", "xx-bb", "ya-bb", "yx-bb", "ay-bb", "xy-bb", "yy-bb"},
	},
}

var testSubCaseSensitive = []struct {
	s string
	m map[string]string
	r []string
}{
	{
		"a",
		map[string]string{"a": "b"},
		[]string{"a", "b"},
	},
	{
		"ab-cd",
		map[string]string{"a": "x", "c": "y"},
		[]string{"ab-cd", "ab-yd", "xb-cd", "xb-yd"},
	},
	{
		"ab-cd-ef-gh",
		map[string]string{"a": "x", "c": "y", "e": "b"},
		[]string{"ab-cd-ef-gh", "ab-cd-bf-gh", "ab-yd-ef-gh", "ab-yd-bf-gh", "xb-cd-ef-gh", "xb-cd-bf-gh", "xb-yd-ef-gh", "xb-yd-bf-gh"},
	},
	{
		"aA-bB",
		map[string]string{"a": "x", "A": "C"},
		[]string{"aA-bB", "aC-bB", "xA-bB", "xC-bB"},
	},
	{
		"aA-bB-12",
		map[string]string{"a": "x", "A": "C", "1": "99"},
		[]string{"aA-bB-12", "aA-bB-992", "aC-bB-12", "aC-bB-992", "xA-bB-12", "xA-bB-992", "xC-bB-12", "xC-bB-992"},
	},
}

func TestSub(t *testing.T) {
	for i, ts := range testSub {
		t.Run("Sub-"+strconv.Itoa(i), func(t *testing.T) {
			result := Sub(ts.s, ts.m)
			if len(result) != len(ts.r) {
				fmt.Println(result)
				t.Fatalf("Len of result should be %d elements got %d", len(ts.r), len(result))
			}
			if !equal(result, ts.r) {
				t.Fatalf("Result isn't right. Expected: %s. Actual: %s", ts.r, result)
			}
		})
	}
}

func TestSubMultiMap(t *testing.T) {
	for i, ts := range testSubMM {
		t.Run("SubMultiMap-"+strconv.Itoa(i), func(t *testing.T) {
			result := SubMultiMap(ts.s, ts.m)
			if len(result) != len(ts.r) {
				fmt.Println(result)
				t.Fatalf("Len of result should be %d elements got %d", len(ts.r), len(result))
			}
			if !equal(result, ts.r) {
				t.Fatalf("Result isn't right. Expected: %s. Actual: %s", ts.r, result)
			}
		})
	}
}

func TestSubCaseSensitive(t *testing.T) {
	for i, ts := range testSubCaseSensitive {
		t.Run("Sub-Case-Sensitive-"+strconv.Itoa(i), func(t *testing.T) {
			result := SubCaseSensitive(ts.s, ts.m)
			if len(result) != len(ts.r) {
				fmt.Println(result)
				t.Fatalf("Len of result should be %d elements got %d", len(ts.r), len(result))
			}
			if !equal(result, ts.r) {
				t.Fatalf("Result isn't right. Expected: %s. Actual: %s", ts.r, result)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
