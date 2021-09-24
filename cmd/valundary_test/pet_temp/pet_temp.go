package pettemp

const ColdMessage string = "寒い"
const SuitableMessage string = "適温"
const HotMessage string = "暑い"

type PetTemp struct {
	Temperture float64
}

func newPetTemp() *PetTemp {
	return &PetTemp{}
}

// GetTempCondition
// ペットにとって現在の温度がどのような状況か文字列で返却する
func (petTemp *PetTemp) GetTempCondtion() string {
	message := ""
	switch {
	case petTemp.isCold():
		message = ColdMessage
	case petTemp.isSuitable():
		message = SuitableMessage
	case petTemp.isHot():
		message = HotMessage
	}
	return message
}

func (petTemp *PetTemp) isCold() bool {
	temp := petTemp.Temperture
	result := false
	if temp < 24.0 {
		result = true
	}
	return result
}

func (petTemp *PetTemp) isHot() bool {
	temp := petTemp.Temperture
	result := false
	if temp >= 26.0 {
		result = true
	}
	return result
}

func (petTemp *PetTemp) isSuitable() bool {
	temp := petTemp.Temperture
	result := false
	if temp >= 24.0 && temp < 26.0 {
		result = true
	}
	return result
}
