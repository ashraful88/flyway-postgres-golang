package flyway

import (
	"fmt"
	"os/exec"
)

func CreateConnectionString(dbHost, dbPort, dbName, dbUser, dbPass, sqlLoc string, retries int64) string {
	if retries == 0 {
		retries = 60
	}
	if sqlLoc == "" {
		sqlLoc = "filesystem:/flyway/sql"
	}
	return fmt.Sprintf("-url=jdbc:postgresql://%s:%s/%s -user=%s -password=%s -connectRetries=%d -locations='%s'",
		dbHost, dbPort, dbName, dbUser, dbPass, retries, sqlLoc)
}

// Info get current migration version info
func Info(cs string) (string, error) {
	cmd := "/flyway/flyway info -outputType=json "
	outInfo, err := exec.Command("/bin/bash", "-c", cmd+cs).CombinedOutput()
	return string(outInfo), err
}

// Clean if need to drop database, all data will be deleted
func Clean(cs string) (string, error) {
	cmd := "/flyway/flyway clean "
	cleanSchema, err := exec.Command("/bin/bash", "-c", cmd+cs).CombinedOutput()
	return string(cleanSchema), err
}

// Migrate runs flyway database schema migration
func Migrate(cs string) (string, error) {
	cmd := "/flyway/flyway migrate "
	out, err := exec.Command("/bin/bash", "-c", cmd+cs).CombinedOutput()
	return string(out), err
}
