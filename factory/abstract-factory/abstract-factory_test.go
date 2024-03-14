package abstractfactory

import "testing"

func TestAbstract(t *testing.T) {
	companyFactoryStore := CompanyFactoryStore{
		shower: ChinaComFactory{},
	}
	companyFactoryStore.shower.ShowCompany("bilibili").slogan()
}
