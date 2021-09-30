package equip

import (
	"testing"
)

func TestWeapon_MakeStrong(t *testing.T) {
	type fields struct {
		Name          string
		AttackPoint   int
		EquipmentType int
		Reinfoceable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "強化不可クラスの代表値",
			fields: fields{
				Name:          "人参の短剣",
				AttackPoint:   20,
				EquipmentType: 1,
				Reinfoceable:  false,
			},
			want: "武器の強化に失敗しました",
		},
		{
			name: "強化可能クラスの代表値",
			fields: fields{
				Name:          "草薙の剣",
				AttackPoint:   90,
				EquipmentType: 2,
				Reinfoceable:  true,
			},
			want: "強化に成功しました",
		},
		{
			name: "これ以上強化できないクラスの代表値",
			fields: fields{
				Name:          "草薙の剣+",
				AttackPoint:   100,
				EquipmentType: 2,
				Reinfoceable:  true,
			},
			want: "これ以上この武器の強化はできません",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weapon{
				Name:          tt.fields.Name,
				AttackPoint:   tt.fields.AttackPoint,
				EquipmentType: tt.fields.EquipmentType,
				Reinfoceable:  tt.fields.Reinfoceable,
			}
			_, got1 := w.MakeStrong()
			if got1 != tt.want {
				t.Errorf("Weapon.MakeStrong() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
