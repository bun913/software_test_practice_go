package calculator

type Discount struct {
	Day int
}

func (d *Discount) CalcDiscountRate() float64 {
	result := 0.0
	if d.isDisCount() {
		result = 0.2
	}
	return result
}

func (d *Discount) isDisCount() bool {
	day := d.Day
	result := false
	if (day >= 1 && day <= 5) || (day >= 28 && day <= 30) {
		result = true
	}
	return result
}
