package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/CharlesChou03/_git/amt.git/internal/db"
	"github.com/CharlesChou03/_git/amt.git/models"
	modelsReq "github.com/CharlesChou03/_git/amt.git/models/request"
	modelsRes "github.com/CharlesChou03/_git/amt.git/models/response"
)

func GetTutor(req *modelsReq.GetTutorReq, res *modelsRes.GetTutorRes) (int, *modelsRes.GetTutorRes, models.TutorInfoError) {
	if !req.Validate() {
		return 400, res, models.BadRequestError
	}

	cacheStatus, cacheResult := GetTutorInfoFromCache(req.Tutor)
	if cacheStatus == false {
		return 204, res, models.NotFoundError
	}

	return 200, &cacheResult, models.NoError
}

func GetTutorInfoFromCache(slug string) (bool, modelsRes.GetTutorRes) {
	tutorInfoFromCache := db.RedisDB.GetValueFromCache(slug, 30)
	tutorInfo := modelsRes.GetTutorRes{}
	if tutorInfoFromCache == "" {
		generateResult, newTutorInfo := GenerateTutorCache(slug)
		return generateResult, newTutorInfo
	}
	// [TODO] handle expire situation
	expired := false
	if expired == true {
		go GenerateTutorCache(slug)
		return true, tutorInfo
	}
	if err := json.Unmarshal([]byte(tutorInfoFromCache), &tutorInfo); err != nil {
		fmt.Print("decode aggregate result error")
		return false, tutorInfo
	}
	return true, tutorInfo
}

func GenerateTutorCache(slug string) (bool, modelsRes.GetTutorRes) {
	db.RedisDB.DeleteCacheByKey(slug)
	tutorInfo := modelsRes.GetTutorRes{}
	searchResult, tutor, tutorLessonPrices, tutorLanguages := GetTutorInfoFromDB(slug)
	if searchResult == true {
		aggregateResult := AggregateTutorInfo(tutor, tutorLessonPrices, tutorLanguages)
		aggregateResultEncoding, err := json.Marshal(aggregateResult)
		if err != nil {
			fmt.Print("encode aggregate result error")
			return false, tutorInfo
		}
		db.RedisDB.SetValueToCache(slug, aggregateResultEncoding)
		return true, aggregateResult
	} else {
		return false, tutorInfo
	}
}

func GetTutorInfoFromDB(slug string) (bool, db.Tutor, db.TutorLessonPrice, []db.TutorLanguage) {
	tutor := db.Tutor{}
	tutorLessonPrice := db.TutorLessonPrice{}
	tutorLanguages := []db.TutorLanguage{}

	tutor.Slug = slug
	searchTutorResult := db.MySQLDB.SearchTutors(&tutor)
	if searchTutorResult == false {
		return false, tutor, tutorLessonPrice, tutorLanguages
	}
	tutorId := tutor.ID
	searchTutorResult = db.MySQLDB.SearchTutorLessonPrices(tutorId, &tutorLessonPrice)
	if searchTutorResult == false {
		return false, tutor, tutorLessonPrice, tutorLanguages
	}

	searchTutorResult = db.MySQLDB.SearchTutorLanguage(tutorId, &tutorLanguages)
	if searchTutorResult == false {
		return false, tutor, tutorLessonPrice, tutorLanguages
	}

	return true, tutor, tutorLessonPrice, tutorLanguages
}

func AggregateTutorInfo(tutor db.Tutor,
	tutorLessonPrice db.TutorLessonPrice,
	tutorLanguages []db.TutorLanguage) modelsRes.GetTutorRes {

	tutorInfo := modelsRes.GetTutorRes{}
	tutorInfo.ID = strconv.Itoa(tutor.ID)
	tutorInfo.Slug = tutor.Slug
	tutorInfo.Name = tutor.Name
	tutorInfo.Headline = tutor.Headline
	tutorInfo.Introduction = tutor.Introduction
	priceInfo := modelsRes.PriceInfo{}
	priceInfo.Trial = tutorLessonPrice.TrialPrice
	priceInfo.Normal = tutorLessonPrice.NormalPrice
	tutorInfo.PriceInfo = priceInfo

	teachingLanguages := []int{}
	for _, v := range tutorLanguages {
		teachingLanguages = append(teachingLanguages, v.LanguageId)
	}
	tutorInfo.TeachingLanguages = teachingLanguages
	return tutorInfo
}
