package calculator

import "errors"

const unit_of_meter float64 = 1.65

// Tatami 畳の乗数等の情報を保持する
type Tatami struct {
	SheetsNum int8
}

// NewTatami 新規Tatami型の返却
func NewTatami(sheetnum int8) *Tatami {
	return &Tatami{
		SheetsNum: sheetnum,
	}
}

// CalcSquareMeter 畳数から平方メートルを計算する
func (tatami *Tatami) CalcSquareMeter() (float64, error) {
	var e error
	var result float64
	// エラーチェック
	if tatami.isInvalid() {
		e = errors.New("無効な値です。１以上127以下の整数を入力してください")
		return result, e
	}
	// 誤差防止のため一旦整数値に戻す
	result = float64(tatami.SheetsNum) * 100 * unit_of_meter / 100
	return result, e
}

// isInvalid SheetNumのバリデーションチェック
// 計算できない無効な値ならtrueを返す
func (tatami *Tatami) isInvalid() bool {
	sheetnum := tatami.SheetsNum
	result := false
	if sheetnum <= 0 {
		result = true
	}
	return result
}
