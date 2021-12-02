package exercisemodel

type Filter struct {
	BodyParts string `json:"body_parts" gorm:"column:body_parts;"`
	Mechanics string `json:"mechanics" gorm:"column:mechanics;"`
	Force     string `json:"force" gorm:"column:force;"`
}
