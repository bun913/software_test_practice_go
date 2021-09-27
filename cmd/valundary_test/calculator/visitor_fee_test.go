package calculator

import "testing"

func TestVisitor_CalcCharge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		want    int
		wantErr bool
	}{
		{"マイナス側無効クラスの代表値", -10, 0, true},
		{"マイナス側無効クラスの上限境界値", -1, 0, true},
		{"無料クラスの下限境界値", 0, 0, false},
		{"無料クラスの代表値", 3, 0, false},
		{"無料クラスの上限境界値", 5, 0, false},
		{"500円クラスの下限境界値", 6, 500, false},
		{"500円クラスの代表値", 10, 500, false},
		{"500円クラスの上限境界値", 12, 500, false},
		{"1000円クラスの下限境界値", 13, 1000, false},
		{"1000円クラスの代表値", 15, 1000, false},
		{"1000円クラスの上限境界値", 17, 1000, false},
		{"1500円クラスの下限境界値", 18, 1500, false},
		{"1500円クラスの代表値", 50, 1500, false},
		{"1500円クラスの上限境界値", 120, 1500, false},
		{"プラス側無効クラスの下限境界値", 121, 0, true},
		{"プラス側無効クラスの代表値", 150, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visitor := &Visitor{
				Age: tt.age,
			}
			got, err := visitor.CalcCharge()
			if (err != nil) != tt.wantErr {
				t.Errorf("Visitor.CalcCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Visitor.CalcCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
