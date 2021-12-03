package db

import (
	"fmt"

	"github.com/CharlesChou03/_git/amt.git/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLTutorDB struct {
	DB *gorm.DB
}

var (
	MySQLDB                *MySQLTutorDB
	TableTutors            string = "tutors"
	TableTutorLessonPrices string = "tutor_lesson_prices"
	TableTutorLanguages    string = "tutor_languages"
	TableLanguages         string = "languages"
)

func SetupMySQLDB() *MySQLTutorDB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("connection to mysql failed:", err)
	}

	return &MySQLTutorDB{DB: db}
}

func (db *MySQLTutorDB) CreateTutorTable() {
	db.DB.Set("gorm:table_options", "COLLATE=utf8mb4_general_ci").AutoMigrate(&Tutor{})
	db.DB.Set("gorm:table_options", "COLLATE=utf8mb4_general_ci").AutoMigrate(&TutorLessonPrice{})
	db.DB.Set("gorm:table_options", "COLLATE=utf8mb4_general_ci").AutoMigrate(&TutorLanguage{})
	db.DB.Set("gorm:table_options", "COLLATE=utf8mb4_general_ci").AutoMigrate(&Language{})
}

func (db *MySQLTutorDB) Close() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func (db *MySQLTutorDB) SearchTutors(tutor *Tutor) bool {
	if result := db.DB.Where("slug = ?", &tutor.Slug).Take(&tutor); result.Error != nil {
		fmt.Print("+v%", result.Error)
		return false
	}
	return true
}

func (db *MySQLTutorDB) SearchTutorLessonPrices(tutorId int, tutorLessonPrices *TutorLessonPrice) bool {
	if result := db.DB.Where("tutor_id = ?", tutorId).Take(&tutorLessonPrices); result.Error != nil {
		fmt.Print("+v%", result.Error)
		return false
	}
	return true
}

func (db *MySQLTutorDB) SearchTutorLanguage(tutorId int, tutorLanguages *[]TutorLanguage) bool {
	if result := db.DB.Where("tutor_id = ?", tutorId).Find(&tutorLanguages); result.Error != nil {
		fmt.Print("+v%", result.Error)
		return false
	}
	return true
}

// func (db *MySQLTutorDB) SearchLanguages(languageIds []string, languages *[]Language) bool {
// 	if result := db.DB.Where("id IN ?", languageIds).Find(&languages); result.Error != nil {
// 		fmt.Print("+v%", result.Error)
// 		return false
// 	}
// 	return true
// }
