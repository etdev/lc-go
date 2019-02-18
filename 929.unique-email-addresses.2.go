import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	seen := map[string]bool{}
	count := 0

	for _, email := range emails {
		split := strings.Split(email, "@")
		username, domain := split[0], split[1]
		username = strings.Replace(username, ".", "", -1)
		username = strings.Split(username, "+")[0]
		cleaned := fmt.Sprintf("%s%s", username, domain)

		if !seen[cleaned] {
			count++
			seen[cleaned] = true
		}
	}

	return count
}
