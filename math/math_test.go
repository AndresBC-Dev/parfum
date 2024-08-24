package math

import "testing"

func TestSum(t *testing.T) {
	result := Sum(8, 7)
	expected := 15
	if result != expected {
		t.Errorf("sum was incorrect, got %v, and expeted %v", result, expected)
	}

}

func TestSumString(t *testing.T) {
	testCases := []struct {
		Name     string
		A        string
		B        string
		Expected int
	}{{

		Name:     "Sum two numbers",
		A:        "1",
		B:        "2",
		Expected: 3,
	},
		{
			Name:     "Incorrect A parameter",
			A:        "a",
			B:        "2",
			Expected: 0,
		},
		{
			Name:     "Incorrect B parameter",
			A:        "1",
			B:        "b",
			Expected: 0,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			result := sumString(tc.A, tc.B)

			t.Log("Result: ", result)

			if result != tc.Expected {
				t.Errorf("sum was incorrect, got: %d, want: %d.", result, tc.Expected)
			}

		})

	}
}
