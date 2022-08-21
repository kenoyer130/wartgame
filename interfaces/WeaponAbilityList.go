package interfaces

type WeaponAbilityList interface {
	Init()
	ApplyWeaponAbilities(phase WeaponAbilityPhase, die int, weapon IWeaponAbility) IWeaponAbility
}

type IWeaponAbility interface {
	SetArmorPiercing(value int)
}

type WeaponAbilityPhase string

const (
	WeaponAbilityPhaseWounds WeaponAbilityPhase = "Wounds"
)