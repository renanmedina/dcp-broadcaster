package questions_solver

type LanguageInfo struct {
	LanguageName  string
	FileExtension string
	Enabled       bool
}

var SolvingLanguages = map[string]LanguageInfo{
	"golang":     {"golang", "go", true},
	"ruby":       {"ruby", "rb", true},
	"php":        {"php", "php", true},
	"python":     {"python", "py", true},
	"elixir":     {"elixir", "ex", true},
	"csharp":     {"csharp", "cs", true},
	"java":       {"java", "java", true},
	"javascript": {"javascript", "js", true},
	"c++":        {"c++", "cpp", true},
	"c":          {"c", "c", true},
	"dart":       {"dart", "dart", true},
}
