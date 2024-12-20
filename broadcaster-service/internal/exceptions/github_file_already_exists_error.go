package exceptions

type GithubFileAlreadyExistsError struct{}

func (e GithubFileAlreadyExistsError) Error() string {
	return "File already exists"
}
