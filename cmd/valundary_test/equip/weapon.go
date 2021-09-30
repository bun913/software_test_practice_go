package equip

import (
	"strings"
)

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

func (w *Weapon) MakeStrong() (*Weapon, string) {
	var m string
	// 強化不可な武器種
	if !w.Reinfoceable {
		m = "武器の強化に失敗しました"
		return w, m
	}
	// すでに強化済み
	isAlreday := strings.HasSuffix(w.Name, "+")
	if isAlreday {
		m = "これ以上この武器の強化はできません"
		return w, m
	}
	w.AttackPoint += 10
	w.Name = w.Name + "+"
	m = "強化に成功しました"
	return w, m
}
