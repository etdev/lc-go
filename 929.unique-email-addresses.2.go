import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	seen := map[string]bool{}

	for _, email := range emails {
		split := strings.Split(email, "@")
		username := strings.Replace(split[0], ".", "", -1)
		username = strings.Split(username, "+")[0]
		domain := split[1]
		seen[strings.Join([]string{username, domain}, "")] = true
	}

	return len(seen)
}
