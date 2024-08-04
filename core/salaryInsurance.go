package core

import (
	"SalaryCalaulater_TwVersion/structure"
	"github.com/BurntSushi/toml"
)

func getInsuranceSalaryGrades() []structure.InsuranceSalaryGroup {
	var insuranceSalaryGrades []structure.InsuranceSalaryGroup

	if _, tomlErr := toml.DecodeFile("SalaryCalaulater_TwVersion/config/toml/insuranceSalaryGrade.toml", &insuranceSalaryGrades); tomlErr != nil {
	}

	return insuranceSalaryGrades
}
