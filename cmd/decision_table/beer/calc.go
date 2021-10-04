package beer

type PriceCondition struct {
	Coupon    bool
	HappyHour bool
}

func (cp *PriceCondition) Calc() int {
	result := 490
	if cp.isHappyHour() {
		result = 290
	}
	if cp.isCoupon() {
		result = 100
	}
	return result
}

func (cp *PriceCondition) isCoupon() bool {
	result := false
	if cp.Coupon {
		result = true
	}
	return result
}

func (cp *PriceCondition) isHappyHour() bool {
	result := false
	if cp.HappyHour {
		result = true
	}
	return result
}
