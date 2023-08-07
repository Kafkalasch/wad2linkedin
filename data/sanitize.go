package data

import (
	"fmt"
	"regexp"
	"strings"
)

var commonMailProviders = []string{
	"gmail",
	"outlook",
	"yahoo",
	"icloud",
	"web",
	"gmx",
	"t-online",
	"proton",
	"zoho",
	"msn",
}

func ExtractCompanyName(lead Lead) (string, error) {
	companyName := lead.Company
	if companyName != "" {
		return companyName, nil
	}

	mail := lead.Email
	if mail != "" {
		guessedName, err := guessCompanyNameFromEmail(mail)
		if err != nil {
			return "", err
		}
		return guessedName, nil
	}
	return "", fmt.Errorf("no company found")
}

func guessCompanyNameFromEmail(email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	regex := regexp.MustCompile(`@([^.]+)`)

	match := regex.FindStringSubmatch(email)
	if len(match) > 1 {
		companyName := match[1]

		for _, provider := range commonMailProviders {
			if strings.EqualFold(companyName, provider) {
				return "", fmt.Errorf("provided a private mail adress")
			}
		}

		return companyName, nil
	}
	return "", fmt.Errorf("no match found")
}
