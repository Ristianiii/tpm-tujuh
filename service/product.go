package service

import (
	"hari-ketiga/tpm-keenam/model"
	"hari-ketiga/tpm-keenam/repository"
)

type ProductService struct {
	ProductPgRepo *repository.ProductPgRepo
}

func (u *ProductService) Get() ([]*model.Product, error) {
	return u.ProductPgRepo.Get()
}

func (s *ProductService) Create(product *model.Product) (int, error) {
	return s.ProductPgRepo.Create(product)
}

func (u *ProductService) Update(id int, product *model.ProductUpdate) error {
	return u.ProductPgRepo.Update(id, product)
}

func (u *ProductService) Delete(id int) error {
	return u.ProductPgRepo.Delete(id)
}
