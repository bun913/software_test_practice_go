package calculator

import "errors"

const freeCharge int = 0
const level1Charge int = 500
const level2Charge int = 1000
const level3Charge int = 1500

type Visitor struct {
	Age int
}

func (visitor *Visitor) CalcCharge() (int, error) {
	var e error
	var charge int
	if visitor.isInValid() {
		e = errors.New("年齢が不正です。0から120までの整数で入力してください")
	}
	switch {
	case visitor.isLevel1Charge():
		charge = level1Charge
	case visitor.isLevel2Charge():
		charge = level2Charge
	case visitor.isLevel3Charge():
		charge = level3Charge
	}
	return charge, e
}

func (visitor *Visitor) isInValid() bool {
	result := false
	age := visitor.Age
	if age < 0 || age > 120 {
		result = true
	}
	return result
}

func (visitor *Visitor) isFree() bool {
	result := false
	age := visitor.Age
	if age < 6 {
		result = false
	}
	return result
}

func (visitor *Visitor) isLevel1Charge() bool {
	result := false
	age := visitor.Age
	if age >= 6 && age < 13 {
		result = true
	}
	return result
}

func (visitor *Visitor) isLevel2Charge() bool {
	result := false
	age := visitor.Age
	if age >= 13 && age < 18 {
		result = true
	}
	return result
}

func (visitor *Visitor) isLevel3Charge() bool {
	result := false
	age := visitor.Age
	if age >= 18 && age <= 120 {
		result = true
	}
	return result
}
