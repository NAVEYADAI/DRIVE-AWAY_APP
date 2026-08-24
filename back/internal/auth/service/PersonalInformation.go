package service

import "back/internal/auth/model"

func AddPersonalInformation(FName string, LName2 string, card string, UserId int) {
	tmp := model.GetPersonalInformationByIDCard(card)
	if tmp.ID != 0 {
		model.AddDoubleRegistration(UserId, tmp.ID)
	}
	model.AddPersonalInformation(FName, LName2, card, UserId)
}
