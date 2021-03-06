/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package conf

import (
	"log"
	"os"
	"strconv"
)

var (
	MysqlServerAddr         string = "database"
	MySqlDbMain                    = "wplay"
	MysqlServerPort                = ""
	MysqlLogin                     = "root"
	MysqlPassword                  = "root"
	MysqlMaxUserConnections        = 151
)

func GetEnvVars() {
	MysqlServerAddr = parseEnvVar(os.Getenv("GO_APP_MYSQL_SERVER_ADDR"), "database")
	MySqlDbMain = parseEnvVar(os.Getenv("GO_APP_MYSQL_DB_MAIN"), "wplay")
	MysqlServerPort = ":" + parseEnvVar(os.Getenv("GO_APP_MYSQL_SERVER_PORT"), "3306")
	MysqlLogin = parseEnvVar(os.Getenv("GO_APP_MYSQL_LOGIN"), "root")
	MysqlPassword = parseEnvVar(os.Getenv("GO_APP_MYSQL_PASSWORD"), "WP95123657")
	var err error
	MysqlMaxUserConnections, err = strconv.Atoi(parseEnvVar(os.Getenv("GO_APP_MYSQL_MAX_USER_CONNECTIONS"), "151"))
	if err != nil {
		log.Fatal(err)
	}
}

func parseEnvVar(envVar string, defaultVal string) string {
	if len(envVar) == 0 {
		return defaultVal
	} else {
		return envVar
	}
}
