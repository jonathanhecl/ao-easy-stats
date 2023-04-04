package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

const (
	Version = "0.0.1"
)

func main() {
	fmt.Println("Starting AO Easy Stats generator v" + Version)

	var (
		days       = 15 // days to show
		line       = charts.NewLine()
		labels     = generateDateItems(days)
		characters = make([]opts.LineData, 0)
		records    = make([]opts.LineData, 0)
		maxValue   int
	)

	// get data
	for l := range labels {
		chars, record := GetTotalCharsByDay(labels[l])
		characters = append(characters, opts.LineData{Value: chars, Name: labels[l]})
		records = append(records, opts.LineData{Value: record, Name: labels[l]})
		if chars > maxValue {
			maxValue = chars
		} else if record > maxValue {
			maxValue = record
		}
	}

	// set global options
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "90%",
			Height: "300px",
			Theme:  types.ThemeInfographic,
		}),
		charts.WithYAxisOpts(opts.YAxis{Type: "value", SplitNumber: maxValue}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    true,
			Trigger: "axis",
			AxisPointer: &opts.AxisPointer{
				Snap: true,
			},
		}),
	)

	// invert labels / characters / records
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		labels[i], labels[j] = labels[j], labels[i]
		characters[i], characters[j] = characters[j], characters[i]
		records[i], records[j] = records[j], records[i]
	}

	// set axis & series
	line.SetXAxis(labels)
	line.AddSeries("Personajes logueados", characters, charts.WithItemStyleOpts(
		opts.ItemStyle{Color: "#84b6f4"}))
	line.AddSeries("Record de conectados", records, charts.WithItemStyleOpts(
		opts.ItemStyle{Color: "#77dd77"}))

	// set series options
	line.SetSeriesOptions(
		charts.WithMarkLineNameTypeItemOpts(
			opts.MarkLineNameTypeItem{Name: "Max", Type: "max"}),
		charts.WithMarkPointNameTypeItemOpts(
			opts.MarkPointNameTypeItem{Name: "Max", Type: "max"}),
		charts.WithAreaStyleOpts(opts.AreaStyle{Opacity: 0.5}),
		//charts.WithMarkPointStyleOpts(
		//	opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		charts.WithLineChartOpts(opts.LineChart{Smooth: true}), //  Step: "middle"
		charts.WithLabelOpts(opts.Label{Show: false, Formatter: "{value}"}))

	// render
	f, _ := os.Create("stats.html")
	err := line.Render(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
