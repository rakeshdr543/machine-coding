package main

import (
	"fmt"
	"time"

	tracker "github.com/rakeshdr543/machine_coding/pandemic_tracker"
)

func main() {

	pandemicTracker := tracker.NewTracker()

	name1, name2, name3 := "john", "bob", "Rock"

	pandemicTracker.RegisterUser(name1, 897897987, 560056)
	pandemicTracker.RegisterUser(name2, 89789765567, 560056)

	pandemicTracker.RegisterUser(name3, 656656567, 560055)

	assessRes1, err := pandemicTracker.SelfAssessment(name2, []string{"Cough", "Fever"}, true, true)
	if err != nil {
		fmt.Printf("error occurred %s\n", err.Error())
	}

	fmt.Printf("Assessment result for %s is %d\n", name2, assessRes1)

	assessRes2, err := pandemicTracker.SelfAssessment(name1, []string{}, false, false)
	if err != nil {
		fmt.Printf("error occurred %s\n", err.Error())
	}
	fmt.Printf("Assessment result for %s is %d\n", name1, assessRes2)

	adminId := "admin1"

	pandemicTracker.PandemicResult(adminId, name1, true, time.Now())
	pandemicTracker.PandemicResult(adminId, name2, true, time.Now())
	pandemicTracker.PandemicResult(adminId, name3, true, time.Now())

	zone1Res, err := pandemicTracker.GetZone(adminId, 560056)
	if err != nil {
		fmt.Printf("error occurred %s\n", err.Error())
	}
	fmt.Printf("Pandemic result for zone %d is %s\n", 560056, zone1Res)

	zone2Res, err := pandemicTracker.GetZone(adminId, 560055)
	if err != nil {
		fmt.Printf("error occurred %s\n", err.Error())
	}
	fmt.Printf("Pandemic result for zone %d is %s\n", 560055, zone2Res)

}
