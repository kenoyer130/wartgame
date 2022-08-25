package interfaces

type WeaponAbilityList interface {
	Init()
	ApplyWeaponAbilities(phase WeaponAbilityPhase, die int, weapon IWeaponAbility)
}

type IWeaponAbility interface {
	SetArmorPiercing(value int)
	GetWeaponAbilities() []string
}

type WeaponAbilityPhase string

const (
	WeaponAbilityPhaseWounds WeaponAbilityPhase = "Wounds"
)