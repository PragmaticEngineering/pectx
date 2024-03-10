package pectx_test

import (
	"testing"

	"github.com/pragmaticengineering/pectx"
)

// helper function to create a new Field
func newField(key, value string, t *testing.T) *pectx.Field {
	t.Helper()
	f := &pectx.Field{}
	f.Set(key, value)
	return f
}

func TestGet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *pectx.Field
		key      string
		exists   bool
		val      string
	}{
		{
			desc:     "validate Field values",
			expected: newField("key", "val", t),
			key:      "key",
			exists:   true,
			val:      "val",
		},
		{
			desc:     "validate that the value doesn't exist",
			expected: newField("key", "", t),
			key:      "key",
			exists:   false,
			val:      "",
		},
		//{
		//	desc: "validate that the value is overwritten",
		//	expected: newField("key", "val"),
		//	key: 	"key",
		//	exists: true,
		//	val: "val2",
		//},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			key, value := testCase.expected.Get()
			if key != testCase.key {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.key, key)
			}
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
		expected *pectx.Field
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is set",
			expected: newField("key", "val", t),
			key:      "key",
			value:    "val",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := pectx.Field{}
			f.Set(testCase.key, testCase.value)

			// Check if the value is set
			key, value := testCase.expected.Get()
			if key != testCase.key {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.key, key)
			}
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
		expected *pectx.Field
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is overwritten",
			expected: newField("key", "val2", t),
			key:      "key",
			value:    "val2",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := pectx.Field{}
			f.Set(testCase.key, "val")
			f.Set(testCase.key, testCase.value)

			// Check if the value is set
			key, value := testCase.expected.Get()
			if key != testCase.key {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.key, key)
			}
			if value != testCase.value {
				t.Errorf("%s: expected %s, got %s", testCase.desc, testCase.value, value)
			}

		})
	}
}
