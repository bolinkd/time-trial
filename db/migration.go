package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
)

func RunMigrationOnDb(conn *Connection) error {
	driver := goose.DBDriver{Name: "postgres", OpenStr: cnx}
	driver.Import = "github.com/lib/pq"
	driver.Dialect = &goose.PostgresDialect{}

	// migrateConf.Env is ignored with a db connection
	migrateConf := &goose.DBConf{
		MigrationsDir: "db",
		Env:           "timetrial",
		Driver:        driver,
	}

	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		return err
	}

	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, conn.DB.DB)
	if err != nil {
		return err
	}

	latest, _ = goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)

	return nil
}
