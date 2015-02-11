package openhr

import (
	"fmt"
	"github.com/internavenue/unlinked.in/schemas"
)

type AccessCredential struct {
	UserID   string `json:",omitempty"`
	TypeCode string `json:",omitempty"`
	Value    string `json:",omitempty"`
}

type Address struct {
	CountryCode            string `json:",omitempty"`
	CountrySubDivisionCode string `json:",omitempty"`
	CityName               string `json:",omitempty"`
	PostalCode             string `json:",omitempty"`
	AddressLine            string `json:",omitempty"`
	Label                  string `json:",omitempty"`
}

type Attachment struct {
	EmbeddedData      string                 `json:",omitempty"`
	EmbeddedText      string                 `json:",omitempty"`
	URI               string                 `json:",omitempty"`
	FileName          string                 `json:",omitempty"`
	Description       string                 `json:",omitempty"`
	FileType          string                 `json:",omitempty"`
	DocumentTitle     string                 `json:",omitempty"`
	AccessCredentials []AccessCredential     `json:",omitempty"`
	UserArea          map[string]interface{} `json:",omitempty"`
}

type Certification struct {
	CertificationTypeCode string `json:",omitempty"`
	CertificationName     string `json:",omitempty"`
	Description           string `json:",omitempty"`
	IssuingAuthorityName  string `json:",omitempty"`
	FirstIssuedDate       string `json:",omitempty"`
	EndDate               string `json:",omitempty"`
}

type DegreeType struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type EducationLevel struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type AttendanceStatusCode string

const (
	STATUS_CURRENT       AttendanceStatusCode = "Current"
	STATUS_PRIOR                              = "Prior"
	STATUS_UNKNOWN                            = "Unknown"
	STATUS_NOT_SPECIFIED                      = "NotSpecified"
)

type Education struct {
	School               string               `json:",omitempty"`
	SubSchool            []string             `json:",omitempty"`
	ReferenceLocation    ReferenceLocation    `json:",omitempty"`
	EducationLevel       []EducationLevel     `json:",omitempty"`
	AttendanceStatusCode AttendanceStatusCode `json:",omitempty"`
	AttendanceStartDate  string               `json:",omitempty"`
	AttendanceEndDate    string               `json:",omitempty"`
	EducationScore       []string             `json:",omitempty"`
	DegreeType           []DegreeType         `json:",omitempty"`
	DegreeDate           string               `json:",omitempty"`
	MajorProgramName     []string             `json:",omitempty"`
	MinorProgramName     []string             `json:",omitempty"`
	AcademicHonors       string               `json:",omitempty"`
	Comment              string               `json:",omitempty"`
	UserArea             map[string]interface{}
}

type Preferred string

const (
	PRIMARY   Preferred = "primary"
	SECONDARY           = "secondary"
)

type Label string

const (
	LABEL_PERSONAL Label = "personal"
	LABEL_WORK           = "work"
	LABEL_MOBILE         = "mobile"
	LABEL_OTHER          = "other"
	LABEL_BLOG           = "blog"
	LABEL_SOCIAL         = "social"
)

type Email struct {
	Address   string    `json:",omitempty"`
	Label     Label     `json:",omitempty"`
	Preferred Preferred `json:",omitempty"`
}

type RefereeTypeCode string

const (
	PROFESSIONAL RefereeTypeCode = "Professional"
	PERSONAL                     = "Personal"
	VERIFICATION                 = "Verification"
)

type EmploymentReference struct {
	RefereeTypeCode      RefereeTypeCode `json:",omitempty"`
	FormattedName        string          `json:",omitempty"`
	PositionTitle        string          `json:",omitempty"`
	OrganizationName     string          `json:",omitempty"`
	PreferredPhone       string          `json:",omitempty"`
	PreferredEmail       string          `json:",omitempty"`
	YearsKnownNumber     float32         `json:",omitempty"`
	RelationshipTypeCode string          `json:",omitempty"`
	Comment              string          `json:",omitempty"`
}

type Gender string

const (
	MALE    Gender = "Male"
	FEMALE         = "Female"
	UNKNOWN        = "Unknown"
)

type Industry struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type JobCategory struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type JobLevel struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type License struct {
	LicenseTypeCode      string `json:",omitempty"`
	LicenseName          string `json:",omitempty"`
	Description          string `json:",omitempty"`
	IssuingAuthorityName string `json:",omitempty"`
	FirstIssuedDate      string `json:",omitempty"`
	EndDate              string `json:",omitempty"`
}

