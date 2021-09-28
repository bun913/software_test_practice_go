package equip

import "testing"

func TestCharcter_EquipLeft(t *testing.T) {
	type fields struct {
		RightHand *Weapon
		LeftHand  *Sheild
	}
	type args struct {
		weapon *Weapon
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "同値分割1: 片手武器装備可能",
			fields: fields{
				RightHand: NewWwapon("ニンジンの短剣", 20, TypeOneHand, false),
				LeftHand:  NewSheild("盾"),
			},
			wantErr: false,
		},
		{
			name: "同値分割2: 両手武器装備可能",
			fields: fields{
				RightHand: NewWwapon("竜の槍", 50, TypeTwoHand, true),
				LeftHand:  NewSheild("盾"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char := &Charcter{
				RightHand: tt.fields.RightHand,
				LeftHand:  tt.fields.LeftHand,
			}
			if err := char.EquipLeft(tt.fields.RightHand); (err != nil) != tt.wantErr {
				t.Errorf("Charcter.EquipLeft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
