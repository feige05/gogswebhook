package conf

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"gogsWebHook/api/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	flagsline = `
	-config-file Path to the server configuration file
	-static-dir Path to the static files dir
`
	Config     = &defs.Config{}
	ConfigFile string
	StaticDir  string
	Flags      = flag.NewFlagSet("gogs", flag.ContinueOnError)
)

func init() {
	Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, flagsline)
	}
	Flags.StringVar(&ConfigFile, "config-file", "", "Path to the server configuration file")
	Flags.StringVar(&StaticDir, "static-dir", "", "Path to the static files dir")
}

func Parse(arguments []string) error {
	if err := Flags.Parse(arguments); err != nil {
		switch err {
		case nil:
		case flag.ErrHelp:
			fmt.Println(flagsline)
		default:
		}
		if len(Flags.Args()) != 0 {
			fmt.Errorf("'%s' is not a valid flag", Flags.Arg(0))
		}
		return err
	}
	if ConfigFile != "" {
		// fmt.Printf("conf-file path:%s", path)
		// load extra conf information
		b, err := ioutil.ReadFile(ConfigFile)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(b, &Config); err != nil {
			return err
		}
	} else {
		return errors.New("config-file is not a valid flag")
	}

	return nil
}
