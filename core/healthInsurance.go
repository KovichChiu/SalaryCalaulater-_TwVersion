package core

import (
	"SalaryCalaulater_TwVersion/structure"
	"github.com/BurntSushi/toml"
	"math"
)

func getHealthInsurance() structure.HealthInsurance {
	var insuranceHealthGrades structure.HealthInsurance

	if _, tomlErr := toml.DecodeFile("config/toml/insuranceHealthGrade.toml", &insuranceHealthGrades); tomlErr != nil {
		println("get Salary Grades Error:", tomlErr.Error())
	}

	return insuranceHealthGrades
}

func GetHealthInsurancePay(salary int) float64 {
	var healthInsuranceGrades = getHealthInsurance()

	// 取 保費級距
	groups := healthInsuranceGrades.InsuranceSalaryGroups
	var cost int

	for _, group := range groups {
		// 符合級距再進去跑迴圈
		if salary >= group.Min && salary <= group.Max {
			for _, level := range group.Levels {
				if salary >= level.Min && salary <= level.Max {
					cost = level.InsuranceSalary
				}
			}
		}
	}

	// 取 費率們
	healthInsurancePremiumRate := healthInsuranceGrades.HealthSetting.HealthInsurancePremiumRate / 100
	employeeContributionRate := healthInsuranceGrades.HealthSetting.EmployeeContributionRate / 100

	// 取 加保人數(超過三人以三人計)
	insured := 1       // 被保險人
	insuredFamily := 0 // 一起加保的被保險人家屬
	if insuredFamily >= 3 {
		insuredFamily = 3
	}

	// 計算
	var insurancePay = (float64(cost) * healthInsurancePremiumRate * employeeContributionRate) * float64(insured+insuredFamily)

	return math.Ceil(insurancePay)
}
