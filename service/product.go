package service

import (
	"tpm-tujuh/model"
	"tpm-tujuh/repository"
)

type ProductService struct {
	ProductPgRepo *repository.ProductPgRepo
}

func (u *ProductService) Get(userId int) ([]*model.Product, error) {
	return u.ProductPgRepo.Get(userId)
}

func (s *ProductService) Create(product *model.Product) (int, error) {
	return s.ProductPgRepo.Create(product)
}

func (u *ProductService) Update(id int, uid int, product *model.ProductUpdate) error {
	return u.ProductPgRepo.Update(id, uid, product)
}

func (u *ProductService) Delete(id int, uid int) error {
	return u.ProductPgRepo.Delete(id, uid)
}
