package tax

import "testing"

// func TestCalculateTax(t *testing.T) {
// 	amount := 500.0
// 	expected := 5.0

// 	result := CalculateTax(amount)

// 	if result != expected {
// 		t.Errorf("Expected %f but got %f", expected, result)
// 	}
// }

// func TestCalculateTaxBatch(t *testing.T) {
// 	type calcTax struct {
// 		amount, expected float64
// 	}

// 	tests := []calcTax{
// 		{500.0, 5.0},
// 		{1000.0, 10.0},
// 		{1500.0, 10.0},
// 		{0.0, 0.0},
// 	}

// 	for i, test := range tests {
// 		result := CalculateTax(test.amount)
// 		if result != test.expected {
// 			t.Errorf("Expected %f but got %f on index %d", test.expected, result, i+1)
// 		}
// 	}
// }

// func BenchmarkCalculateTax(b *testing.B) {
// 	for b.Loop() {
// 		CalculateTax(500.0)
// 	}
// }

// func BenchmarkCalculateTax2(b *testing.B) {
// 	for b.Loop() {
// 		CalculateTax2(500.0)
// 	}
// }

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500, 1000, 1500}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expect 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expect 20", result)
		}
	})
}
