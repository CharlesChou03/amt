package service

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/CharlesChou03/_git/amt.git/internal/db"
	modelsRes "github.com/CharlesChou03/_git/amt.git/models/response"
	"github.com/CharlesChou03/_git/amt.git/services"
)

func TestAggregateTutorInfo(t *testing.T) {

	t.Run("assert equal", func(t *testing.T) {
		tutorEntityId := 1
		slug := "foo-bar"
		name := "Amazing Teacher 1"
		headline := "Hi I'm a English Teacher"
		introduction := "........."
		normalPrice := float32(10)
		trialPrice := float32(5)

		tutor := tutorEntity(tutorEntityId, slug, name, headline, introduction)

		tutorLessonEntityId := 1
		tutorLessonPrice := tutorLessonPriceEntity(tutorLessonEntityId, tutorEntityId, normalPrice, trialPrice)

		tutorLanguages := []db.TutorLanguage{}
		tutorLanguageEntityId := 1
		tutorLanguageId := 121
		tutorLanguage1 := tutorLanguageEntity(tutorLanguageEntityId, tutorEntityId, tutorLanguageId)
		tutorLanguages = append(tutorLanguages, tutorLanguage1)
		tutorLanguageEntityId = 2
		tutorLanguageId = 123
		tutorLanguage2 := tutorLanguageEntity(tutorLanguageEntityId, tutorEntityId, tutorLanguageId)
		tutorLanguages = append(tutorLanguages, tutorLanguage2)

		wantTeachingLanguages := []int{121, 123}
		want := getTutorResEntity(tutorEntityId, slug, name, headline, introduction, normalPrice, trialPrice, wantTeachingLanguages)

		got := services.AggregateTutorInfo(tutor, tutorLessonPrice, tutorLanguages)
		assertEqual(t, got, want)
	})

	t.Run("assert not equal", func(t *testing.T) {
		tutorEntityId := 1
		slug := "foo-bar"
		name1 := "Amazing Teacher 1"
		name2 := "Amazing Teacher 2"
		headline := "Hi I'm a English Teacher"
		introduction := "........."
		normalPrice := float32(10)
		trialPrice := float32(5)

		tutor := tutorEntity(tutorEntityId, slug, name1, headline, introduction)

		tutorLessonEntityId := 1
		tutorLessonPrice := tutorLessonPriceEntity(tutorLessonEntityId, tutorEntityId, normalPrice, trialPrice)

		tutorLanguages := []db.TutorLanguage{}
		tutorLanguageEntityId := 1
		tutorLanguageId := 121
		tutorLanguage1 := tutorLanguageEntity(tutorLanguageEntityId, tutorEntityId, tutorLanguageId)
		tutorLanguages = append(tutorLanguages, tutorLanguage1)
		tutorLanguageEntityId = 2
		tutorLanguageId = 123
		tutorLanguage2 := tutorLanguageEntity(tutorLanguageEntityId, tutorEntityId, tutorLanguageId)
		tutorLanguages = append(tutorLanguages, tutorLanguage2)

		wantTeachingLanguages := []int{121, 123}
		want := getTutorResEntity(tutorEntityId, slug, name2, headline, introduction, normalPrice, trialPrice, wantTeachingLanguages)

		got := services.AggregateTutorInfo(tutor, tutorLessonPrice, tutorLanguages)
		assertNotEqual(t, got, want)
	})

}

func tutorEntity(id int, slug, name, headline, introduction string) db.Tutor {
	tutor := db.Tutor{}
	tutor.ID = id
	tutor.Slug = slug
	tutor.Name = name
	tutor.Headline = headline
	tutor.Introduction = introduction
	return tutor
}

func tutorLessonPriceEntity(id, tutorId int, normalPrice, trialPrice float32) db.TutorLessonPrice {
	tutorLessonPrice := db.TutorLessonPrice{}
	tutorLessonPrice.ID = id
	tutorLessonPrice.TutorId = tutorId
	tutorLessonPrice.NormalPrice = normalPrice
	tutorLessonPrice.TrialPrice = trialPrice
	return tutorLessonPrice
}

func tutorLanguageEntity(id, tutorId, languageId int) db.TutorLanguage {
	tutorLanguage := db.TutorLanguage{}
	tutorLanguage.ID = id
	tutorLanguage.TutorId = tutorId
	tutorLanguage.LanguageId = languageId
	return tutorLanguage
}

func getTutorResEntity(tutorId int,
	slug, name, headline, introduction string,
	normalPrice, trialPrice float32,
	teachingLanguages []int) modelsRes.GetTutorRes {

	getTutorRes := modelsRes.GetTutorRes{}
	getTutorRes.ID = strconv.Itoa(tutorId)
	getTutorRes.Slug = slug
	getTutorRes.Name = name
	getTutorRes.Headline = headline
	getTutorRes.Introduction = introduction
	priceInfo := modelsRes.PriceInfo{}
	priceInfo.Trial = trialPrice
	priceInfo.Normal = normalPrice
	getTutorRes.PriceInfo = priceInfo
	getTutorRes.TeachingLanguages = teachingLanguages
	return getTutorRes
}

func assertEqual(t *testing.T, got, want modelsRes.GetTutorRes) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func assertNotEqual(t *testing.T, got, want modelsRes.GetTutorRes) {
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
