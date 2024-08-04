package core

import (
	"SalaryCalaulater_TwVersion/structure"
	"github.com/BurntSushi/toml"
	"math"
)

// 勞保相關

// getLaborInsurance 取勞保資料
func getLaborInsurance() structure.LaborInsurance {
	var insuranceSalaryGrades structure.LaborInsurance

	if _, tomlErr := toml.DecodeFile("config/toml/insuranceSalaryGrade.toml", &insuranceSalaryGrades); tomlErr != nil {
		println("get Salary Grades Error:", tomlErr.Error())
	}

	return insuranceSalaryGrades
}

// GetLaborInsurancePay 取得勞保自付額(四捨五入)
func GetLaborInsurancePay(salary int) float64 {
	var insuredSalaryGrades = getLaborInsurance() // 取得勞保級距列表

	var insuredSalary int
	for _, grade := range insuredSalaryGrades.InsuranceSalaryGroups {
		if salary >= grade.SalaryRange.Min && salary <= grade.SalaryRange.Max {
			insuredSalary = grade.InsuranceSalary
			break
		}
	}

	// 保險費率
	var laborInsurancePremiumRate = insuredSalaryGrades.LaborSetting.LaborInsurancePremiumRate / 100

	// 受僱員工自付比例(勞工部分負擔)
	var employeeContributionRate = insuredSalaryGrades.LaborSetting.EmployeeContributionRate / 100

	// 計算 -> 勞保級距 * 費率 * 部分負擔
	var insurancePay = float64(insuredSalary) * laborInsurancePremiumRate * employeeContributionRate

	// 四捨五入後回傳
	return math.Ceil(insurancePay)
}
