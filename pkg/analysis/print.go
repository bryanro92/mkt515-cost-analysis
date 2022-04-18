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
	fmt.Fprintln(w, "FixedCost\tVariableCost\tProfitPerCx\tTotalNumCx\tAnnualExpenses\tDailyRev\tDailyProfit\tAnnualRevenue\tAnnualProfit\tMargin\tBreakEvenDays")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
		s.itousd(s.data.fixedCost),
		s.itousd(s.data.variableCost),
		s.itousd(s.profitPerCustomer()),
		s.cxReachString(),
		s.itousd(s.totalAnnualCost()),
		s.itousd(int(s.dailyRevenue())),
		s.itousd(int(s.dailyProfit())),
		s.itousd(s.totalAnnualRevenue()),
		s.itousd(s.annualProfit()),
		s.margin(),
		int(s.breakEvenDays()))
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

func (s situation) itousd(x int) string {
	return s.a.FormatMoney(x)
}