type PersonCompetency struct {
	CompetencyID    string `json:",omitempty"`
	CompetencyName  string `json:",omitempty"`
	CompetencyLevel string `json:",omitempty"`
}

type Phone struct {
	Number    string    `json:",omitempty"`
	Label     Label     `json:",omitempty"`
	Preferred Preferred `json:",omitempty"`
}

type PositionHistory struct {
	Employer             string                 `json:",omitempty"`
	OrganizationUnitName string                 `json:",omitempty"`
	PositionTitle        string                 `json:",omitempty"`
	ReferenceLocation    ReferenceLocation      `json:",omitempty"`
	StartDate            string                 `json:",omitempty"`
	EndDate              string                 `json:",omitempty"`
	CurrentIndicator     bool                   `json:",omitempty"`
	Industry             []Industry             `json:",omitempty"`
	JobCategory          []JobCategory          `json:",omitempty"`
	JobLevel             []JobLevel             `json:",omitempty"`
	Description          string                 `json:",omitempty"`
	UserArea             map[string]interface{} `json:",omitempty"`
}

type ReferenceLocation struct {
	CountryCode            string   `json:",omitempty"`
	CountrySubDivisionCode []string `json:",omitempty"`
	CityName               string   `json:",omitempty"`
}

type SpecialCommitment struct {
	Name string `json:",omitempty"`
	Code string `json:",omitempty"`
}

type Web struct {
	Address          string `json:",omitempty"`
	Label            Label  `json:",omitempty"`
	LabelDescription string `json:",omitempty"`
}

type WorkEligibility struct {
	CountryCode string `json:",omitempty"`
	Permanent   bool   `json:",omitempty"`
}

type OpenHRSchema struct {
	CandidateID          string                 `json:",omitempty"`
	CandidateURI         string                 `json:",omitempty"`
	GivenName            string                 `json:",omitempty"`
	FamilyName           string                 `json:",omitempty"`
	FormattedName        string                 `json:",omitempty"`
	DateOfBirth          string                 `json:",omitempty"`
	Gender               Gender                 `json:",omitempty"`
	DisabilityIndicator  bool                   `json:",omitempty"`
	DisabilitySummary    []string               `json:",omitempty"`
	Address              []Address              `json:",omitempty"`
	Phone                []Phone                `json:",omitempty"`
	Email                []Email                `json:",omitempty"`
	Web                  []Web                  `json:",omitempty"`
	PersonCompetency     []PersonCompetency     `json:",omitempty"`
	WorkEligibility      []WorkEligibility      `json:",omitempty"`
	Education            []Education            `json:"EducationOrganizationAttendance,omitempty"`
	PositionHistory      []PositionHistory      `json:",omitempty"`
	Certification        []Certification        `json:",omitempty"`
	License              []License              `json:",omitempty"`
	EmploymentReferences []EmploymentReference  `json:",omitempty"`
	SpecialCommitment    []SpecialCommitment    `json:",omitempty"`
	Attachment           []Attachment           `json:",omitempty"`
	UserArea             map[string]interface{} `json:",omitempty"`
}

func toStringDate(date schemas.Date) string {
	var strDate string
	if date.Day > 0 {
		strDate = fmt.Sprintf("%d", date.Day)
	}
	if date.Month > 0 {
		strDate = fmt.Sprintf("%d-%s", date.Month, strDate)
	}
	if date.Year > 0 {
		strDate = fmt.Sprintf("%d-%s", date.Year, strDate)
	}

	return strDate
}

