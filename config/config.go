package config

import (
	"os"
	"strings"
)

var requiredParams = []string{
	"KAPITOL_MONGO_URL",
	"KAPITOL_MONGO_PASSWORD",
	"KAPITOL_MONGO_MEMBERS_COLLECTION",
	"KAPITOL_MONGO_LEGISLATION_COLLECTION",
	"OCTOPUS_LOG_PATH",
}

type Mongo struct {
	Url string
	Password string
	MembersCollection string
	LegislativeCollection string
}

type LogPath string

type Config struct {
	Mongo Mongo
	LogPath LogPath
}

func Configuration() Config {
	missingEnvVars := make([]string, 0, len(requiredParams))
	for _, v := range requiredParams {
		val := os.Getenv(v)
		if val == "" {
			missingEnvVars = append(missingEnvVars, v)
		}
	}

	if len(missingEnvVars) > 0 {
		panic("Octopus environment variables not set properly. Missing:\n" + strings.Join(missingEnvVars, "\n"))
	}

	m := Mongo{
		Url: os.Getenv("KAPITOL_MONGO_URL"),
		Password: os.Getenv("KAPITOL_MONGO_PASSWORD"),
		MembersCollection: os.Getenv("KAPITOL_MONGO_MEMBERS_COLLECTION"),
		LegislativeCollection: os.Getenv("KAPITOL_MONGO_LEGISLATION_COLLECTION"),
	}

	lp := LogPath(os.Getenv("OCTOPUS_LOG_PATH"))

	return Config{Mongo: m, LogPath: lp}
}
