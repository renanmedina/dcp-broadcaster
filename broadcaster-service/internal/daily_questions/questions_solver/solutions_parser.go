package questions_solver

import (
	"errors"
	"fmt"
	"regexp"
)

func extractSolutionCodeFrom(responseContent string, programmingLanguge string) (string, error) {
	regexString := fmt.Sprintf("(```%s.+```)", programmingLanguge)
	matcher := regexp.MustCompile(regexString)
	matches := matcher.FindStringSubmatch(responseContent)

	if len(matches) > 0 {
		return matches[len(matches)-1], nil
	}

	return "", errors.New("No code found to be extracted")
}
