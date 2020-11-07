package pkg

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

type View struct{}

func (v View) Display(stocks []Stock) {
	var stockRows []table.Row
	for _, s := range stocks {
		tmp := table.Row{s.Symbol, s.Price, s.DollarChange, s.PercentChange}
		stockRows = append(stockRows, tmp)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Symbol", "Price", "$ Change", "% Change"})
	t.AppendRows(stockRows)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        "Price",
			Transformer: text.NewNumberTransformer("%.2f"),
			Align: text.AlignRight,
			WidthMin: 20,
		},
		{
			Name:        "$ Change",
			Transformer: text.NewNumberTransformer("%.2f"),
			Align: text.AlignRight,
			WidthMin: 20,
		},
		{
			Name:        "% Change",
			Transformer: text.NewNumberTransformer("%.4f%%"),
			Align: text.AlignRight,
			WidthMin: 20,
		},
	})

	t.SetStyle(table.StyleColoredDark)
	t.Render()
}
