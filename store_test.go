package pectx_test

import (
	"testing"

	"github.com/pragmaticengineering/pectx"
)

// helper function to create a new Store
func newStore(key, value string, t *testing.T) *pectx.Store {
	t.Helper()
	f := &pectx.Store{}
	f.Set(key, value)
	return f
}

func TestGet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *pectx.Store
		key      string
		exists   bool
		val      string
	}{
		{
			desc:     "validate Store values",
			expected: newStore("key", "val", t),
			key:      "key",
			exists:   true,
			val:      "val",
		},
		{
			desc:     "validate that the value doesn't exist",
			expected: newStore("key", "", t),
			key:      "key",
			exists:   false,
			val:      "",
		},
		//{
		//	desc: "validate that the value is overwritten",
		//	expected: newStore("key", "val"),
		//	key: 	"key",
		//	exists: true,
		//	val: "val2",
		//},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			value := testCase.expected.Get(testCase.key)
			if value != testCase.val {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.val, value)
			}

		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *pectx.Store
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is set",
			expected: newStore("key", "val", t),
			key:      "key",
			value:    "val",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := pectx.Store{}
			f.Set(testCase.key, testCase.value)

			// Check if the value is set
			value := testCase.expected.Get(testCase.key)

			if value != testCase.value {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.value, value)
			}

		})
	}
}

func TestOverrideValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *pectx.Store
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is overwritten",
			expected: newStore("key", "val2", t),
			key:      "key",
			value:    "val2",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := pectx.Store{}
			f.Set(testCase.key, "val")
			f.Set(testCase.key, testCase.value)

			// Check if the value is set
			value := testCase.expected.Get(testCase.key)
			if value != testCase.value {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.value, value)
			}

		})
	}
}
