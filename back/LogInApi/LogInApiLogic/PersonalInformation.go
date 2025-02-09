package LogInApiLogic

import "back/LogInApi/LogInApiGormModels"

func AddPersonalInformation(FName string, LName2 string, card string, UserId int) {
	tmp := LogInApiGormModels.GetPersonalInformationByIDCard(card)
	if tmp.ID != 0 {
		LogInApiGormModels.AddDoubleRegistration(UserId, tmp.ID)
	}
	LogInApiGormModels.AddPersonalInformation(FName, LName2, card, UserId)
}
