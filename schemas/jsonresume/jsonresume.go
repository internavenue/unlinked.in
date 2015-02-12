package jsonresume

import (
	"fmt"
	"github.com/internavenue/unlinked.in/schemas"
	"strings"
)

type Award struct {
	Awarder string `json:"awarder,omitempty"`
	Date    string `json:"date,omitempty"`
	Summary string `json:"summary,omitempty"`
	Title   string `json:"title,omitempty"`
}

type Location struct {
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Region      string `json:"region,omitempty"`
}

type Profile struct {
	Network  string `json:"network,omitempty"`
	URL      string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
}

type Basics struct {
	Email    string    `json:"email,omitempty"`
	Label    string    `json:"label,omitempty"`
	Location Location  `json:"location"`
	Name     string    `json:"name,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Picture  string    `json:"picture,omitempty"`
	Profiles []Profile `json:"profiles,omitempty"`
	Summary  string    `json:"summary,omitempty"`
	Website  string    `json:"website,omitempty"`
}

type Education struct {
	Area        string   `json:"area,omitempty"`
	Courses     []string `json:"courses,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	Gpa         string   `json:"gpa,omitempty"`
	Institution string   `json:"institution,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	StudyType   string   `json:"studyType,omitempty"`
}

type Interest struct {
	Keywords []string `json:"keywords,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type Language struct {
	Fluency  string `json:"fluency,omitempty"`
	Language string `json:"language,omitempty"`
}

type Publication struct {
	Name        string `json:"name,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Website     string `json:"website,omitempty"`
}

type Reference struct {
	Name      string `json:"name,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type Skill struct {
	Keywords []string `json:"keywords,omitempty"`
	Level    string   `json:"level,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type Volunteer struct {
	EndDate      string   `json:"endDate,omitempty"`
	Highlights   []string `json:"highlights,omitempty"`
	Organization string   `json:"organization,omitempty"`
	Position     string   `json:"position,omitempty"`
	StartDate    string   `json:"startDate,omitempty"`
	Summary      string   `json:"summary,omitempty"`
	Website      string   `json:"website,omitempty"`
}

type Work struct {
	Company    string   `json:"company,omitempty"`
	EndDate    string   `json:"endDate,omitempty"`
	Highlights []string `json:"highlights,omitempty"`
	Position   string   `json:"position,omitempty"`
	StartDate  string   `json:"startDate,omitempty"`
	Summary    string   `json:"summary,omitempty"`
	Website    string   `json:"website,omitempty"`
}

type JSONResumeSchema struct {
	Awards       []Award       `json:"awards,omitempty"`
	Basics       Basics        `json:"basics"`
	Education    []Education   `json:"education,omitempty"`
	Interests    []Interest    `json:"interests,omitempty"`
	Languages    []Language    `json:"languages,omitempty"`
	Publications []Publication `json:"publications,omitempty"`
	References   []Reference   `json:"references,omitempty"`
	Skills       []Skill       `json:"skills,omitempty"`
	Volunteer    []Volunteer   `json:"volunteer,omitempty"`
	Work         []Work        `json:"work,omitempty"`
}

func FromLinkedInSchema(linked *schemas.LinkedInSchema) *JSONResumeSchema {
	r := &JSONResumeSchema{}
	r.Basics.Email = linked.EmailAddress
	r.Basics.Label = linked.Headline
	r.Basics.Name = linked.FormattedName
	if linked.PhoneNumbers.Total > 0 {
		r.Basics.Phone = linked.PhoneNumbers.Values[0].PhoneNumber
	}
	r.Basics.Picture = linked.PictureUrl
	r.Basics.Summary = linked.Summary
	r.Basics.Website = linked.PublicProfileUrl
	r.Basics.Location.CountryCode = linked.Location.Country.Code
	r.Basics.Location.Address = linked.Location.Name
	r.Basics.Profiles = make([]Profile, 0, linked.BoundAccountTypes.Total)
	for _, bat := range linked.BoundAccountTypes.Values {
		if bat.NumAccounts == 0 {
			continue
		}

		for _, ba := range bat.BoundAccounts.Values {
			r.Basics.Profiles = append(r.Basics.Profiles, Profile{
				Network:  ba.AccountType,
				Username: ba.ProviderAccountName,
			})
		}
	}

	r.Education = make([]Education, 0, linked.Educations.Total)
	for _, edu := range linked.Educations.Values {
		r.Education = append(r.Education, Education{
			Area:        edu.FieldOfStudy,
			Institution: edu.SchoolName,
			StudyType:   edu.Degree,
			StartDate:   toStringDate(edu.StartDate),
			EndDate:     toStringDate(edu.EndDate),
		})
	}

	if keys := strings.Split(linked.Interests, "\n"); len(keys) > 0 {
		r.Interests = []Interest{Interest{Keywords: keys}}
	}

	r.Languages = make([]Language, 0, linked.Languages.Total)
	for _, lang := range linked.Languages.Values {
		r.Languages = append(r.Languages, Language{
			Language: lang.Language.Name,
			Fluency:  lang.Proficiency.Name,
		})
	}

	r.Publications = make([]Publication, 0, linked.Publications.Total)
	for _, pub := range linked.Publications.Values {
		r.Publications = append(r.Publications, Publication{
			Name:        pub.Title,
			Publisher:   pub.Publisher.Name,
			Summary:     pub.Summary,
			Website:     pub.URL,
			ReleaseDate: toStringDate(pub.Date),
		})
	}

	r.References = make([]Reference, 0, linked.Recommendations.Total)
	for _, ref := range linked.Recommendations.Values {
		r.References = append(r.References, Reference{
			Name:      ref.Recommender.FirstName + " " + ref.Recommender.LastName,
			Reference: ref.RecommendationText,
		})
	}

	r.Skills = make([]Skill, 0, linked.Skills.Total)
	for _, s := range linked.Skills.Values {
		r.Skills = append(r.Skills, Skill{
			Name: s.Skill.Name,
		})
	}

	r.Volunteer = make([]Volunteer, 0, linked.Volunteer.VolunteerExperiences.Total)
	for _, v := range linked.Volunteer.VolunteerExperiences.Values {
		r.Volunteer = append(r.Volunteer, Volunteer{
			Organization: v.Organization.Name,
			Position:     v.Role,
		})
	}

	r.Work = make([]Work, 0, linked.Positions.Total)
	for _, w := range linked.Positions.Values {
		r.Work = append(r.Work, Work{
			Company:   w.Company.Name,
			Position:  w.Title,
			Summary:   w.Summary,
			StartDate: toStringDate(w.StartDate),
			EndDate:   toStringDate(w.EndDate),
		})
	}

	return r
}

func toStringDate(date schemas.Date) string {
	if date.Year == 0 {
		return ""
	}

	month := date.Month
	if 1 < month || month > 12 {
		month = 1
	}

	day := date.Day
	if 1 < day || day > 31 {
		day = 1
	}

	return fmt.Sprintf("%d-%02d-%02d", date.Year, month, day)
}
