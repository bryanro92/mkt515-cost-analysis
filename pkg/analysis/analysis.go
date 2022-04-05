package analysis

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type situation struct {
	answers   []string
	questions []string
	data      data
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
	s.sayBreakdown()
	s.sayDataSummary()

	fmt.Sprintf("Break Even (days): %.2f", float64(s.breakEvenDays()))
	fmt.Println("")

	if s.annualProfit() > s.totalAnnualRevenue() {
		return "Do not advise: total annual cost: $" + strconv.Itoa(s.totalAnnualCost()) + " is greater than total revenue: $" + strconv.Itoa(s.totalAnnualRevenue())
	}
	if s.breakEvenDays() >= 1825 {
		return "Do not advise: break even point is grater than 5 years"
	}
	return "Should proceed as planned."
}

func (s *situation) breakEvenDays() float64 {
	return float64(s.data.fixedCost) / s.dailyRevenue()
}

func (s situation) totalAnnualCost() int {
	return s.data.variableCost * s.data.totalTargetMarket
}

func (s situation) totalAnnualRevenue() int {
	return s.cxReach() * s.data.avgRevenuePerCx
}

func (s situation) annualProfit() int {
	return s.totalAnnualRevenue() - s.data.fixedCost - s.data.totalTargetMarket*s.data.variableCost
}

func (s situation) cxReach() int {
	return int(s.data.marketAwareness * float64(s.data.totalTargetMarket))
}

func (s situation) dailyRevenue() float64 {
	n := 365
	return float64(float64(s.totalAnnualRevenue()) / float64(n))
}
