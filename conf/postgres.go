package conf

type postgresStruct struct {
	User     string
	Password string
	DBName   string
	Port     string
}

var Postgres postgresStruct

func init() {
	Postgres = postgresStruct{
		User:     "postgres",
		Password: "1538",
		DBName:   "db_proj",
		Port:     "5432",
	}
}
