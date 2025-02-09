package test

import (
	"back/LogInApi/LogInApiGormModels"
	"back/LogInApi/config"
	"testing"
)

func TestGetFNameByUserId(t *testing.T) {
	config.UpDBLogIn()
	tmp := LogInApiGormModels.GetFNameByUserID(5)
	println(tmp)
}
