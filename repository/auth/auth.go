package auth

import (
	"errors"
	"gokomodo-api/delivery/middlerwares"
	password2 "gokomodo-api/delivery/password"
	"gokomodo-api/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) LoginSeller(email string, password string) (string, error) {
	var seller entities.Seller
	tx := ar.database.Where("email = ?", email).Find(&seller)
	if tx.Error != nil {
		return "failed", tx.Error
	}

	//jika data user dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}

	//jika ada, maka cek
	if password2.CheckPasswordHash(password, seller.Password) != true {
		return "password incorrect", errors.New("password incorrect")
	}

	//jika password sama
	token, err := middlerwares.CreateToken(int(seller.ID), seller.Name)
	if err != nil {
		return "create token failed", err
	}

	return token, nil

}

func (ar *AuthRepository) LoginBuyer(email string, password string) (string, error) {
	var buyer entities.Buyer

	tx := ar.database.Where("email = ?", email).Find(&buyer)

	if tx.Error != nil {
		return "failed", tx.Error
	}

	//jika data user dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}

	//jika ada, maka cek
	if password2.CheckPasswordHash(password, buyer.Password) != true {
		return "password incorrect", errors.New("password incorrect")
	}

	//jika password sama
	token, err := middlerwares.CreateToken(int(buyer.ID), buyer.Name)
	if err != nil {
		return "create token failed", err
	}

	return token, nil

}
