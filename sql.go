package testish

import (
	"fmt"
	"github.com/omidfth/testish/internal/router"
	"github.com/omidfth/testish/internal/types/serviceNames"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"time"
)

type sql struct {
}

const (
	portRegex            = `({PORT})`
	initDBRegex          = `(INIT_DB_PATH)`
	dumpRegex            = `(DUMP_PATH)`
	restoreRegex         = `(RESTORE_PATH)`
	pgTemplate           = "/docker-conf/postgres_template.yml"
	pgRestoreTemplate    = "/docker-conf/data/postgres_restore_template.sh"
	mysqlTemplate        = "/docker-conf/mysql_template.yml"
	mysqlRestoreTemplate = "/docker-conf/data/mysql_restore_template.sh"
	pgInitDBTemplate     = "/docker-conf/data/db_init_pg_template.sql"
	mysqlInitDBTemplate  = "/docker-conf/data/db_init_mysql_template.sql"
	dockerComposeOutput  = "/docker-compose/"
	dataOutput           = "/docker-compose/data/"
)

func newSql(r router.Router) {
	r.On(serviceNames.POSTGRESQL, createPostgresql)
	r.On(serviceNames.MYSQL, createMysql)
}

func generateFiles(option *Option, initDBTemplate string, composeTemplate string, restoreTemplate string, dockerComposeName string, initDbName string, dumpName string, restoreName string, defaultPort int) string {
	dir, _ := dirname()
	tempPath := dir + composeTemplate
	initTempPath := dir + initDBTemplate
	restorePath := dir + restoreTemplate
	dockerComposeOutputPath := dir + dockerComposeOutput + dockerComposeName
	initDBOutputPath := dir + dataOutput + initDbName
	dumpOutput := dir + dataOutput + dumpName
	restoreOutput := dir + dataOutput + restoreName

	createInitDB(initTempPath, initDBOutputPath)
	createDump(option.DumpPath, dumpOutput)
	createRestore(restorePath, restoreOutput)
	createDockerComposeFile(tempPath, option.ExposePort, initDBOutputPath, dumpOutput, restoreOutput, dockerComposeOutputPath, defaultPort)
	stopDockerCompose(dockerComposeOutputPath)
	runDockerCompose(dockerComposeOutputPath)

	return dockerComposeOutputPath
}

func getGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
}

func createMysql(i interface{}) (interface{}, string) {
	option := castInterfaceToOption(i)
	out := generateFiles(
		option,
		mysqlInitDBTemplate,
		mysqlTemplate,
		mysqlRestoreTemplate,
		"mysql.yml",
		"msql.sql",
		"mysql_dump.sql",
		"mysql_restore.sh", 3307)

	time.Sleep(time.Duration(10000) * time.Millisecond)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"db_user",
		"db_user",
		"127.0.0.1",
		option.ExposePort,
		"db_test",
	)

	log.Println("DSN:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), getGormConfig())

	if err == nil {
		log.Println("Database Connected!")
		return db, out
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(1000*i) * time.Millisecond)
		fmt.Println("connect to db ... try:", i+2)
		db, err = gorm.Open(mysql.Open(dsn), getGormConfig())
		if err == nil {
			log.Println("Database Connected!")
			return db, out
		}
	}

	os.Exit(14)
	return nil, out

}

func createPostgresql(i interface{}) (interface{}, string) {
	option := castInterfaceToOption(i)
	out := generateFiles(
		option,
		pgInitDBTemplate,
		pgTemplate,
		pgRestoreTemplate,
		"pg.yml",
		"pg.sql",
		"pg_dump.sql",
		"pg_restore.sh", 5432)

	time.Sleep(time.Duration(10000) * time.Millisecond)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"host.docker.internal",
		option.ExposePort,
		"db_user",
		"db_user",
		"db_test")
	db, err := gorm.Open(postgres.Open(dsn), getGormConfig())

	if err == nil {
		log.Println("Database Connected!")
		return db, out
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(1000*i) * time.Millisecond)
		fmt.Println("connect to db ... try:", i+2)
		db, err = gorm.Open(postgres.Open(dsn), getGormConfig())
		if err == nil {
			log.Println("Database Connected!")
			return db, out
		}
	}

	os.Exit(14)
	return nil, out
}

func createInitDB(initDBPath string, output string) {
	data, _ := os.ReadFile(initDBPath)
	err := os.WriteFile(output, data, os.ModePerm)
	if err != nil {
		log.Println("ERR", err.Error())
	}
}

func createDump(dumpPath string, output string) {
	data, _ := os.ReadFile(dumpPath)
	err := os.WriteFile(output, data, os.ModePerm)
	if err != nil {
		log.Println("ERR", err.Error())
	}
}

func createRestore(restorePath string, output string) {
	data, _ := os.ReadFile(restorePath)
	err := os.WriteFile(output, data, os.ModePerm)
	if err != nil {
		log.Println("ERR", err.Error())
	}
}

func createDockerComposeFile(path string, port int, initDBPath string, dumpPath string, restorePath string, output string, defaultPort int) {
	data, _ := os.ReadFile(path)
	portPattern := regexp.MustCompile(portRegex)
	p := fmt.Sprintf("\"%d:%d\"", port, defaultPort)
	str := portPattern.ReplaceAllString(string(data), p)

	dbPattern := regexp.MustCompile(initDBRegex)
	str = dbPattern.ReplaceAllString(str, initDBPath)

	dumpPattern := regexp.MustCompile(dumpRegex)
	str = dumpPattern.ReplaceAllString(str, dumpPath)

	restorePattern := regexp.MustCompile(restoreRegex)
	str = restorePattern.ReplaceAllString(str, restorePath)

	os.WriteFile(output, []byte(str), 0666)
}
