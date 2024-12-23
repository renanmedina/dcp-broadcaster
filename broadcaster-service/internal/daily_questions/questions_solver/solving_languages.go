package questions_solver

type LanguageInfo struct {
	LanguageName  string `gorm:"primaryKey"`
	FileExtension string
	Enabled       bool
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
