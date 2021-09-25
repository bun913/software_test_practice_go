package pettemp

import (
	"testing"
)

func TestPetTemp_GetTempCondtion(t *testing.T) {
	tests := []struct {
		name string
		temp float64
		want string
	}{
		{"寒いの代表値:20.0", 20.0, ColdMessage},
		{"寒いの境界値:23.9", 23.9, ColdMessage},
		{"快適の境界値:24.0", 24.0, SuitableMessage},
		{"快適の代表値:25.0", 25.0, SuitableMessage},
		{"快適の境界値:25.9", 25.9, SuitableMessage},
		{"暑いの境界値:26.0", 26.0, HotMessage},
		{"暑いの代表値:30.0", 30.0, HotMessage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			petTemp := &PetTemp{
				Temperture: tt.temp,
			}
			if got := petTemp.GetTempCondtion(); got != tt.want {
				t.Errorf("PetTemp.GetTempCondtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
