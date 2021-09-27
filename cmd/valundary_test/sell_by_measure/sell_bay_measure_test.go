package sellbymeasure

import "testing"

func TestMaterial_CalcResult(t *testing.T) {
	tests := []struct {
		name    string
		meater  float64
		want    int
		wantErr bool
	}{
		{"無効値エラーの上限境界値（代表値）", 0.0, 0, true},
		{"有効(400円単価)の下限境界値", 0.1, 40, false},
		{"有効(400円単価)の代表値", 1.0, 400, false},
		{"有効(400円単価)の上限境界値", 3.0, 1200, false},
		{"有効(350円単価)の下限境界値", 3.1, 1085, false},
		{"有効(350円単価)の代表値", 4.0, 1400, false},
		{"有効(350円単価)の上限境界値", 100.0, 35000, false},
		{"無効値エラーの下限境界値", 100.1, 0.0, true},
		{"無効値エラーの代表値", 200.0, 0.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			material := NewMaterial(tt.meater)
			got, err := material.CalcResult()
			if (err != nil) != tt.wantErr {
				t.Errorf("Material.CalcResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Material.CalcResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
