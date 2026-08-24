package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestGetCompanyByName_FindsExistingCompany(t *testing.T) {
	testutil.SetupTestDB(t)

	// Two rows so a WHERE clause that doesn't actually filter (returning an
	// arbitrary/first row) is distinguishable from one that filters correctly.
	first := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	second := model.CreateCompanyWithReturnId("Globex", "GX-1", "2 Side St")
	require.NotZero(t, first.ID)
	require.NotZero(t, second.ID)

	got := model.GetCompanyByName("Globex")
	require.Equal(t, second.ID, got, "GetCompanyByName should resolve the id of the company actually named Globex, not just any row")
}

func TestGetCompanyByName_UnknownNameReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")

	got := model.GetCompanyByName("DoesNotExist")
	require.Zero(t, got, "GetCompanyByName should return 0 for a name that doesn't exist, not some other row's id")
}

func TestCreateCompany_Basic(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateCompany(model.NewCompany(0, "Acme", "AC-1", "1 Main St"))

	all := model.GetAllCompanies()
	require.Len(t, all, 1)
	require.Equal(t, "Acme", all[0].Name)
}

func TestGetAllCompanies_EmptyWhenNoneCreated(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Empty(t, model.GetAllCompanies())
}

func TestGetCompanyById_UnknownIdReturnsZeroValue(t *testing.T) {
	testutil.SetupTestDB(t)

	got := model.GetCompanyById(999999)
	require.Zero(t, got.ID)
}

func TestUpdateCompany_PersistsChanges(t *testing.T) {
	testutil.SetupTestDB(t)

	created := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	created.Name = "Acme Renamed"
	model.UpdateCompany(created)

	got := model.GetCompanyById(created.ID)
	require.Equal(t, "Acme Renamed", got.Name)
}

func TestDeleteCompany_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	created := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	model.DeleteCompany(created)

	require.Empty(t, model.GetAllCompanies())
}
