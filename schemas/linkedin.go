package schemas

type Date struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type EducationList struct {
	Total  int         `json:"_total,omitempty"`
	Values []Education `json:"values,omitempty"`
}

type Education struct {
	Degree       string `json:"degree,omitempty"`
	EndDate      Date   `json:"endDate,omitempty"`
	FieldOfStudy string `json:"fieldOfStudy,omitempty"`
	ID           int    `json:"id,omitempty"`
	Notes        string `json:"notes,omitempty"`
	SchoolName   string `json:"schoolName,omitempty"`
	StartDate    Date   `json:"startDate,omitempty"`
	Activities   string `json:"activities,omitempty"`
}

type LanguageList struct {
	Total  int            `json:"_total,omitempty"`
	Values []LanguageData `json:"values,omitempty"`
}

type Language struct {
	Name string `json:"name,omitempty"`
}

type Proficiency struct {
	Level string `json:"level,omitempty"`
	Name  string `json:"name,omitempty"`
}

type LanguageData struct {
	ID          int         `json:"id,omitempty"`
	Language    Language    `json:"language,omitempty"`
	Proficiency Proficiency `json:"proficiency,omitempty"`
}

type Company struct {
	ID       int         `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Ticker   string      `json:"ticker,omitempty"`
	Type     interface{} `json:"type,omitempty"`
	Industry interface{} `json:"industry,omitempty"`
}

type RecommendationList struct {
	Total  int              `json:"_total,omitempty"`
	Values []Recommendation `json:"values,omitempty"`
}

type Recommendation struct {
	ID                 int    `json:"id,omitempty"`
	RecommendationText string `json:"recommendation-text,omitempty"`
	// handle recommendation-type and recommender
}

type SkillList struct {
	Total  int         `json:"_total,omitempty"`
	Values []SKillData `json:"values,omitempty"`
}

type SKillData struct {
	ID    int   `json:"id,omitempty"`
	Skill Skill `json:"skill,omitempty"`
}

type Skill struct {
	Name string `json:"name,omitempty"`
}

type PositionList struct {
	Total  int        `json:"_total,omitempty"`
	Values []Position `json:"values,omitempty"`
}

type Position struct {
	Company   Company `json:"company,omitempty"`
	ID        int     `json:"id,omitempty"`
	IsCurrent bool    `json:"isCurrent,omitempty"`
	StartDate Date    `json:"startDate,omitempty"`
	EndDate   Date    `json:"endDate,omitempty"`
	Summary   string  `json:"summary,omitempty"`
	Title     string  `json:"title,omitempty"`
}

type CauseList struct {
	Total  int     `json:"_total,omitempty"`
	Values []Cause `json:"values,omitempty"`
}

type Cause struct {
	Name string `json:"name,omitempty"`
}

type Opportunity struct {
	BoardMember bool `json:"boardMember,omitempty"`
	ProBono     bool `json:"proBono,omitempty"`
}

type VolunteerExperienceList struct {
	Total  int                   `json:"_total,omitempty"`
	Values []VolunteerExperience `json:"values,omitempty"`
}

type Organization struct {
	Name string `json:"name,omitempty"`
}

type VolunteerExperience struct {
	ID           int          `json:"id,omitempty"`
	Organization Organization `json:"organization,omitempty"`
	Role         string       `json:"role,omitempty"`
}

type Volunteer struct {
	Causes               CauseList               `json:"causes,omitempty"`
	Opportunities        Opportunity             `json:"opportunities,omitempty"`
	VolunteerExperiences VolunteerExperienceList `json:"volunteerExperiences,omitempty"`
}

type Authority struct {
	Name string `json:"name,omitempty"`
}

type Certification struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Number    string    `json:"number,omitempty"`
	StartDate Date      `json:"startDate,omitempty"`
	EndDate   Date      `json:"endDate,omitempty"`
	Authority Authority `authority,omitempty"`
}

type CertificationList struct {
	Total  int             `json:"_total,omitempty"`
	Values []Certification `json:"values,omitempty"`
}

type LinkedInProfile struct {
	ID               string             `json:"id,omitempty"`
	DateOfBirth      Date               `json:"dateOfBirth,omitempty"`
	Educations       EducationList      `json:"educations,omitempty"`
	FirstName        string             `json:"firstName,omitempty"`
	Headline         string             `json:"headline,omitempty"`
	Industry         string             `json:"industry,omitempty"`
	Interests        string             `json:"interests,omitempty"`
	Languages        LanguageList       `json:"languages,omitempty"`
	LastName         string             `json:"lastName,omitempty"`
	NumRecommenders  int                `json:"numRecommenders,omitempty"`
	PictureUrl       string             `json:"pictureUrl,omitempty"`
	Positions        PositionList       `json:"positions,omitempty"`
	Recommendations  RecommendationList `json:"recommendationsReceived,omitempty"`
	Skills           SkillList          `json:"skills,omitempty"`
	Summary          string             `json:"summary,omitempty"`
	CurrentPositions PositionList       `json:"threeCurrentPositions,omitempty"`
	PastPositions    PositionList       `json:"threePastPositions,omitempty"`
	Volunteer        Volunteer          `json:"volunteer,omitempty"`
	Certifications   CertificationList  `json:"certifications,omitempty"`
}
