package data

import "fmt"

type Lead struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	JobTitle      string `json:"jobTitle"`
	Company       string `json:"company"`
	Email         string `json:"email"`
	MobilePhone   string `json:"mobilePhone"`
	LandlinePhone string `json:"landlinePhone"`
	Website       string `json:"website"`
	Place         string `json:"place"`
	Street        string `json:"street"`
	ZipCode       string `json:"zipCode"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Biography     string `json:"biography"`
	Twitter       string `json:"twitter"`
	Linkedin      string `json:"linkedin"`
	ConnectedVia  string `json:"connectedVia"`
	Scoring       int    `json:"scoring"`
	Tags          string `json:"tags"`
	Note          string `json:"note"`
}

func (l Lead) PrintInfo() {
	fmt.Printf("Name:\t\t\t%s %s\n", l.FirstName, l.LastName)
	fmt.Printf("Scoring:\t\t%d\n", l.Scoring)
	if l.JobTitle != "" {
		fmt.Printf("JobTitle:\t\t%s\n", l.JobTitle)
	}
	if l.Company != "" {
		fmt.Printf("Company:\t\t%s\n", l.Company)
	}
	if l.Biography != "" {
		fmt.Printf("Biography:\t\t%s\n", l.Biography)
	}
	if l.Tags != "" {
		fmt.Printf("Tags:\t\t\t%s\n", l.Tags)
	}
	if l.Note != "" {
		fmt.Printf("Note:\t\t\t%s\n", l.Note)
	}
	if l.Email != "" {
		fmt.Printf("Email:\t\t\t%s\n", l.Email)
	}
	if l.Linkedin != "" {
		fmt.Printf("Linkedin:\t\t%s\n", l.Linkedin)
	}
	if l.Twitter != "" {
		fmt.Printf("Twitter:\t\t%s\n", l.Twitter)
	}
	if l.MobilePhone != "" {
		fmt.Printf("MobilePhone:\t\t%s\n", l.MobilePhone)
	}
	if l.LandlinePhone != "" {
		fmt.Printf("LandlinePhone:\t\t%s\n", l.LandlinePhone)
	}
	if l.Website != "" {
		fmt.Printf("Website:\t\t%s\n", l.Website)
	}
	if l.Country != "" {
		fmt.Printf("Country:\t\t%s\n", l.Country)
	}
	if l.State != "" {
		fmt.Printf("State:\t\t\t%s\n", l.State)
	}
	if l.City != "" {
		fmt.Printf("City:\t\t\t%s\n", l.City)
	}
}
