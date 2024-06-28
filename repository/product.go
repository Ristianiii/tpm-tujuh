package repository

import (
	"fmt"
	"hari-ketiga/tpm-keenam/model"

	"gorm.io/gorm"
)

type ProductPgRepo struct {
	DB *gorm.DB
}

func (s *ProductPgRepo) Get() ([]*model.Product, error) {
	products := []*model.Product{}
	err := s.DB.Debug().Find(&products).Error
	return products, err
}

func (r *ProductPgRepo) Create(product *model.Product) (int, error) {
	err := r.DB.Debug().Create(&product).Error
	if err != nil {
		return 0, err
	}
	return product.Id, err
}

func (r *ProductPgRepo) Update(id int, productUpdate *model.ProductUpdate) error {
	result := r.DB.Debug().
		Where("id = ?", id).
		Updates(model.Product{
			Name:  productUpdate.Name,
			Price: productUpdate.Price,
		})
	fmt.Printf("Rows affected: %v", result.RowsAffected)
	return result.Error
}

func (s *ProductPgRepo) Delete(id int) error {
	err := s.DB.Debug().
		Where("id = ?", id).
		Delete(&model.Product{}).Error
	return err
}
