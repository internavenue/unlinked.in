package europass

import (
	"github.com/internavenue/unlinked.in/schemas"
)

type CodeLabel struct {
	Code  string `json:",omitempty"`
	Label string `json:",omitempty"`
}

const (
	ACHIEVEMENT_HONORS_AWARDS        = "honors_awards"
	ACHIEVEMENT_PUBLICATIONS         = "publications"
	ACHIEVEMENT_PRESENTATIONS        = "presentations"
	ACHIEVEMENT_PROJECTS             = "projects"
	ACHIEVEMENT_CITATIONS            = "citations"
	ACHIEVEMENT_MEMBERSHIPS          = "memberships"
	ACHIEVEMENT_CONFERENCES          = "conferences"
	ACHIEVEMENT_SEMINARS             = "seminars"
	ACHIEVEMENT_WORKSHOPS            = "workshops"
	ACHIEVEMENT_REFERENCES           = "references"
	ACHIEVEMENT_SIGNATURE_EQUIVALENT = "signature_equivalent"
)

type Achievement struct {
	Description string    `json:",omitempty"`
	Title       CodeLabel `json:",omitempty"`
	ReferenceToList
}

type Address struct {
	AddressLine  string `json:",omitempty"`
	AddressLine2 string `json:",omitempty"`
	Municipality string `json:",omitempty"`
	PostalCode   string `json:",omitempty"`
	Country      CodeLabel
}

type ContactAddress struct {
	Contact Address `json:",omitempty"`
}

type Email struct {
	Contact string `json:",omitempty"`
}

const (
	INSTANTMESSAGING_GTALK = "gtalk"
	INSTANTMESSAGING_SKYPE = "skype"
	INSTANTMESSAGING_ICQ   = "icq"
	INSTANTMESSAGING_AIM   = "aim"
	INSTANTMESSAGING_MSN   = "msn"
	INSTANTMESSAGING_YAHOO = "yahoo"
)

type InstantMessaging struct {
	Contact string `json:",omitempty"`
	Use     CodeLabel
}

const (
	TELEPHONE_HOME   = "home"
	TELEPHONE_WORK   = "work"
	TELEPHONE_MOBILE = "mobile"
)

type Telephone struct {
	Contact string `json:",omitempty"`
	Use     CodeLabel
}

const (
	WEBSITE_PERSONAL  = "personal"
	WEBSITE_BUSINESS  = "business"
	WEBSITE_WORK      = "work"
	WEBSITE_PORTFOLIO = "portfolio"
)

type Website struct {
	Contact string `json:",omitempty"`
	Use     CodeLabel
}

type ContactInfo struct {
	Address          ContactAddress
	Email            Email
	InstantMessaging []InstantMessaging
	Telephone        []Telephone
	Website          []Website
}

type OrganisationContactInfo struct {
	Address ContactAddress
	Website Website
}

type Organisation struct {
	ContactInfo OrganisationContactInfo
	Name        string `json:",omitempty"`
}

type Date struct {
	Year  int `json:",omitempty"`
	Month int `json:",omitempty"`
	Day   int `json:",omitempty"`
}

type Period struct {
	Current bool `json:",omitempty"`
	From    Date
	To      Date
}

type Education struct {
	BasicExperience
	Organisation Organisation
	Title        string `json:",omitempty"`
	Activities   string `json:",omitempty"`
	Field        CodeLabel
	Level        CodeLabel
}

const (
	HEADLINE_PREFERRED_JOB       = "preferred_job"
	HEADLINE_JOB_APPLIED_FOR     = "job_applied_for"
	HEADLINE_STUDIES_APPLIED_FOR = "studies_applied_for"
	HEADLINE_POSITION            = "position"
	HEADLINE_PERSONAL_STATEMENT  = "personal_statement"
)

type Headline struct {
	Description CodeLabel `json:",omitempty"`
	Type        CodeLabel `json:",omitempty"`
}

const (
	GENDER_MALE   = "M"
	GENDER_FEMALE = "F"
)

