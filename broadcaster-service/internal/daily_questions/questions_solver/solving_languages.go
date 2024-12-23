package questions_solver

import "github.com/renanmedina/dcp-broadcaster/utils"

type LanguageInfo struct {
	LanguageName  string
	FileExtension string
	Enabled       bool
}

type SolvingLanguagesRepository struct {
	db     *utils.DatabaseAdapdater
	logger *utils.ApplicationLogger
}

const (
	QUESTIONS_SOLUTION_LANGUAGES_TABLE_NAME = "daily_questions_solutions"
	QUESTIONS_SOLUTION_LANGUAGES_FIELDS     = "language_name, file_extension, enabled"
)

func (r *SolvingLanguagesRepository) GetAll() ([]LanguageInfo, error) {
	rows, err := r.db.Select(QUESTIONS_SOLUTION_LANGUAGES_FIELDS, QUESTIONS_SOLUTION_LANGUAGES_TABLE_NAME, nil)

	if err != nil {
		return make([]LanguageInfo, 0), err
	}

	defer rows.Close()

	var languages []LanguageInfo

	for rows.Next() {
		var languageInfo LanguageInfo
		err = rows.Scan(
			&languageInfo.LanguageName,
			&languageInfo.FileExtension,
			&languageInfo.Enabled,
		)

		if err != nil {
			panic(err.Error())
		}

		languages = append(languages, languageInfo)
	}

	return languages, nil
}

func (r *SolvingLanguagesRepository) GetAllEnabled() ([]LanguageInfo, error) {
	rows, err := r.db.Select(
		QUESTIONS_SOLUTION_LANGUAGES_FIELDS,
		QUESTIONS_SOLUTION_LANGUAGES_TABLE_NAME,
		map[string]interface{}{
			"enabled": 1,
		},
	)

	if err != nil {
		return make([]LanguageInfo, 0), err
	}

	defer rows.Close()

	var languages []LanguageInfo

	for rows.Next() {
		var languageInfo LanguageInfo
		err = rows.Scan(
			&languageInfo.LanguageName,
			&languageInfo.FileExtension,
			&languageInfo.Enabled,
		)

		if err != nil {
			panic(err.Error())
		}

		languages = append(languages, languageInfo)
	}

	return languages, nil
}

func NewSolveLanguagesRepository() SolvingLanguagesRepository {
	return SolvingLanguagesRepository{
		utils.GetDatabase(),
		utils.GetApplicationLogger(),
	}
}

var SolvingLanguages = map[string]LanguageInfo{
	"golang":     {"golang", "go", true},
	"ruby":       {"ruby", "rb", true},
	"php":        {"php", "php", true},
	"python":     {"python", "py", true},
	"elixir":     {"elixir", "ex", true},
	"javascript": {"javascript", "js", true},
	"c":          {"c", "c", true},
	// "csharp":     {"csharp", "cs", true},
	// "java":       {"java", "java", true},
	// "c++":        {"c++", "cpp", true},
	// "dart":       {"dart", "dart", true},
}
