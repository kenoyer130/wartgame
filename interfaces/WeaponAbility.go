package interfaces

type WeaponAbility interface {
	GetType() string
	GetPhase() WeaponAbilityPhase
	Apply(die int, weapon IWeaponAbility)
}