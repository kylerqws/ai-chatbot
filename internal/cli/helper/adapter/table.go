package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const EmptyTableColumn = "-"

type TableAdapter struct {
	command *cobra.Command
	table   *table.Writer
	tables  []*table.Writer
}

func NewTableAdapter(cmd *cobra.Command) *TableAdapter {
	return &TableAdapter{command: cmd}
}

func (h *TableAdapter) CreateTable() uint8 {
	tbl := table.NewWriter()

	tbl.SetOutputMirror(h.command.OutOrStdout())
	tbl.SetStyle(h.DefaultTableStyle())

	h.table = &tbl
	h.tables = append(h.tables, h.table)

	id := len(h.tables) - 1
	return uint8(id)
}

func (h *TableAdapter) ResetTables() {
	h.table = nil
	h.tables = nil
}

func (h *TableAdapter) Table() *table.Writer {
	return h.table
}

func (h *TableAdapter) Tables() []*table.Writer {
	return h.tables
}

func (h *TableAdapter) ExistTables() bool {
	return len(h.tables) > 0
}

func (h *TableAdapter) SwitchTable(id uint8) error {
	if id >= uint8(len(h.tables)) {
		return fmt.Errorf("table with ID %d not found", id)
	}

	h.table = h.tables[id]
	return nil
}

func (*TableAdapter) DefaultTableStyle() table.Style {
	style := table.StyleBold

	style.Format.HeaderAlign = text.AlignCenter
	style.Color.Header = text.Colors{text.Bold}

	return style
}

func (*TableAdapter) ColumnConfig(
	index uint8, align text.Align, width uint8, colors text.Colors,
) table.ColumnConfig {
	return table.ColumnConfig{Number: int(index), Align: align, WidthMin: int(width), Colors: colors}
}

func (h *TableAdapter) SetColumnTableConfigs(configs ...table.ColumnConfig) {
	(*h.table).SetColumnConfigs(configs)
}

func (h *TableAdapter) AppendTableHeader(headers ...any) {
	(*h.table).AppendHeader(append(table.Row{}, headers...))
}

func (h *TableAdapter) AppendTableRow(rows ...any) {
	(*h.table).AppendRow(append(table.Row{}, rows...))
}

func (h *TableAdapter) RenderTable() {
	(*h.table).Render()
}