func FromLinkedInSchema(linked *schemas.LinkedInSchema) *OpenHRSchema {
	p := &OpenHRSchema{}
	// TODO(nvcnvn): handle commented field and consider when to use UserArea
	p.CandidateID = linked.ID
	p.CandidateURI = linked.PublicProfileUrl
	p.GivenName = linked.FirstName
	p.FamilyName = linked.LastName
	p.FormattedName = linked.FormattedName
	// p.FormattedName
	p.DateOfBirth = toStringDate(linked.DateOfBirth)
	// p.Gender
	// p.DisabilityIndicator
	// p.DisabilitySummary
	p.Address = []Address{Address{AddressLine: linked.MainAddress}}

	p.Phone = make([]Phone, 0, len(linked.PhoneNumbers.Values))
	for _, fone := range linked.PhoneNumbers.Values {
		switch fone.PhoneType {
		case schemas.PHONETYPE_HOME:
			p.Phone = append(p.Phone, Phone{
				Number: fone.PhoneNumber,
				Label:  LABEL_PERSONAL,
			})
		case schemas.PHONETYPE_MOBILE:
			p.Phone = append(p.Phone, Phone{
				Number: fone.PhoneNumber,
				Label:  LABEL_MOBILE,
			})
		case schemas.PHONETYPE_WORK:
			p.Phone = append(p.Phone, Phone{
				Number: fone.PhoneNumber,
				Label:  LABEL_WORK,
			})
		default:
			p.Phone = append(p.Phone, Phone{
				Number: fone.PhoneNumber,
				Label:  LABEL_OTHER,
			})
		}
	}
	p.Email = []Email{Email{Address: linked.EmailAddress}}

	p.Web = make([]Web, 0, linked.Resources.Total)
	for _, rs := range linked.Resources.Values {
		p.Web = append(p.Web, Web{Address: rs.URL, LabelDescription: rs.Name})
	}

	p.PersonCompetency = make([]PersonCompetency, 0,
		linked.Languages.Total+linked.Skills.Total)
	for _, lang := range linked.Languages.Values {
		p.PersonCompetency = append(p.PersonCompetency, PersonCompetency{
			CompetencyID:    fmt.Sprintf("linkedin-lang#%d", lang.ID),
			CompetencyName:  lang.Language.Name,
			CompetencyLevel: lang.Proficiency.Name,
		})
	}
	for _, skill := range linked.Skills.Values {
		p.PersonCompetency = append(p.PersonCompetency, PersonCompetency{
			CompetencyID:   fmt.Sprintf("linkedin-skill#%d", skill.ID),
			CompetencyName: skill.Skill.Name,
		})
	}

	p.Education = make([]Education, 0, linked.Educations.Total)
	for _, edu := range linked.Educations.Values {
		p.Education = append(p.Education, Education{
			School: edu.SchoolName,
			EducationLevel: []EducationLevel{EducationLevel{
				Name: edu.Degree,
			}},
			AttendanceStartDate: toStringDate(edu.StartDate),
			AttendanceEndDate:   toStringDate(edu.EndDate),
			MajorProgramName:    []string{edu.FieldOfStudy},
		})
	}

	p.PositionHistory = make([]PositionHistory, 0,
		linked.CurrentPositions.Total+linked.PastPositions.Total)
	for _, pos := range linked.CurrentPositions.Values {
		p.PositionHistory = append(p.PositionHistory, PositionHistory{
			Employer:         pos.Company.Name,
			PositionTitle:    pos.Title,
			Description:      pos.Summary,
			StartDate:        toStringDate(pos.StartDate),
			EndDate:          toStringDate(pos.EndDate),
			CurrentIndicator: pos.IsCurrent,
			Industry:         []Industry{Industry{Name: pos.Company.Industry}},
		})
	}

	for _, pos := range linked.PastPositions.Values {
		p.PositionHistory = append(p.PositionHistory, PositionHistory{
			Employer:         pos.Company.Name,
			PositionTitle:    pos.Title,
			Description:      pos.Summary,
			StartDate:        toStringDate(pos.StartDate),
			EndDate:          toStringDate(pos.EndDate),
			CurrentIndicator: pos.IsCurrent,
			Industry:         []Industry{Industry{Name: pos.Company.Industry}},
		})
	}

	p.Certification = make([]Certification, 0, linked.Certifications.Total)
	for _, cert := range linked.Certifications.Values {
		p.Certification = append(p.Certification, Certification{
			CertificationTypeCode: cert.Number,
			CertificationName:     cert.Name,
			IssuingAuthorityName:  cert.Authority.Name,
			FirstIssuedDate:       toStringDate(cert.StartDate),
			EndDate:               toStringDate(cert.EndDate),
		})
	}

	p.EmploymentReferences = make([]EmploymentReference, 0, linked.Recommendations.Total)
	for _, ref := range linked.Recommendations.Values {
		p.EmploymentReferences = append(p.EmploymentReferences, EmploymentReference{
			Comment:       ref.RecommendationText,
			FormattedName: ref.Recommender.FirstName + " " + ref.Recommender.LastName,
			// RefereeTypeCode
		})
	}

	p.Attachment = make([]Attachment, 0, linked.Resources.Total)
	for _, rs := range linked.Resources.Values {
		p.Attachment = append(p.Attachment, Attachment{
			URI:      rs.URL,
			FileName: rs.Name,
		})
	}

	return p
}
