package equip

/*
	Weapon
	武器・盾を表す構造体
	EquipmentType 1:片手持ち、2:両手持ち
*/
const TypeOneHand int = 1
const TypeTwoHand int = 2

func NewWwapon(name string, attackpoint int, typeNum int, reinforceable bool) *Weapon {
	return &Weapon{
		Name:          name,
		AttackPoint:   attackpoint,
		EquipmentType: typeNum,
		Reinfoceable:  reinforceable,
	}
}

type Weapon struct {
	Name          string
	AttackPoint   int
	EquipmentType int
	Reinfoceable  bool
}

type Sheild struct {
	Name string
}

func NewSheild(name string) *Sheild {
	return &Sheild{
		Name: name,
	}
}
