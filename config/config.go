package config

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

var requiredParams = []string{
	"KAPITOL_MONGO_URL",
	"KAPITOL_MONGO_DB",
	"KAPITOL_MONGO_PASSWORD",
	"KAPITOL_MONGO_MEMBERS_COLLECTION",
	"KAPITOL_MONGO_LEGISLATION_COLLECTION",
	"OCTOPUS_LOG_PATH",
	"OCTOPUS_LOG_LEVEL",
	"KAPITOL_PRO_PUBLICA_CONGRESS_API_KEY",
}

type Mongo struct {
	Url string
	Db string
	Password string
	MembersCollection string
	LegislativeCollection string
}

type LogInfo struct {
	Path string
	Level int
}

type ApiKeys struct {
	ProPublicaCongress string
}

type Config struct {
	Mongo Mongo
	LogInfo LogInfo
	ApiKeys ApiKeys
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

	level, err := strconv.Atoi(os.Getenv("OCTOPUS_LOG_LEVEL"))
	if err != nil {
		fmt.Println("Error: get octopus log level:", err)
		level = 0
	}
	li := LogInfo{Path: os.Getenv("OCTOPUS_LOG_PATH"), Level: level}

	m := Mongo{
		Url: os.Getenv("KAPITOL_MONGO_URL"),
		Db: os.Getenv("KAPITOL_MONGO_DB"),
		Password: os.Getenv("KAPITOL_MONGO_PASSWORD"),
		MembersCollection: os.Getenv("KAPITOL_MONGO_MEMBERS_COLLECTION"),
		LegislativeCollection: os.Getenv("KAPITOL_MONGO_LEGISLATION_COLLECTION"),
	}

	api := ApiKeys{ProPublicaCongress: os.Getenv("KAPITOL_PRO_PUBLICA_CONGRESS_API_KEY")}

	return Config{Mongo: m, LogInfo: li, ApiKeys: api}
}
