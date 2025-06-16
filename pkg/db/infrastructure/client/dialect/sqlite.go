package dialect

import (
	"database/sql"
	"fmt"
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

// sqliteDialect implements the Dialect interface using SQLite.
type sqliteDialect struct {
	config ctrcfg.Config
	db     *bun.DB
}

// NewSQLite creates a new SQLite dialect.
func NewSQLite(cfg ctrcfg.Config) ctrdlt.Dialect {
	return &sqliteDialect{config: cfg}
}

// Connect opens the connection via the dialect.
func (d *sqliteDialect) Connect() error {
	dsn, err := d.prepareDSN(d.config.GetDSN())
	if err != nil {
		return fmt.Errorf("prepare DSN: %w", err)
	}

	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("open SQLite connection: %w", err)
	}

	d.db = bun.NewDB(sqldb, sqlitedialect.New())

	if err := d.init(); err != nil {
		return fmt.Errorf("init database: %w", err)
	}
	if err := d.db.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	return nil
}

// Close closes the database connection.
func (d *sqliteDialect) Close() error {
	if err := d.db.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}
	return nil
}

// DB returns the low-level database instance.
func (d *sqliteDialect) DB() *bun.DB {
	if d.db == nil {
		panic("database not initialized")
	}
	return d.db
}

// prepareDSN normalizes and ensures the DSN path is valid.
func (*sqliteDialect) prepareDSN(dsn string) (string, error) {
	if dsn = strings.TrimSpace(dsn); dsn == "" {
		dsn = "var/database.sqlite"
	}

	if strings.HasPrefix(dsn, ":memory:") {
		return dsn, nil
	}

	dir := filepath.Dir(dsn)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("create directory '%s': %w", dir, err)
	}

	return dsn, nil
}

// init applies SQLite-specific settings and debug mode if enabled.
func (d *sqliteDialect) init() error {
	if d.config.IsDebug() {
		d.db.AddQueryHook(bundebug.NewQueryHook())
	}

	if _, err := d.db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return fmt.Errorf("enable foreign keys: %w", err)
	}

	return nil
}
