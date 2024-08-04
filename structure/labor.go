package structure

// Range 薪資範圍
type Range struct {
	Min int `toml:"min"`
	Max int `toml:"max"`
}

// LaborInsuranceGrade 勞保投保薪資等級
type LaborInsuranceGrade struct {
	Grade           int   `toml:"grade"`
	SalaryRange     Range `toml:"salary_range"`
	InsuranceSalary int   `toml:"insured_salary"`
}

type LaborSetting struct {
	LaborInsurancePremiumRate float64 `toml:"labor_insurance_premium_rate"`
	EmployeeContributionRate  float64 `toml:"employee_contribution_rate"`
}

// LaborInsurance 勞保
type LaborInsurance struct {
	LaborSetting          LaborSetting          `toml:"labor_setting"`
	InsuranceSalaryGroups []LaborInsuranceGrade `toml:"labor_insurance_grade"` // 投保薪資等級列表
}
