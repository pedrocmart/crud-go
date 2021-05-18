package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pedrocmart/crud-go/api/models"
)

type UserRepository interface {
	Save(*models.User) (*models.User, error)
	Find(uint64) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(*models.User) error
	Delete(uint64) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Save(user *models.User) (*models.User, error) {
	tx := r.db.Debug().Begin()
	err := tx.Debug().Model(&models.User{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, tx.Commit().Error
}

func (r *userRepositoryImpl) Find(user_id uint64) (*models.User, error) {
	user := &models.User{}
	err := r.db.Debug().Model(&models.User{}).Where("id = ?", user_id).Find(user).Error
	return user, err
}

func (r *userRepositoryImpl) FindAll() ([]*models.User, error) {
	user := []*models.User{}
	err := r.db.Debug().Model(&models.User{}).Find(&user).Error
	return user, err
}

func (r *userRepositoryImpl) Update(user *models.User) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"name":        user.Name,
		"description": user.Description,
		"dob":         user.DOB,
		"address":     user.Address,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.User{}).Where("id = ?", user.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *userRepositoryImpl) Delete(user_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.User{}).Where("id = ?", user_id).Delete(&models.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
