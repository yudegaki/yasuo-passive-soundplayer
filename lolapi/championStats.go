package lolapi

type ChampionStats struct {
	AbilityHaste float32 `json:"abilityHaste"`
	AbilityPower float32 `json:"abilityPower"`
	Armor float32 `json:"armor"`
	ArmorPenetrationFlat float32 `json:"armorPenetrationFlat"`
	ArmorPenetrationPercent float32 `json:"armorPenetrationPercent"`
	AttackDamage float32 `json:"attackDamage"`
	AttackRange float32 `json:"attackRange"`
	AttackSpeed float32 `json:"attackSpeed"`
	BonusArmorPenetrationPercent float32 `json:"bonusArmorPenetrationPercent"`
	BonusMagicPenetrationPercent float32 `json:"bonusMagicPenetrationPercent"`
	CritChance float32 `json:"critChance"`
	CritDamage float32 `json:"critDamage"`
	CurrentHealth float32 `json:"currentHealth"`
	HealthRegenRate float32 `json:"healthRegenRate"`
	LifeSteal float32 `json:"lifeSteal"`
	MagicLethality float32 `json:"magicLethality"`
	MagicPenetrationFlat float32 `json:"magicPenetrationFlat"`
	MagicPenetrationPercent float32 `json:"magicPenetrationPercent"`
	MagicResist float32 `json:"magicResist"`
	MaxHealth float32 `json:"maxHealth"`
	MoveSpeed float32 `json:"moveSpeed"`
	PhysicalLethality float32 `json:"physicalLethality"`
	ResourceMax float32 `json:"resourceMax"`
	ResourceRegenRate float32 `json:"resourceRegenRate"`
	ResourceType string `json:"resourceType"`
	ResourceValue float32 `json:"resourceValue"`
	SpellVamp float32 `json:"spellVamp"`
	Tenacity float32 `json:"tenacity"`
}
