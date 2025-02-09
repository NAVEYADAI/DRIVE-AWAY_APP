package MainControllerGormModels

import "back/MainController/config"

type Company struct {
	ID        int    `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	IdCompany string `gorm:"column:idcompany"`
	Address   string `gorm:"column:address"`
}

func (receiver Company) TableName() string {
	return "company"
}
func NewCompany(id int, name, idCompany, address string) Company {
	return Company{id, name, idCompany, address}
}
func (receiver Company) GetId() int {
	return receiver.ID
}
func (receiver Company) GetName() string {
	return receiver.Name
}
func (receiver Company) GetIdCompany() string {
	return receiver.IdCompany
}
func (receiver Company) GetAddress() string {
	return receiver.Address
}
func (receiver Company) SetName(name string) {
	receiver.Name = name
}
func (receiver Company) SetIdCompany(idCompany string) {
	receiver.IdCompany = idCompany
}
func (receiver Company) SetAddress(address string) {
	receiver.Address = address
}

// Gorm functions
func CreateCompany(receiver Company) {
	config.DB.Create(&receiver)
}
func CreateCompanyWithReturnId(companyName, companyId, address string) Company {
	var company Company
	company = Company{Name: companyName, IdCompany: companyId, Address: address}
	config.DB.Create(&company)
	return company
}
func GetAllCompanies() []Company {
	var companies []Company
	config.DB.Find(&companies)
	return companies
}
func GetCompanyById(id int) Company {
	var company Company
	config.DB.Where("ID = ?", id).Find(&company)
	return company
}
func GetCompanyByName(name string) int {
	var company Company
	config.DB.Where("name", name).Find(&company)
	return company.ID
}
func UpdateCompany(receiver Company) {
	config.DB.Save(&receiver)
}
func DeleteCompany(receiver Company) {
	config.DB.Delete(&receiver)
}
