import (
	"fmt"
)

func numUniqueEmails(emails []string) int {
	seen := map[string]bool{}
	count := 0

	for _, email := range emails {
		cleaned := ""
		recording := true
		host := false

		for _, chr := range email {
			if chr == '.' {
				if !host {
					continue
				}
			}
			if chr == '+' {
				recording = false
			}
			if chr == '@' {
				recording = true
				host = true
				continue
			}

			if recording {
				cleaned = fmt.Sprintf("%s%s", cleaned, string(chr))
			}
		}

		if !seen[cleaned] {
			count++
		}
		seen[cleaned] = true
	}

	return count
}
