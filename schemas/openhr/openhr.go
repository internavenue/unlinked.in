package openhr

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

type OpenHRProfile struct {
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