type Demographics struct {
	Birthdate   Date      `json:",omitempty"`
	Gender      CodeLabel `json:",omitempty"`
	Nationality CodeLabel `json:",omitempty"`
}

const (
	PERSON_TITLE_MR   = "mr"
	PERSON_TITLE_MS   = "ms"
	PERSON_TITLE_MRS  = "mrs"
	PERSON_TITLE_MISS = "miss"
	PERSON_TITLE_DR   = "dr"
)

type PersonName struct {
	Title     CodeLabel
	FirstName string `json:",omitempty"`
	Surname   string `json:",omitempty"`
}

type Metadata struct {
	Key   string `json:",omitempty"`
	Value string `json:",omitempty"`
}

type FileData struct {
	ID      string     `json:"id,omitempty"`
	Name    string     `json:",omitempty"`
	Data    string     `json:",omitempty"`
	Metadta []Metadata `json:",omitempty"`
}

const (
	PHOTO_JPEG  = "image/jpeg"
	PHOTO_PJPEG = "image/pjpeg"
	PHOTO_PNG   = "image/png"
	PHOTO_XPNG  = "image/x-png"
)

type Photo struct {
	FileData
	MimeType string
}

type Identification struct {
	ContactInfo  ContactInfo
	Demographics Demographics
	PersonName   PersonName `json:",omitempty"`
	Photo        Photo
}

type Reference struct {
	IDRef string `json:"idref,omitempty"`
}

type ReferenceToList struct {
	ReferenceTo []Reference
}

type MotherTongue struct {
	CodeLabel
	ReferenceToList
}

const (
	LEVEL_A1 = "A1"
	LEVEL_A2 = "A2"
	LEVEL_B1 = "B1"
	LEVEL_B2 = "B2"
	LEVEL_C1 = "C1"
	LEVEL_C2 = "C2"
)

type Certificate struct {
	Title        string `json:",omitempty"`
	AwardingBody string
	Date         Date
	Level        string
}

type ProficiencyLevel struct {
	Listening         string `json:",omitempty"`
	Reading           string `json:",omitempty"`
	SpokenInteraction string `json:",omitempty"`
	SpokenProduction  string `json:",omitempty"`
	Writing           string `json:",omitempty"`
}

type BasicExperience struct {
	Period
	Description string
	ReferenceToList
}

const (
	LANGUAGE_AREA_STUDYING         = "studying_training_language"
	LANGUAGE_AREA_WORK             = "work_language"
	LANGUAGE_AREA_LIVING_TRAVELING = "living_traveling_language"
	LANGUAGE_AREA_MEDIATING_GROUPS = "mediating_groups_language"
)

type LanguageExperience struct {
	BasicExperience
	Area CodeLabel
}

type ForeignLanguage struct {
	CodeLabel
	Certificate      []Certificate
	Description      CodeLabel
	ProficiencyLevel ProficiencyLevel
	Experience       LanguageExperience
}

type Linguistic struct {
	ForeignLanguage []ForeignLanguage `json:",omitempty"`
	MotherTongue    []CodeLabel
}

type GenericSkill struct {
	ReferenceToList
	Description string `json:",omitempty"`
}

type DrivingSkill struct {
	Description []string `json:",omitempty"`
	ReferenceToList
}

type Skills struct {
	Linguistic     Linguistic
	Driving        DrivingSkill
	Communication  GenericSkill
	Computer       GenericSkill
	Organisational GenericSkill
	JobRelated     GenericSkill
	Other          GenericSkill
}

type Employer struct {
	Organisation
	Sector CodeLabel
	ReferenceToList
}

type WorkExperience struct {
	BasicExperience
	Activities string `json:",omitempty"`
	Employer   Employer
	Position   CodeLabel
}

type EuropassSchema struct {
	Achievement    []Achievement `json:",omitempty"`
	Education      []Education   `json:",omitempty"`
	Headline       Headline
	Identification Identification
	Skills         Skills
	WorkExperience []WorkExperience `json:",omitempty"`
}

