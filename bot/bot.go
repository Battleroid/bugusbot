package bot

import "fmt"

// Start watching using given configuration
func (config Config) Start() {
	fmt.Println("Started on:", config.Channel)
}
