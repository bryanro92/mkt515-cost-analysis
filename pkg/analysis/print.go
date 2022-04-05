package analysis

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const breakdown = `
___.                         __       .___                   
\_ |_________   ____ _____  |  | __ __| _/______  _  ______  
 | __ \_  __ \_/ __ \\__  \ |  |/ // __ |/  _ \ \/ \/ /    \ 
 | \_\ \  | \/\  ___/ / __ \|    </ /_/ (  <_> )     /   |  \
 |___  /__|    \___  >____  /__|_ \____ |\____/ \/\_/|___|  /
     \/            \/     \/     \/    \/                 \/ `

const dash = "------------------------------------------------------------------"

func (s situation) sayInputReceived() {
	fmt.Println(dash)
	fmt.Println(".............Input received.............")
	fmt.Println(".............Computing..................")
	fmt.Println(dash)
}

func (s situation) sayBreakdown() {
	fmt.Println(breakdown)
	fmt.Println(dash)
}

func (s situation) sayDataSummary() {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', 0)
	fmt.Fprintln(w, "FixedCost\tVariableCost\tAnnualRevenue\tAnnualProfit")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", itousd(s.data.fixedCost), itousd(s.data.variableCost), itousd(s.totalAnnualRevenue()), itousd(s.annualProfit()))
	w.Flush()
}

type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "–-----"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}

func itousd(x int) string {
	return fmt.Sprintf("$%.2f", float64(x))
}
