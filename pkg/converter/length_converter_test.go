package converter

import (
	"testing"
)

func TestLengthConverter(t *testing.T) {
	testCases := []struct {
		length      Length
		expected    float64
		expectError bool
	}{
		{Length{10.0, METRIC}, 25.4, false},
		{Length{-10.0, METRIC}, -25.4, false},
		{Length{10.0, IMPERIAL}, 3.94, false},
		{Length{-10.0, IMPERIAL}, -3.94, false},
		{Length{0, METRIC}, 0, true},
	}

	for _, testCase := range testCases {
		convertedLength, err := testCase.length.Convert()

		if testCase.expectError && err == nil {
			t.Errorf("expected error but got none")
		}

		if !testCase.expectError && err != nil {
			t.Errorf("expected no error but got %s", err.Error())
		}

		if convertedLength != testCase.expected {
			t.Errorf("expected %f but got %f", testCase.expected, convertedLength)
		}
	}
}

func TestNewLength(t *testing.T) {
	testCases := []struct {
		value          float64
		system         LengthType
		expectedLength float64
	}{
		{2.456, METRIC, 2.46},
		{-2.456, IMPERIAL, -2.46},
		{0, METRIC, 0},
	}

	for _, testCase := range testCases {
		length := NewLength(testCase.value, testCase.system)

		if length.value != testCase.expectedLength {
			t.Errorf("expected %f but got %f", testCase.expectedLength, length.value)
		}
	}
}
