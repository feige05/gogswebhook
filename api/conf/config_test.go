package conf

import (
	"gogsWebHook/api/defs"
	"testing"
)

func TestConfigParsingFlags(t *testing.T) {
	args := []string{
		"-conf-file=./config.yml",
	}

	err := Parse(args)
	if err != nil {
		t.Fatal(err)
	}

	validateConfigParsingFlags(t, Config)
}

func validateConfigParsingFlags(t *testing.T, cfg *defs.Config) {
	if cfg.JenkinsCliPath != "/root/.jenkins/war/WEB-INF/jenkins-cli.jar" {
		t.Errorf("configFile = %v, want %v", cfg.JenkinsCliPath, "/root/.jenkins/war/WEB-INF/jenkins-cli.jar")
	}
}
