package bot

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Address   string   // Server address
	Port      uint16   // Server port
	Username  string   // Username to assume
	MoveTo    string   // Channel to move to upon joining
	Channel   string   // Channel to watch
	Threshold uint8    // Threshold of players before notifying
	Interval  int      // How often to check
	Sayings   []string // List of phrases to say
}

// Load yaml configuration
func LoadYAML(filename string) *Config {
	var err error
	c := Config{}

	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read YAML from filename %s", filename)
		os.Exit(1)
	}

	err = yaml.Unmarshal(fb, &c)
	if err != nil {
		log.Fatal("Could not unmarshal YAML, is the format correct?")
		os.Exit(1)
	}

	return &c
}
