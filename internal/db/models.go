package db

type Tutor struct {
	ID           int    `gorm:"type:int(10) NOT NULL auto_increment;primary_key;" json:"id"`
	Slug         string `gorm:"type:varchar(50) DEFAULT NULL" json:"slug"`
	Name         string `gorm:"type:varchar(50) DEFAULT NULL" json:"name"`
	Headline     string `gorm:"type:varchar(100) DEFAULT NULL" json:"headline"`
	Introduction string `gorm:"type:text;" json:"introduction"`
}

type TutorLessonPrice struct {
	ID          int     `gorm:"type:int(10) NOT NULL auto_increment;primary_key;" json:"id"`
	TutorId     int     `gorm:"type:int(11) DEFAULT NULL" json:"tutor_id"`
	TrialPrice  float32 `gorm:"type:float(11, 0) DEFAULT NULL" json:"trial_price"`
	NormalPrice float32 `gorm:"type:float(11, 0) DEFAULT NULL" json:"normal_price"`
}

type TutorLanguage struct {
	ID         int `gorm:"type:int(10) NOT NULL auto_increment;primary_key;" json:"id"`
	TutorId    int `gorm:"type:int(11) DEFAULT NULL" json:"tutor_id"`
	LanguageId int `gorm:"type:int(11) DEFAULT NULL" json:"language_id"`
}

type Language struct {
	ID   int    `gorm:"type:int(10) NOT NULL auto_increment;primary_key;" json:"id"`
	Slug string `gorm:"type:varchar(30) DEFAULT NULL" json:"slug"`
}
