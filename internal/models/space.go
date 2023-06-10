package models

type Space struct {
	ID                int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string     `json:"name"`
	Color1            string     `json:"color1"`
	Color2            string     `json:"color2"`
	Color3            string     `json:"color3"`
	Color4            string     `json:"color4"`
	ImageUrl          string     `json:"imageUrl"`
	PrimaryLanguageID int        `json:"primaryLanguageId"`
	PrimaryLanguage   Language   `json:"primaryLanguage" gorm:"foreignkey:PrimaryLanguageID"`
	Languages         []Language `json:"languages" gorm:"many2many:space_languages;"`
}
