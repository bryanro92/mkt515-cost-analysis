package analysis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/leekchan/accounting"
)

type situation struct {
	answers   []string
	questions []string
	data      data
	a         accounting.Accounting
}

type data struct {
	fixedCost         int
	variableCost      int
	totalTargetMarket int
	marketAwareness   float64
	avgRevenuePerCx   int
}

func Run(ctx context.Context) error {
	var analysisMgr = new(situation)

	err := analysisMgr.init()
	if err != nil {
		return err
	}
	analysisMgr.sayInputReceived()

	fmt.Println(analysisMgr.analysis())
	return nil
}

func (s situation) analysis() string {
	var bar Bar
	bar.NewOption(0, 10)
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
	// s.sayBreakdown()
	s.sayDataSummary()

	fmt.Sprintf("Break Even (days): %.2f", float64(s.breakEvenDays()))
	fmt.Println("")

	if s.annualProfit() < 0 {
		return "Do not advise: total annual cost: $" + strconv.Itoa(s.totalAnnualCost()) + " is greater than total revenue: $" + strconv.Itoa(s.totalAnnualRevenue())
	}
	if s.breakEvenDays() >= 1825 {
		return "Do not advise: break even point is greater than 5 years"
	}
	return "Should proceed as planned."
}

func (s situation) profitPerCustomer() int {
	return s.data.avgRevenuePerCx - s.data.variableCost
}

func (s *situation) breakEvenDays() float64 {
	return float64(s.data.fixedCost) / s.dailyProfit()
}

func (s situation) totalAnnualCost() int {
	return s.data.variableCost*s.cxReach() + s.data.fixedCost
}

func (s situation) margin() string {
	return fmt.Sprintf("%0.2f", (float64(s.data.avgRevenuePerCx) / float64(s.data.variableCost)))
}

func (s situation) totalAnnualRevenue() int {
	return s.cxReach() * s.data.avgRevenuePerCx
}

func (s situation) annualProfit() int {
	return s.totalAnnualRevenue() - s.totalAnnualCost()
}

func (s situation) cxReach() int {
	return int(s.data.marketAwareness * float64(s.data.totalTargetMarket))
}

func (s situation) cxReachString() string {
	return humanize.Comma(int64(s.data.marketAwareness * float64(s.data.totalTargetMarket)))
}

func (s situation) dailyRevenue() float64 {
	n := 365
	return float64(float64(s.totalAnnualRevenue()) / float64(n))
}

func (s situation) dailyProfit() float64 {
	n := 365
	return float64(s.annualProfit()) / float64(n)
}
