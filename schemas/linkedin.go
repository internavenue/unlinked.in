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
	Industry string      `json:"industry,omitempty"`
}

type RecommendationList struct {
	Total  int              `json:"_total,omitempty"`
	Values []Recommendation `json:"values,omitempty"`
}

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type Recommendation struct {
	ID                 int    `json:"id,omitempty"`
	RecommendationText string `json:"recommendationText,omitempty"`
	RecommendationType string `json:"recommendationType,omitempty"` // just guess
	Recommender        Person `json:"recommender,omitempty"`
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
	Authority Authority `json:"authority,omitempty"`
}

type CertificationList struct {
	Total  int             `json:"_total,omitempty"`
	Values []Certification `json:"values,omitempty"`
}

type CountryCode struct {
	Code string `json:"code,omitempty"`
}

type Location struct {
	Name string      `json:"name,omitempty"`
	Code CountryCode `json:"country,omitempty"`
}

type Publisher struct {
	Name string `json:"name,omitempty"`
}

type AuthorList struct {
	Total  int      `json:"_total,omitempty"`
	Values []Author `json:"values,omitempty"`
}

type Author struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Person Person `json:"person,omitempty"`
}

type Publication struct {
	ID        string     `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	URL       string     `json:"url,omitempty"`
	Summary   string     `json:"summary,omitempty"`
	Date      Date       `json:"date,omitempty"`
	Publisher Publisher  `json:"publisher,omitempty"`
	Authors   AuthorList `json:"authors,omitempty"` // I guess this field.
}

type PublicationList struct {
	Total  int           `json:"_total,omitempty"`
	Values []Publication `json:"values,omitempty"`
}

type Inventor struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Person Person `json:"person,omitempty"`
}

type InventorList struct {
	Total  int        `json:"_total,omitempty"`
	Values []Inventor `json:"values,omitempty"`
}

type StatusCode int

const (
	STATUS_APPLICATION StatusCode = iota
	STATUS_GRANTED
)

type Status struct {
	ID   StatusCode `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`
}

type Office struct {
	Name string `json:"name,omitempty"`
}

type Patent struct {
	ID        string       `json:"id,omitempty"`
	Title     string       `json:"title,omitempty"`
	Summary   string       `json:"summary,omitempty"`
	Number    string       `json:"number,omitempty"`
	Date      Date         `json:"date,omitempty"`
	URL       string       `json:"url,omitempty"`
	Status    Status       `json:"status,omitempty"`
	Office    Office       `json:"office,omitempty"`
	Inventors InventorList `json:"inventors,omitempty"`
}

type PatentList struct {
	Total  int      `json:"_total,omitempty"`
	Values []Patent `json:"values,omitempty"`
}

type Course struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}

type CourseList struct {
	Total  int      `json:"_total,omitempty"`
	Values []Course `json:"values,omitempty"`
}

type Resource struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type ResourceList struct {
	Total  int        `json:"_total,omitempty"`
	Values []Resource `json:"values,omitempty"`
}

type PhoneType string

const (
	PHONETYPE_HOME   PhoneType = "home"
	PHONETYPE_WORK             = "work"
	PHONETYPE_MOBILE           = "mobile"
)

type Phone struct {
	PhoneType   PhoneType `json:"phoneType,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
}

type PhoneList struct {
	Total  int     `json:"_total,omitempty"`
	Values []Phone `json:"values,omitempty"`
}

type BindingStatus string

const (
	BINDINGSTATUS_AUTHENTICATED BindingStatus = "authenticated"
)

type BoundAccount struct {
	ID                  int           `json:"id,omitempty"`
	IsPrimary           bool          `json:"isPrimary,omitempty"`
	AccountType         string        `json:"accountType,omitempty"`
	BindingStatus       BindingStatus `json:"bindingStatus,omitempty"`
	ProviderAccountId   string        `json:"providerAccountId,omitempty"`
	ProviderAccountName string        `json:"providerAccountName,omitempty"`
}

type BoundAccountList struct {
	Total  int            `json:"_total,omitempty"`
	Values []BoundAccount `json:"values,omitempty"`
}

type BoundAccountType struct {
	NumAccounts   int              `json:"numAccounts,omitempty"`
	BoundAccounts BoundAccountList `json:"boundAccounts,omitempty"`
}

