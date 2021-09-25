package kichenscale

import "strconv"

type Scale struct {
	Weight        int
	minceDeadLine int
	overDeadLine  int
	errorString   string
	unit          string
}

func NewScale(weight int) *Scale {
	return &Scale{
		Weight:        weight,
		minceDeadLine: 0,
		overDeadLine:  2000,
		errorString:   "EEEE",
		unit:          "g",
	}
}

// GetResult 計測の結果を返却する
func (scale *Scale) GetResult() string {
	switch {
	case scale.isError():
		return scale.errorString
	default:
		return scale.getDispString()
	}
}

// エラー表示とするべきか判断
func (scale *Scale) isError() bool {
	result := false
	if scale.isMinceError() || scale.isOverError() {
		result = true
	}
	return result
}

// isMinceError 0gより小さい値の場合エラーと判断する
func (scale *Scale) isMinceError() bool {
	result := false
	weight := scale.Weight
	if weight < scale.minceDeadLine {
		result = true
	}
	return result
}

// isOverError 上限(2000)より大きい値の場合エラーと判断する
func (scale *Scale) isOverError() bool {
	result := false
	weight := scale.Weight
	if weight > scale.overDeadLine {
		result = true
	}
	return result
}

// getDispString 正常表示のグラム数を返却する
func (scale *Scale) getDispString() string {
	i := scale.Weight
	return strconv.Itoa(i) + scale.unit
}
