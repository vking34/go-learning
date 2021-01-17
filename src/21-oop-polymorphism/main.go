package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	cpc        int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (ad Advertisement) source() string {
	return ad.adName
}

func (ad Advertisement) calculate() int {
	return ad.cpc * ad.noOfClicks
}

func calculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}

	fmt.Printf("Net income of organisation $%d\n", netIncome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 3000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 5000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	project4 := Advertisement{adName: "Ads 1", cpc: 2, noOfClicks: 3000}
	incomeStreams := []Income{project1, project2, project3, project4}
	calculateNetIncome(incomeStreams)
}