// BoundAccountTypeList is the funniest struct
// ...that really what I get from my profile!
type BoundAccountTypeList struct {
	Total  int                `json:"_total,omitempty"`
	Values []BoundAccountType `json:"values,omitempty"`
}

type IMAccountType string

const (
	IMACCOUNTTYPE_SKYPE IMAccountType = "skype"
	IMACCOUNTTYPE_AIM                 = "aim"
	IMACCOUNTTYPE_GTALK               = "gtalk"
	IMACCOUNTTYPE_ICQ                 = "icq"
	IMACCOUNTTYPE_MSN                 = "msn"
	IMACCOUNTTYPE_YAHOO               = "yahoo"
)

type IMAccount struct {
	IMAccountType IMAccountType `json:"imAccountType,omitempty"`
	IMAccountName string        `json:"imAccountName,omitempty"`
}

type IMAccountList struct {
	Total  int         `json:"_total,omitempty"`
	Values []IMAccount `json:"values,omitempty"`
}

type TwitterAccount struct {
	ProviderAccountID   string `json:"providerAccountId,omitempty"`
	ProviderAccountName string `json:"providerAccountName,omitempty"`
}

type TwitterAccountList struct {
	Total  int              `json:"_total,omitempty"`
	Values []TwitterAccount `json:"values,omitempty"`
}

type LinkedInProfile struct {
	ID                          string               `json:"id,omitempty"`
	EmailAddress                string               `json:"emailAddress,omitempty"`
	MainAddress                 string               `json:"mainAddress,omitempty"` // guess
	FirstName                   string               `json:"firstName,omitempty"`
	LastName                    string               `json:"lastName,omitempty"`
	MaidenNamestring            string               `json:"maidenName,omitempty"`
	FormattedNamestring         string               `json:"formattedName,omitempty"`
	PhoneticFirstNamestring     string               `json:"phoneticFirstName,omitempty"`
	PhoneticLastNamestring      string               `json:"phoneticLastName,omitempty"`
	FormattedPhoneticNamestring string               `json:"formattedPhoneticName,omitempty"`
	Headline                    string               `json:"headline,omitempty"`
	Specialties                 string               `json:"specialties,omitempty"`
	Summary                     string               `json:"summary,omitempty"`
	ProposalComments            string               `json:"proposalComments,omitempty"`
	Associations                string               `json:"associations,omitempty"`
	Industry                    string               `json:"industry,omitempty"`
	Interests                   string               `json:"interests,omitempty"`
	PublicProfileUrl            string               `json:"publicProfileUrl,omitempty"`
	PictureUrl                  string               `json:"pictureUrl,omitempty"`
	NumRecommenders             int                  `json:"numRecommenders,omitempty"`
	LastModifiedTimestamp       int64                `json:"lastModifiedTimestamp,omitempty"`
	DateOfBirth                 Date                 `json:"dateOfBirth,omitempty"`
	Location                    Location             `json:"location,omitempty"`
	Volunteer                   Volunteer            `json:"volunteer,omitempty"`
	Recommendations             RecommendationList   `json:"recommendationsReceived,omitempty"`
	Educations                  EducationList        `json:"educations,omitempty"`
	Positions                   PositionList         `json:"positions,omitempty"`
	CurrentPositions            PositionList         `json:"threeCurrentPositions,omitempty"`
	PastPositions               PositionList         `json:"threePastPositions,omitempty"`
	Skills                      SkillList            `json:"skills,omitempty"`
	Languages                   LanguageList         `json:"languages,omitempty"`
	Certifications              CertificationList    `json:"certifications,omitempty"`
	Publications                PublicationList      `json:"publications,omitempty"`
	Courses                     CourseList           `json:"courses,omitempty"`
	Resources                   ResourceList         `json:"resources,omitempty"`
	PhoneNumbers                PhoneList            `json:"phoneNumbers,omitempty"`
	BoundAccountTypes           BoundAccountTypeList `json:"boundAccountTypes,omitempty"`
	IMAccounts                  IMAccountList        `json:"imAccounts,omitempty"`
	TwitterAccounts             TwitterAccountList   `json:"twitterAccounts,omitempty"` // they seem really liek to twitt
	PrimaryTwitterAccount       TwitterAccount       `json:"primaryTwitterAccount,omitempty"`
}
