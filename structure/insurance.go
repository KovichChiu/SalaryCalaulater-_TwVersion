package structure

type Range struct {
	Min int `toml:"min"`
	Max int `toml:"max"`
}

type InsuranceSalaryGroup struct {
	Grade                  int `toml:"grade"`
	MonthlySalaryTotal     int `toml:"monthly_salary_total"`
	MonthlyInsuranceSalary int `toml:"monthly_insurance_salary"`
}
