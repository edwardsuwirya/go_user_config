package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(GetDataSourceName(newConfigHardCoded()))
	//fmt.Println(GetDataSourceName(newConfigWithPositionalArgument()))
	//fmt.Println(GetDataSourceName(newConfigWithFlag()))
	//fmt.Println(GetDataSourceName(newConfigWithFlagShortLong()))
	//fmt.Println(GetDataSourceName(newConfigWithEnv()))
	//for{}
}

func newConfigHardCoded() Config{
	// Kalau ganti environment server, ganti code, commit-push
	return Config{
		dbHost:     "localhost",
		dbPort:     "5432",
		dbName:     "test",
		dbUser:     "root",
		dbPassword: "P@ssw0rd",
	}
}

/*
go run app.go config.go "10.0.1.1" 5432 postgres 12345 enigma
*/
func newConfigWithPositionalArgument() Config {
	args := os.Args[1:]
	fmt.Println(os.Args)
	return Config{
		dbHost:     args[0],
		dbPort:     args[1],
		dbName:     args[2],
		dbUser:     args[3],
		dbPassword: args[4],
	}
}

func newConfigWithFlag() Config {
	dbHost := flag.String("host", "localhost", "Database Host Server")
	dbPort := flag.String("port", "5432", "Database Server Port")
	dbUser := flag.String("user", "postgres", "Database User Name")
	dbPassword := flag.String("password", "", "Database User Password")
	dbName := flag.String("db", "test", "Database Name")

	flag.Parse()
	return Config{
		dbHost:     *dbHost,
		dbPort:     *dbPort,
		dbName:     *dbName,
		dbUser:     *dbUser,
		dbPassword: *dbPassword,
	}
}

func newConfigWithFlagShortLong() Config {
	var dbHost string
	var dbPort string
	var dbName string
	var dbUser string
	var dbPassword string

	flag.StringVar(&dbHost, "host", "localhost", "Database Host Server")
	flag.StringVar(&dbPort, "port", "5432", "Database Server Port")
	flag.StringVar(&dbUser, "user", "postgres", "Database User Name")
	flag.StringVar(&dbPassword, "password", "", "Database User Password")
	flag.StringVar(&dbName, "db", "test", "Database Name")

	flag.StringVar(&dbHost, "h", "localhost", "Database Host Server")
	flag.StringVar(&dbPort, "p", "5432", "Database Server Port")
	flag.StringVar(&dbUser, "u", "postgres", "Database User Name")
	flag.StringVar(&dbPassword, "P", "", "Database User Password")
	flag.StringVar(&dbName, "d", "test", "Database Name")

	flag.Usage = func() {
		fmt.Print(`
Usage of Config With Flag:
  -P, --password string
        Database User Password
  -d, --db string
        Database Name (default "test")
  -h, --host string
        Database IP (default "localhost")
  -p, --port string
        Database Port (default "5432")
  -u, --user string
        Database User Name (default "postgres")
`)
	}
	flag.Parse()
	return Config{
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword,
	}
}

func newConfigWithEnv() Config {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	return Config{
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword,
	}
}
