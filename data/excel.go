package data

import (
	"fmt"
	"github.com/thedatashed/xlsxreader"
	"strconv"
)

// ImportFromExcel reads an excel file and returns a slice of leads.
func ImportFromExcel(filePath string) ([]Lead, error) {
	// Create an instance of the reader by opening a target file
	xl, err := xlsxreader.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	// Ensure the file reader is closed once utilised
	defer xl.Close()

	sheets := xl.XlsxFile.Sheets
	if len(sheets) != 1 {
		return nil, fmt.Errorf("expected 1 sheet, got %d", len(sheets))
	}

	sheetName := sheets[0]

	var headers map[string]string = nil

	allLeads := make([]Lead, 0)
	for row := range xl.ReadRows(sheetName) {
		if headers == nil {
			// first line -> extract headers
			headers = extractHeaders(row)
			continue
		}
		rowData := extractData(row, headers)
		lead := convertData(rowData)
		allLeads = append(allLeads, lead)
	}

	return allLeads, nil
}

func extractHeaders(row xlsxreader.Row) map[string]string {
	cells := row.Cells
	headers := make(map[string]string, len(cells))
	for _, cell := range cells {
		headers[cell.Column] = cell.Value
	}
	return headers
}

func extractData(row xlsxreader.Row, headers map[string]string) map[string]string {
	cells := row.Cells
	data := make(map[string]string)
	for _, cell := range cells {
		name := headers[cell.Column]
		value := cell.Value
		data[name] = value
	}
	return data
}

func convertData(rowData map[string]string) Lead {
	const firstNameKey = "First name"
	const lastNameKey = "Last name"
	const jobTitleKey = "Job title"
	const companyKey = "Company"
	const emailKey = "Email"
	const mobilePhoneKey = "Mobile phone"
	const landlinePhoneKey = "Landline phone"
	const websiteKey = "Website"
	const placeKey = "Place"
	const streetKey = "Street"
	const zipCodeKey = "Zip code"
	const cityKey = "City"
	const stateKey = "State"
	const countryKey = "Country"
	const biographyKey = "Biography"
	const twitterKey = "Twitter"
	const linkedinKey = "Linkedin"
	const connectedVia = "Connected via"
	const scoringKey = "Scoring"
	const tagsKey = "Tags"
	const noteKey = "Note"

	score, err := strconv.Atoi(rowData[scoringKey])
	if err != nil {
		score = 0
	}
	return Lead{
		FirstName:     rowData[firstNameKey],
		LastName:      rowData[lastNameKey],
		JobTitle:      rowData[jobTitleKey],
		Company:       rowData[companyKey],
		Email:         rowData[emailKey],
		MobilePhone:   rowData[mobilePhoneKey],
		LandlinePhone: rowData[landlinePhoneKey],
		Website:       rowData[websiteKey],
		Place:         rowData[placeKey],
		Street:        rowData[streetKey],
		ZipCode:       rowData[zipCodeKey],
		City:          rowData[cityKey],
		State:         rowData[stateKey],
		Country:       rowData[countryKey],
		Biography:     rowData[biographyKey],
		Twitter:       rowData[twitterKey],
		Linkedin:      rowData[linkedinKey],
		ConnectedVia:  rowData[connectedVia],
		Scoring:       score,
		Tags:          rowData[tagsKey],
		Note:          rowData[noteKey],
	}
}
