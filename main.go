package main

import (
	"SalaryCalaulater_TwVersion/core"
	"fmt"
)

func main() {

	var salary = 63000

	var labor = core.GetLaborInsurancePay(salary)
	var health = core.GetHealthInsurancePay(salary)

	fmt.Println(fmt.Sprintf("勞保: %v, 健保: %v", labor, health))
}
