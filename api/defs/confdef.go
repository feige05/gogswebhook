package defs

type Config struct {
	JavaPath       string `yaml:"JavaPath"`
	JenkinsCliPath string `yaml:"JenkinsCliPath"`
	JenkinsURL     string `yaml:"JenkinsURL"`
	JenkinsAuth    string `yaml:"JenkinsAuth"`
	Port           string `yaml:"Port"`
	Store          []*Job `yaml:"Store"`
}

type Job struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}
