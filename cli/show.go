package cli

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"wad2linkedin/browser"
	"wad2linkedin/data"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "opens the browser with the search results lead by lead",
	Long: `

	Example:
	wad2linkedin show -file ./contacts.xlsx`,
	Run: showLeads,
}

func showLeads(cmd *cobra.Command, args []string) {

	leads, err := data.ImportFromExcel(wad_export_path)
	if err != nil {
		panic(fmt.Errorf("could not read data from file: %v\n", err))
	}

	skip := 0

	for _, lead := range leads {
		if skip > 0 {
			skip--
			fmt.Printf("skipping %s, %s\n", lead.FirstName, lead.LastName)
			continue
		}

		keywords := make([]string, 0)

		if lead.FirstName != "" {
			keywords = append(keywords, lead.FirstName)
		}

		if lead.LastName != "" {
			keywords = append(keywords, lead.LastName)
		}

		companyName, err := data.ExtractCompanyName(lead)
		if err == nil && companyName != "" {
			keywords = append(keywords, companyName)
		}

		fmt.Println("Lead:")
		lead.PrintInfo()
		fmt.Println("")
		
		err = browser.OpenSearchInLinked(keywords)
		if err != nil {
			fmt.Printf("error opening browser: %v\n", err)
		}

		c, s, exit := Interact()
		if c {
			continue
		}
		if s > 0 {
			skip = s
			continue
		}
		if exit {
			break
		}
	}
}

func Interact() (c bool, skip int, exit bool) {
	fmt.Print("to continue, press enter\n" +
		"to exit: type 'exit' and press enter\n" +
		"to skip: type 'skip n' to skip the next n leads\n")
	input := bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		txt := input.Text()
		if txt == "" {
			return true, 0, false
		}
		if txt == "exit" {
			return false, 0, true
		}
		if strings.HasPrefix(txt, "skip") {
			rawNumber, found := strings.CutPrefix(txt, "skip")
			if !found {
				fmt.Println("invalid input")
				continue
			}
			trimmedNumber := strings.TrimSpace(rawNumber)
			skip, err := strconv.Atoi(trimmedNumber)
			if err != nil {
				fmt.Println("invalid input")
				continue
			}
			return false, skip, false
		}
		fmt.Println("invalid input")
	}

}

func init() {
	rootCmd.AddCommand(showCmd)
}
