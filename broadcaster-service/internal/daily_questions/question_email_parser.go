package daily_questions

import (
	"regexp"
	"strings"
	"time"

	imap "github.com/BrianLeishman/go-imap"
)

type QuestionEmailParser struct{}

type QuestionEmailMetadata struct {
	MessageId   string
	Date        time.Time
	Title       string
	BodyHtml    string
	BodyText    string
	Difficulty  string
	CompanyName string
}

func (m *QuestionEmailMetadata) Valid() bool {
	if m.Difficulty != "" && m.CompanyName != "" {
		return true
	}

	return false
}

func parseQuestionEmailMessage(msg *imap.Email) QuestionEmailMetadata {
	questionHtml := msg.HTML
	questionText := extractQuestionText(msg.Text)
	companyName := extractCompanyName(questionText)
	difficulty := extractDifficulty(msg.Subject)

	return QuestionEmailMetadata{
		MessageId:   extractMessageId(msg.MessageID),
		Date:        msg.Received,
		Title:       msg.Subject,
		Difficulty:  difficulty,
		BodyHtml:    questionHtml,
		BodyText:    questionText,
		CompanyName: companyName,
	}
}

func extractMessageId(emailMessageIdHeader string) string {
	return extractByRegex(emailMessageIdHeader, `\<(.+)\>`)
}

func extractCompanyName(questionText string) string {
	return extractByRegex(questionText, `asked by (.+)\.`)
}

func extractDifficulty(title string) string {
	return strings.ToLower(
		extractByRegex(title, `Problem #[0-9]+ \[(.+)\]`),
	)
}

func extractQuestionText(bodyContent string) string {
	extracted := extractByRegex(bodyContent, `Good morning!(?s:.+)--------?`)
	splits := strings.Split(extracted, "--------")

	if len(splits) > 0 {
		return strings.TrimSpace(splits[0])
	}

	return ""
}

func extractByRegex(text string, expression string) string {
	rxp := regexp.MustCompile(expression)
	matches := rxp.FindStringSubmatch(text)

	if matches != nil {
		return matches[len(matches)-1]
	}

	return ""
}
