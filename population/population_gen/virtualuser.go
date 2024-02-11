package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type VirtualUser struct {
	ID                      string                  `yaml:"id"`
	PersonalInfo            PersonalInfo            `yaml:"personal_info"`
	PhysicalCharacteristics PhysicalCharacteristics `yaml:"physical_characteristics"`
	HealthProfile           HealthProfile           `yaml:"health_profile"`
	Education               []Education             `yaml:"education"`
	Career                  Career                  `yaml:"career"`
	Personality             Personality             `yaml:"personality"`
	Lifestyle               Lifestyle               `yaml:"lifestyle"`
	Relationships           Relationships           `yaml:"relationships"`
	Goals                   Goals                   `yaml:"goals"`
	Preferences             Preferences             `yaml:"preferences"`
}

type PersonalInfo struct {
	Name            string   `yaml:"name"`
	Gender          string   `yaml:"gender"`
	Birthday        string   `yaml:"birthday"`
	Birthplace      string   `yaml:"birthplace"`
	Residence       string   `yaml:"residence"`
	Nationality     string   `yaml:"nationality"`
	LanguagesSpoken []string `yaml:"languages_spoken"`
}

type PhysicalCharacteristics struct {
	Height                 string   `yaml:"height"`
	Weight                 string   `yaml:"weight"`
	EyeColor               string   `yaml:"eye_color"`
	HairColor              string   `yaml:"hair_color"`
	DistinguishingFeatures []string `yaml:"distinguishing_features"`
}

type HealthProfile struct {
	BloodType          string   `yaml:"blood_type"`
	Allergies          []string `yaml:"allergies"`
	MedicalConditions  []string `yaml:"medical_conditions"`
	RegularMedications []string `yaml:"regular_medications"`
}

type Education struct {
	Degree        string `yaml:"degree"`
	Institution   string `yaml:"institution"`
	YearCompleted int    `yaml:"year_completed"`
}

type Career struct {
	CurrentPosition   string             `yaml:"current_position"`
	Employer          string             `yaml:"employer"`
	YearsAtCompany    int                `yaml:"years_at_company"`
	PreviousPositions []PreviousPosition `yaml:"previous_positions"`
}

type PreviousPosition struct {
	Title    string `yaml:"title"`
	Company  string `yaml:"company"`
	Duration string `yaml:"duration"`
}

type Personality struct {
	MBTI           string   `yaml:"mbti"`
	BigFive        []string `yaml:"big_five"`
	Strengths      []string `yaml:"strengths"`
	Weaknesses     []string `yaml:"weaknesses"`
	Fears          []string `yaml:"fears"`
	Motivations    []string `yaml:"motivations"`
	Interests      []string `yaml:"interests"`
	Values         []string `yaml:"values"`
	LifePhilosophy string   `yaml:"life_philosophy"`
}

type Lifestyle struct {
	DietaryPreferences []string `yaml:"dietary_preferences"`
	ExerciseRoutine    []string `yaml:"exercise_routine"`
	Hobbies            []string `yaml:"hobbies"`
	FavoriteBooks      []string `yaml:"favorite_books"`
	FavoriteMovies     []string `yaml:"favorite_movies"`
}

type Relationships struct {
	Family           []FamilyMember   `yaml:"family"`
	SignificantOther SignificantOther `yaml:"significant_other"`
	CloseFriends     []Friend         `yaml:"close_friends"`
}

type FamilyMember struct {
	Relation string `yaml:"relation"`
	Name     string `yaml:"name"`
}

type SignificantOther struct {
	Name                 string `yaml:"name"`
	RelationshipDuration string `yaml:"relationship_duration"`
}

type Friend struct {
	Name string `yaml:"name"`
}

type Goals struct {
	ShortTerm []string `yaml:"short_term"`
	LongTerm  []string `yaml:"long_term"`
}

type Preferences struct {
	MusicGenres   []string `yaml:"music_genres"`
	Colors        []string `yaml:"colors"`
	PetPeeves     []string `yaml:"pet_peeves"`
	DreamVacation string   `yaml:"dream_vacation"`
}

func main() {
	// Example of initializing a VirtualUser and printing the YAML
	user := VirtualUser{
		ID: "1001",
		PersonalInfo: PersonalInfo{
			Name: "Jordan Rivera",
			// Additional fields here
		},
		// Populate other fields as necessary
	}

	data, err := yaml.Marshal(&user)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(data))
}
