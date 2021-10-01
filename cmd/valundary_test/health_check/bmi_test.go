package healthcheck

import "testing"

func TestTarget_GetCondition(t *testing.T) {
	type fields struct {
		BMI float64
	}
	tests := []struct {
		name    string
		bmi     float64
		want    string
		wantErr bool
	}{
		{"無効値エラーの上限境界値（代表値）", 0.0, "", true},
		{"痩せの下限境界値", 0.1, "痩せ", false},
		{"痩せの代表値", 10.0, "痩せ", false},
		{"痩せの上限境界値", 18.4, "痩せ", false},
		{"普通の下限境界値", 18.5, "普通", false},
		{"普通の代表値", 21.0, "普通", false},
		{"普通の上限境界値", 24.9, "普通", false},
		{"前肥満の下限境界値", 25.0, "前肥満", false},
		{"前肥満の代表値", 27.5, "前肥満", false},
		{"前肥満の上限境界値", 29.9, "前肥満", false},
		{"肥満1の下限境界値", 30.0, "肥満1", false},
		{"肥満1の代表値", 32.5, "肥満1", false},
		{"肥満1の上限境界値", 34.9, "肥満1", false},
		{"肥満2の下限境界値", 35.0, "肥満2", false},
		{"肥満2の代表値", 37.5, "肥満2", false},
		{"肥満2の上限境界値", 39.9, "肥満2", false},
		{"肥満3の下限境界値", 40.0, "肥満3", false},
		{"肥満3の代表値", 60.0, "肥満3", false},
		{"肥満3の上限境界値", 99.9, "肥満3", false},
		{"エラーの下限境界値", 100.0, "", true},
		{"エラーの代表値", 120.0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Target{
				BMI: tt.bmi,
			}
			got, err := tr.GetCondition()
			if (err != nil) != tt.wantErr {
				t.Errorf("Target.GetCondition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Target.GetCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
