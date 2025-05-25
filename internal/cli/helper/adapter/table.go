package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const EmptyTableColumn = "-"

type TableAdapterHelper struct {
	command *cobra.Command
	table   *table.Writer
	tables  []*table.Writer
}

func NewTableAdapterHelper(cmd *cobra.Command) *TableAdapterHelper {
	return &TableAdapterHelper{command: cmd}
}

func (h *TableAdapterHelper) CreateTable() uint8 {
	tbl := table.NewWriter()

	tbl.SetOutputMirror(h.command.OutOrStdout())
	tbl.SetStyle(h.DefaultTableStyle())

	h.table = &tbl
	h.tables = append(h.tables, h.table)

	id := len(h.tables) - 1
	return uint8(id)
}

func (h *TableAdapterHelper) ResetTables() {
	h.table = nil
	h.tables = nil
}

func (h *TableAdapterHelper) Table() *table.Writer {
	return h.table
}

func (h *TableAdapterHelper) Tables() []*table.Writer {
	return h.tables
}

func (h *TableAdapterHelper) SwitchTable(id uint8) error {
	if id >= uint8(len(h.tables)) {
		return fmt.Errorf("table with ID %d not found", id)
	}

	h.table = h.tables[id]
	return nil
}

func (*TableAdapterHelper) DefaultTableStyle() table.Style {
	style := table.StyleBold

	style.Format.HeaderAlign = text.AlignCenter
	style.Color.Header = text.Colors{text.Bold}

	return style
}

func (*TableAdapterHelper) ColumnConfig(
	index uint8, align text.Align, width uint8, colors text.Colors,
) table.ColumnConfig {
	return table.ColumnConfig{Number: int(index), Align: align, WidthMin: int(width), Colors: colors}
}

func (h *TableAdapterHelper) SetColumnTableConfigs(configs ...table.ColumnConfig) {
	(*h.table).SetColumnConfigs(configs)
}

func (h *TableAdapterHelper) AppendTableHeader(headers ...any) {
	(*h.table).AppendHeader(append(table.Row{}, headers...))
}

func (h *TableAdapterHelper) AppendTableRow(rows ...any) {
	(*h.table).AppendRow(append(table.Row{}, rows...))
}

func (h *TableAdapterHelper) RenderTable() {
	(*h.table).Render()
}
