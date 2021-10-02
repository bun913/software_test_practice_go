package calculator

import "testing"

func TestDiscount_CalcDiscountRate(t *testing.T) {
	type fields struct {
		Day int
	}
	tests := []struct {
		name string
		day  int
		want float64
	}{
		{"早い日の割引下限境界値", 1, 0.2},
		{"早い日の割引代表値", 3, 0.2},
		{"早い日の割引上限境界値", 5, 0.2},
		{"割引なしの下限境界値", 6, 0.0},
		{"割引なしの代表値", 16, 0.0},
		{"割引なしの代表値", 27, 0.0},
		{"遅い日の割引下限境界値", 28, 0.2},
		{"遅い日の割引代表値", 29, 0.2},
		{"遅い日の割引上限境界値", 30, 0.2},
		{"割引なしの下限（上限・代表)値", 31, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discount{
				Day: tt.day,
			}
			if got := d.CalcDiscountRate(); got != tt.want {
				t.Errorf("Discount.CalcDiscountRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
