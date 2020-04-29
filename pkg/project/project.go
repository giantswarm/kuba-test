package project

var (
	description        = "Kuba's test package"
	gitSHA             = "n/a"
	name        string = "kuba-test"
	source             = ""
	version            = "0.1.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
