package equip

import "errors"

type Charcter struct {
	RightHand *Weapon
	LeftHand  *Sheild
}

func (char *Charcter) EquipLeft(weapon *Weapon) error {
	var e error
	right := char.RightHand
	if right == nil {
		// 不便なシステム・・・
		e = errors.New("右手から先に装備してください")
		return e
	}
	if !char.IsAbleEquip() {
		e = errors.New("装備できません")
		return e
	}
	// 装備
	char.AttachWeapon(weapon)
	return e
}

func (char *Charcter) AttachWeapon(weapon *Weapon) {
}

// IsAbleEquip 右手が片手武器の場合だけTrueを返す
// 先に左に装備を持たせることはないように別途制御されているものとする
func (char *Charcter) IsAbleEquip() bool {
	right := char.RightHand
	category := right.EquipmentType
	result := false
	if category == TypeOneHand {
		result = true
	}
	return result
}
