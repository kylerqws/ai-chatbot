package dialect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"

	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
)

type sqliteDialect struct {
	config ctrcfg.Config
	db     *bun.DB
}

func NewSQLite(cfg ctrcfg.Config) ctrdlt.Dialect {
	return &sqliteDialect{config: cfg}
}

func (d *sqliteDialect) Connect() error {
	dsn, err := d.prepareDSN(d.config.GetDSN())
	if err != nil {
		return fmt.Errorf("failed to prepare DSN: %w", err)
	}

	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	d.db = bun.NewDB(sqldb, sqlitedialect.New())

	err = d.init()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	err = d.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func (d *sqliteDialect) Close() error {
	err := d.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close SQLite database: %w", err)
	}

	return nil
}

func (d *sqliteDialect) DB() *bun.DB {
	if d.db == nil {
		log.Fatalf("database not initialized")
	}

	return d.db
}

func (_ *sqliteDialect) prepareDSN(dsn string) (string, error) {
	if strings.TrimSpace(dsn) == "" {
		dsn = "var/database.sqlite"
	}

	if strings.HasPrefix(dsn, ":memory:") {
		return dsn, nil
	}

	dir := filepath.Dir(dsn)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory %v: %w", dir, err)
	}

	return dsn, nil
}

func (d *sqliteDialect) init() error {
	if d.config.IsDebug() {
		d.db.AddQueryHook(bundebug.NewQueryHook())
	}

	_, err := d.db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return nil
}
