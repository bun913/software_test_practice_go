package healthcheck

import "errors"

type Target struct {
	BMI float64
}

func (t *Target) GetCondition() (string, error) {
	var e error
	var m string
	switch {
	case t.isError():
		e = errors.New("エラー。BMIは0より大きく、99.9までの値で入力してください")
	case t.isSkinny():
		m = "痩せ"
	case t.isMidium():
		m = "普通"
	case t.isPreObesity():
		m = "前肥満"
	case t.isObesity1():
		m = "肥満1"
	case t.isObesity2():
		m = "肥満2"
	case t.isObesity3():
		m = "肥満3"
	}
	return m, e
}

func (t *Target) isError() bool {
	bmi := t.BMI
	if bmi <= 0.0 || bmi >= 100.0 {
		return true
	}
	return false
}

func (t *Target) isSkinny() bool {
	bmi := t.BMI
	if bmi >= 0.1 && bmi <= 18.4 {
		return true
	}
	return false
}

func (t *Target) isMidium() bool {
	bmi := t.BMI
	if bmi >= 18.5 && bmi <= 24.9 {
		return true
	}
	return false
}

func (t *Target) isPreObesity() bool {
	bmi := t.BMI
	if bmi >= 25.0 && bmi <= 29.9 {
		return true
	}
	return false
}

func (t *Target) isObesity1() bool {
	bmi := t.BMI
	if bmi >= 30.0 && bmi <= 34.9 {
		return true
	}
	return false
}

func (t *Target) isObesity2() bool {
	bmi := t.BMI
	if bmi >= 35.0 && bmi <= 39.9 {
		return true
	}
	return false
}

func (t *Target) isObesity3() bool {
	bmi := t.BMI
	if bmi >= 40.0 && bmi <= 99.9 {
		return true
	}
	return false
}
