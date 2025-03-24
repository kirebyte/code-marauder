package settings

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	CodeListingDirs         []string `envconfig:"CODE_LISTING_DIRS"`
	CodeListingFile         string   `envconfig:"CODE_LISTING_FILE"`
	CodeListingAllowedFiles []string `envconfig:"CODE_LISTING_ALLOWED_FILES"`
}

// Runtime contains all the needed runtime settings
var Runtime Settings

// init is called before the main function.
func init() {
	Runtime = Settings{}

	// check if the environment variables are set and override the default values.
	err := envconfig.Process("", &Runtime)
	if err != nil {
		log.Fatalf("error loading environment variables: %s", err.Error())
	}

	log.Println("config loaded")
}
