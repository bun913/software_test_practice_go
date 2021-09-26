package calculator

import (
	"testing"
)

func TestTatamiCalcSquareMeter(t *testing.T) {
	tests := []struct {
		name     string
		sheetnum int8
		want     float64
		wantErr  bool
	}{
		{"無効値エラーの下限境界値", int8(-128), 0.0, true},
		{"無効値エラーの代表値", int8(-64), 0.0, true},
		{"無効値エラーの上限境界値", int8(0), 0.0, true},
		{"正常値の下限境界値", int8(1), 1.65, false},
		{"正常値の代表値", int8(64), 105.60, false},
		{"正常値の上限代表値", int8(127), 209.55, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tatami := NewTatami(tt.sheetnum)
			got, err := tatami.CalcSquareMeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tatami.CalcSquareMeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Tatami.CalcSquareMeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
