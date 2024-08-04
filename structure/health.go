package structure

type HealthLevel struct {
	Id              int `toml:"id"`
	Min             int `toml:"min"`
	Max             int `toml:"max"`
	InsuranceSalary int `toml:"insurance_salary"`
}

type HealthGroup struct {
	Id     int           `toml:"id"`
	Cost   int           `toml:"cost"`
	Min    int           `toml:"min"`
	Max    int           `toml:"max"`
	Levels []HealthLevel `toml:"levels"`
}

type HealthSetting struct {
	HealthInsurancePremiumRate float64 `toml:"health_insurance_premium_rate"`
	EmployeeContributionRate   float64 `toml:"employee_contribution_rate"`
}

type HealthInsurance struct {
	HealthSetting         HealthSetting `toml:"health_setting"`
	InsuranceSalaryGroups []HealthGroup `toml:"health_insurance_group"`
}
