package flyway

import (
	"fmt"
	"os/exec"
)

// CreateConnectionString creates connection string for flyway
func CreateConnectionString(dbHost, dbPort, dbName, dbUser, dbPass, schemas, sqlLoc, retries string) string {
	if retries == "" {
		retries = "60"
	}
	if schemas == "" {
		schemas = "public"
	}
	if sqlLoc == "" {
		sqlLoc = "filesystem:/flyway/sql"
	}
	return fmt.Sprintf(`-url=jdbc:postgresql://%s:%s/%s -user=%s -password=%s -schemas="%s" -connectRetries=%s -locations='%s'`,
		dbHost, dbPort, dbName, dbUser, dbPass, schemas, retries, sqlLoc)
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
