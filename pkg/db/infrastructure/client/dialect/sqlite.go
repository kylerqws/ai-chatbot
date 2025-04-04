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
		return err
	}

	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("failed to open SQLite dialect: %w", err)
	}

	d.db = bun.NewDB(sqldb, sqlitedialect.New())
	if err := d.init(); err != nil {
		return err
	}
	if err := d.db.Ping(); err != nil {
		return err
	}

	return nil
}

func (d *sqliteDialect) Close() error {
	return d.db.Close()
}

func (d *sqliteDialect) DB() *bun.DB {
	return d.db
}

func (d *sqliteDialect) prepareDSN(dsn string) (string, error) {
	if strings.TrimSpace(dsn) == "" {
		dsn = "var/database.sqlite"
	}
	if strings.HasPrefix(dsn, ":memory:") {
		return dsn, nil
	}

	dir := filepath.Dir(dsn)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return dsn, nil
}

func (d *sqliteDialect) init() error {
	if d.config.IsDebug() {
		d.db.AddQueryHook(bundebug.NewQueryHook())
	}

	_, err := d.db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		return err
	}

	return nil
}
