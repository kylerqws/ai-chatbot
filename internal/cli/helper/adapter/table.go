package adapter

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

const (
	// EmptyTableColumn defines a placeholder for empty table columns.
	EmptyTableColumn = "-"
)

// TableAdapter provides the implementation for CLI adapter with table handling.
type TableAdapter struct {
	command *cobra.Command
	table   *table.Writer
	tables  []*table.Writer
}

// NewTableAdapter creates a new instance of TableAdapter.
func NewTableAdapter(cmd *cobra.Command) *TableAdapter {
	return &TableAdapter{command: cmd}
}

// CreateTable creates a new table and sets it as active.
func (a *TableAdapter) CreateTable() uint8 {
	tbl := table.NewWriter()

	tbl.SetOutputMirror(a.command.OutOrStdout())
	tbl.SetStyle(a.DefaultTableStyle())

	a.table = &tbl
	a.tables = append(a.tables, a.table)

	id := len(a.tables) - 1
	return uint8(id)
}

// ResetTables clears all tables.
func (a *TableAdapter) ResetTables() {
	a.table = nil
	a.tables = nil
}

// Table returns the currently active table.
func (a *TableAdapter) Table() *table.Writer {
	return a.table
}

// Tables returns all created tables.
func (a *TableAdapter) Tables() []*table.Writer {
	return a.tables
}

// ExistTables reports whether any tables have been created.
func (a *TableAdapter) ExistTables() bool {
	return len(a.tables) > 0
}

// SwitchTable sets the active table by ID.
func (a *TableAdapter) SwitchTable(id uint8) error {
	if id >= uint8(len(a.tables)) {
		return fmt.Errorf("table with ID %d not found", id)
	}

	a.table = a.tables[id]
	return nil
}

// DefaultTableStyle returns the default table styling.
func (*TableAdapter) DefaultTableStyle() table.Style {
	style := table.StyleBold

	style.Format.HeaderAlign = text.AlignCenter
	style.Color.Header = text.Colors{text.Bold}

	return style
}

// ColumnConfig creates a column configuration.
func (*TableAdapter) ColumnConfig(
	index uint8, align text.Align, width uint8, colors text.Colors,
) table.ColumnConfig {
	return table.ColumnConfig{Number: int(index), Align: align, WidthMin: int(width), Colors: colors}
}

// SetColumnTableConfigs applies column configurations to the active table.
func (a *TableAdapter) SetColumnTableConfigs(configs ...table.ColumnConfig) {
	(*a.table).SetColumnConfigs(configs)
}

// AppendTableHeader appends headers to the active table.
func (a *TableAdapter) AppendTableHeader(headers ...any) {
	(*a.table).AppendHeader(append(table.Row{}, headers...))
}

// AppendTableRow appends a row to the active table.
func (a *TableAdapter) AppendTableRow(rows ...any) {
	(*a.table).AppendRow(append(table.Row{}, rows...))
}

// RenderTable renders the active table.
func (a *TableAdapter) RenderTable() {
	(*a.table).Render()
}
