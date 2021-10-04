package beer

import "testing"

func TestPriceCondition_Calc(t *testing.T) {
	type fields struct {
		Coupon    bool
		HappyHour bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"デシジョンテーブル1: クーポン利用　ハッピーアワー利用",
			fields{Coupon: true, HappyHour: true},
			100,
		},
		{
			"デシジョンテーブル1: クーポン利用　ハッピーアワー利用しない",
			fields{Coupon: true, HappyHour: false},
			100,
		},
		{
			"デシジョンテーブル2: クーポン利用しない　ハッピーアワー利用",
			fields{Coupon: false, HappyHour: true},
			290,
		},
		{
			"デシジョンテーブル3: クーポン利用しない　ハッピーアワー利用しない",
			fields{Coupon: false, HappyHour: false},
			490,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &PriceCondition{
				Coupon:    tt.fields.Coupon,
				HappyHour: tt.fields.HappyHour,
			}
			if got := cp.Calc(); got != tt.want {
				t.Errorf("PriceCondition.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
