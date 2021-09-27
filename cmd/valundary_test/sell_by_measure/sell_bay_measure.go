package sellbymeasure

import "errors"

type Material struct {
	Meater float64
}

const normalUnitPrice float64 = 400.0
const discountUnitPrice float64 = 350.0

// NewMaterial M数を与えて新規Materialを返却
func NewMaterial(meater float64) *Material {
	return &Material{
		Meater: meater,
	}
}

// CalcResult メーター数から単価を計算する
func (material *Material) CalcResult() (int, error) {
	var e error
	var result int
	if material.isIvalid() {
		e = errors.New("メートル数が不正な値です")
		return result, e
	}
	result = material.calcPrice()
	return result, e
}

func (material *Material) isIvalid() bool {
	result := false
	if material.isTooShort() || material.isTooLong() {
		result = true
	}
	return result
}

func (material *Material) calcPrice() int {
	meater := material.Meater
	var unit_price float64
	if meater <= 3.0 {
		unit_price = normalUnitPrice
	} else {
		unit_price = discountUnitPrice
	}
	result := meater * unit_price
	return int(result)
}

func (material *Material) isTooShort() bool {
	result := false
	if material.Meater <= 0.0 {
		result = true
	}
	return result
}

func (material *Material) isTooLong() bool {
	result := false
	if material.Meater > 100.0 {
		result = true
	}
	return result
}
