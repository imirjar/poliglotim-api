package config

type Config struct {
	*GtwConf
	*StorageConf
}

type GtwConf struct {
	Port string
}

type StorageConf struct {
	MongoConn string
	PsqlConn  string
}

func New() *Config {
	return &Config{
		GtwConf: &GtwConf{
			Port: ":9090",
		},
		StorageConf: &StorageConf{
			MongoConn: "mongodb://imirjar:W6SgTpAcrTjJ41EX7NLYn7oE@db.sleaf.dev/PoliglotimCourses",
			PsqlConn:  "postgres://imirjar:e2h2ey7hgt@91.122.105.45:5432/poliglotim?sslmode=disable",
		},
	}
}
