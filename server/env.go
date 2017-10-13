package server

import "os"

type errMissingEnvironmentVariable string

func (e errMissingEnvironmentVariable) Error() string {
	return "pubg-mainmenu-hack(env): missing environment variable " + string(e)
}

type environment struct {
	APIHost         string
	ServePort       string
	SSLRedirectHost string
	RealPUBGURL     string
}

func getEnvironment() (*environment, error) {
	env := &environment{}

	requiredVarsDestinations := map[string]*string{
		"API_HOST":      &env.APIHost,
		"PORT":          &env.ServePort,
		"REAL_PUBG_URL": &env.RealPUBGURL,
	}

	optionalVarsDestinations := map[string]*string{
		"SSL_REDIR": &env.SSLRedirectHost,
	}

	for varName, dest := range requiredVarsDestinations {
		value, err := requireEnvVar(varName)
		if err != nil {
			return nil, err
		}
		*dest = value
	}

	for varName, dest := range optionalVarsDestinations {
		*dest = os.Getenv(varName)
	}

	return env, nil
}

func requireEnvVar(varName string) (string, *errMissingEnvironmentVariable) {
	value := os.Getenv(varName)
	if value == "" {
		err := errMissingEnvironmentVariable(varName)
		return "", &err
	}
	return value, nil
}
