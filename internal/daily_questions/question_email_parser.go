package daily_questions

import (
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/emersion/go-imap"
)

type QuestionEmailParser struct{}

type QuestionEmailMetadata struct {
	MessageId    string
	Date         time.Time
	Title        string
	OriginalBody string
	BodyHtml     string
	BodyText     string
	Difficulty   string
	CompanyName  string
}

func (m *QuestionEmailMetadata) Valid() bool {
	if m.Difficulty != "" && m.CompanyName != "" {
		return true
	}

	return false
}

func parseQuestionEmailMessage(msg *imap.Message) QuestionEmailMetadata {
	bodyData := readMessageBody(msg.Body)
	questionHtml := parseMessageToHtml(bodyData)
	questionText := parseMessageToText(bodyData)

	return QuestionEmailMetadata{
		MessageId:    extractMessageId(msg.Envelope.MessageId),
		Date:         msg.Envelope.Date,
		Title:        msg.Envelope.Subject,
		Difficulty:   extractDifficulty(msg.Envelope.Subject),
		OriginalBody: bodyData,
		BodyHtml:     questionHtml,
		BodyText:     questionText,
		CompanyName:  extractCompanyName(questionText),
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

func parseMessageToText(bodyContent string) string {
	return extractByRegex(bodyContent, `Good morning!(?s:.+)--------?`)
}

func parseMessageToHtml(bodyContent string) string {
	return extractByRegex(bodyContent, `\<!DOCTYPE html (?s:.+)`)
}

func readMessageBody(messageBody map[*imap.BodySectionName]imap.Literal) string {
	for _, v := range messageBody {
		body, err := io.ReadAll(v)
		if err != nil {
			log.Fatal(err.Error())
		}

		return string(body)
	}

	return ""
}

func extractByRegex(text string, expression string) string {
	rxp := regexp.MustCompile(expression)
	matches := rxp.FindStringSubmatch(text)

	// fmt.Println("Evaluating expression:")
	// fmt.Println(expression)
	// fmt.Println("Results:")
	// fmt.Println(matches)

	if matches != nil {
		return matches[len(matches)-1]
	}

	return ""
}
