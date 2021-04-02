package repository

import (
	model "github.com/birdream/go-micro-service/imooc-stu/user/domain/model"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

//创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

//根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

//根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

//创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

//根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

//更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

//查找所有用
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
