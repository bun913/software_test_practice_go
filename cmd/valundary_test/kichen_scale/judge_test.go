package kichenscale

import "testing"

func TestScale_GetResult(t *testing.T) {
	type fields struct {
		Weight int
	}
	tests := []struct {
		name   string
		weight int
		want   string
	}{
		{"マイナスエラーの代表値", -100, "EEEE"},
		{"マイナスエラーの上限境界値", -1, "EEEE"},
		{"正常クラスの下限境界値", 0, "0g"},
		{"正常クラスの代表値", 1000, "1000g"},
		{"正常クラスの上限境界値", 2000, "2000g"},
		{"オーバークラスの下限境界値", 2001, "EEEE"},
		{"オーバークラスの代表値", 2500, "EEEE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scale := NewScale(tt.weight)
			if got := scale.GetResult(); got != tt.want {
				t.Errorf("Scale.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
