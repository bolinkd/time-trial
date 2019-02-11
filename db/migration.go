package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"fmt"
)

func RunMigrationOnDb(conn *Connection) error {
	driver := goose.DBDriver{Name: "postgres", OpenStr: cnx}
	driver.Import = "github.com/lib/pq"
	driver.Dialect = &goose.PostgresDialect{}

	// migrateConf.Env is ignored with a db connection
	migrateConf := &goose.DBConf{
		MigrationsDir: "db",
		Env:           "time-trial",
		Driver:        driver,
	}

	fmt.Println("HELLO 0")
	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		fmt.Println("HELLO 1")
		return err
	}

	fmt.Println("HELLO 2")
	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, conn.DB.DB)
	if err != nil {
		fmt.Println("HELLO 3")
		return err
	}

	latest, _ = goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)

	return nil
}
