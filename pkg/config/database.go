package config

type DatabaseEnv struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Dbname string
}

var DatabaseConfig *DatabaseEnv = &DatabaseEnv{
	Host:   "localhost",
	Port:   "5432",
	User:   "user",
	Pass:   "pass",
	Dbname: "gofortickets",
}