func toDate(d schemas.Date) Date {
	return Date{
		Day:   d.Day,
		Month: d.Month,
		Year:  d.Year,
	}
}

func FromLinkedInSchema(linked *schemas.LinkedInSchema) *EuropassSchema {
	s := &EuropassSchema{}
	s.Headline.Description.Label = linked.Headline
	s.Achievement = make([]Achievement, 0,
		linked.HonorsAwards.Total+linked.Publications.Total+linked.Recommendations.Total)
	for _, ha := range linked.HonorsAwards.Values {
		s.Achievement = append(s.Achievement, Achievement{
			Description: ha.Name,
			Title:       CodeLabel{Code: ACHIEVEMENT_HONORS_AWARDS},
		})
	}

	for _, p := range linked.Publications.Values {
		s.Achievement = append(s.Achievement, Achievement{
			Description: p.Title,
			Title:       CodeLabel{Code: ACHIEVEMENT_PUBLICATIONS},
		})
	}

	for _, r := range linked.Recommendations.Values {
		s.Achievement = append(s.Achievement, Achievement{
			Description: r.RecommendationText,
			Title:       CodeLabel{Code: ACHIEVEMENT_REFERENCES},
		})
	}

	s.Education = make([]Education, 0, linked.Educations.Total)
	for _, e := range linked.Educations.Values {
		s.Education = append(s.Education, Education{
			BasicExperience: BasicExperience{
				Period:      Period{From: toDate(e.StartDate), To: toDate(e.EndDate)},
				Description: e.Notes,
			},
			Title:        e.Degree,
			Activities:   e.Activities,
			Field:        CodeLabel{Label: e.FieldOfStudy},
			Organisation: Organisation{Name: e.SchoolName},
			Level:        CodeLabel{Label: e.Degree},
		})
	}

	ctInf := ContactInfo{}
	ctInf.Address.Contact = Address{
		AddressLine: linked.MainAddress,
		Country: CodeLabel{
			Label: linked.Location.Name,
			Code:  linked.Location.Country.Code,
		},
	}

	ctInf.Email.Contact = linked.EmailAddress

	ctInf.InstantMessaging = make([]InstantMessaging, 0, linked.IMAccounts.Total)
	for _, a := range linked.IMAccounts.Values {
		ctInf.InstantMessaging = append(ctInf.InstantMessaging, InstantMessaging{
			Contact: a.IMAccountName,
			Use:     CodeLabel{Code: string(a.IMAccountType)},
		})
	}

	ctInf.Telephone = make([]Telephone, 0, linked.PhoneNumbers.Total)
	for _, p := range linked.PhoneNumbers.Values {
		ctInf.Telephone = append(ctInf.Telephone, Telephone{
			Contact: p.PhoneNumber,
			Use:     CodeLabel{Code: string(p.PhoneType)},
		})
	}

	s.Identification.ContactInfo = ctInf
	langs := make([]ForeignLanguage, 0,
		linked.Languages.Total)
	for _, l := range linked.Languages.Values {
		langs = append(langs, ForeignLanguage{
			CodeLabel: CodeLabel{Label: l.Language.Name},
		})
	}

	s.Identification.Demographics.Birthdate = toDate(linked.DateOfBirth)
	s.Identification.PersonName.FirstName = linked.FirstName
	s.Identification.PersonName.Surname = linked.LastName

	s.Skills.Linguistic.ForeignLanguage = langs
	// TODO(nvcnvn): not sure what I should do with skills

	s.WorkExperience = make([]WorkExperience, 0, linked.Positions.Total)
	for _, p := range linked.Positions.Values {
		s.WorkExperience = append(s.WorkExperience, WorkExperience{
			BasicExperience: BasicExperience{
				Period:      Period{From: toDate(p.StartDate), To: toDate(p.EndDate)},
				Description: p.Title,
			},
			Employer: Employer{
				Organisation: Organisation{
					Name: p.Company.Name,
				},
				Sector: CodeLabel{Label: p.Company.Industry},
			},
			Activities: p.Summary,
		})
	}

	return s
}
