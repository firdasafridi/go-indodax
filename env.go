package indodax

import (
	"os"
)

//
// environment contains default and dynamics values that gathered from external
// resources, for example system environment variables.
//
type environment struct {
	BaseHostPublic  string
	BaseHostPrivate string
	apiKey          string
	apiSecret       string
}

func newEnvironment() (env *environment) {
	env = &environment{
		BaseHostPublic:  UrlPublic,
		BaseHostPrivate: UrlPrivate,
		apiKey:          os.Getenv("INDODAX_KEY"),
		apiSecret:       os.Getenv("INDODAX_SECRET"),
	}

	pubHost := os.Getenv("INDODAX_PUB_HOST")
	if pubHost != "" {
		env.BaseHostPublic = pubHost
	}

	privHost := os.Getenv("INDODAX_PRIV_HOST")
	if privHost != "" {
		env.BaseHostPrivate = privHost
	}

	return env
}
