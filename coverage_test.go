package coverage

import "testing"

func TestConvertToPercentage(t *testing.T) {
	expected := "5.83"
	input := 0.05829249107521026
	tested := ConvertToPercentage(input)

	if tested != expected {
		t.Errorf("Percentage incorrect, got: %s, expected: %s", tested, expected)
	}
}
