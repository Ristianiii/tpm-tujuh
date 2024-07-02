package repository

import (
	"errors"
	"tpm-tujuh/model"

	"gorm.io/gorm"
)

type ProductPgRepo struct {
	DB *gorm.DB
}

func (s *ProductPgRepo) Get(userId int) ([]*model.Product, error) {
	products := []*model.Product{}
	err := s.DB.Debug().Where("user_id = ?", userId).Find(&products).Error
	return products, err
}

func (r *ProductPgRepo) Create(product *model.Product) (int, error) {
	err := r.DB.Debug().Create(&product).Error
	if err != nil {
		return 0, err
	}
	return product.Id, err
}

func (r *ProductPgRepo) Update(id int, uid int, productUpdate *model.ProductUpdate) error {
	result := r.DB.Debug().
		Where("id = ? AND user_id = ?", id, uid).
		Updates(model.Product{
			Name:  productUpdate.Name,
			Price: productUpdate.Price,
		})
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return result.Error
}

func (s *ProductPgRepo) Delete(id int, uid int) error {
	result := s.DB.Debug().
		Where("id = ? AND user_id = ?", id, uid).
		Delete(&model.Product{})

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return result.Error
}
