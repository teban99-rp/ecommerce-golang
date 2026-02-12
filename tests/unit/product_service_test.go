func TestCreateProduct_InvalidPrice(t *testing.T) {
	repo := mocks.NewProductRepositoryMock()
	service := ProductServiceImpl{Repo: repo}

	err := service.CreateProduct(models.Product{
		Name:  "Test",
		Price: 0,
	})

	assert.NotNil(t, err)
}
