package scripts

type Script struct {
	name     string
	commands []string
}

func AddScript(name string, commands []string) *Script {
	return &Script{name, commands}
}
