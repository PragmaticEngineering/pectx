package mapStore_test

import (
	"reflect"
	"testing"

	ms "github.com/pragmaticengineering/pectx/stores/map"
)

func TestNewMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *ms.Map
		key      string
		val      string
	}{
		{
			desc:     "validate that the value is set",
			expected: ms.NewMap(map[string]string{"key": "val"}),
			key:      "key",
			val:      "val",
		},
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

func TestGet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected *ms.Map
		key      string
		exists   bool
		val      string
	}{
		{
			desc:     "validate Store values",
			expected: ms.NewMap(map[string]string{"key": "val"}),
			key:      "key",
			exists:   true,
			val:      "val",
		},
		{
			desc:     "validate that the value doesn't exist",
			expected: ms.NewMap(map[string]string{"key": ""}),
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
		expected *ms.Map
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is set",
			expected: ms.NewMap(map[string]string{"key": "val"}),
			key:      "key",
			value:    "val",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := ms.Map{}
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
		expected *ms.Map
		key      string
		value    string
	}{
		{
			desc:     "validate that the value is overwritten",
			expected: ms.NewMap(map[string]string{"key": "val2"}),
			key:      "key",
			value:    "val2",
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := ms.Map{}
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

func TestListKeys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected []string
		input    map[string]string
	}{
		{
			desc:     "validate that the keys are listed",
			expected: []string{"key", "key2"},
			input:    map[string]string{"key": "val", "key2": "val2"},
		},
		{
			desc:     "validate that no keys are listed",
			expected: []string{},
			input:    map[string]string{},
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			f := ms.NewMap(testCase.input)
			keys := f.ListKeys()
			if !reflect.DeepEqual(keys, testCase.expected) {
				t.Errorf("%s: expected %v, got %v", testCase.desc, testCase.expected, keys)
			}

		})
	}
}
