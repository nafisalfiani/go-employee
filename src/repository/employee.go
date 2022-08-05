package repository

import (
	"context"
	"math"

	"github.com/nafisalfiani/go-employee/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateEmployee(ctx context.Context, employee model.Employee) error {
	db := getDB()
	err := db.Create(&employee).Error
	if err != nil {
		return err
	}

	return nil
}

func GetEmployee(ctx context.Context, id int) (model.Employee, error) {
	db := getDB()
	var employee model.Employee
	err := db.First(&employee, id).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func GetEmployees(ctx context.Context, page, pageSize int) ([]model.Employee, *model.Pagination, error) {
	db := getDB()
	var employee []model.Employee
	err := db.Limit(pageSize).Offset(page).Find(&employee).Error
	if err != nil {
		return nil, nil, err
	}

	var pg model.Pagination
	if err := db.Model(employee).Count(&pg.TotalElement).Error; err != nil {
		return nil, nil, err
	}

	totalPages := int64(math.Ceil(float64(pg.TotalElement) / float64(pageSize)))
	pg.TotalPages = totalPages

	return employee, &pg, nil
}

func UpdateEmployee(ctx context.Context, employee model.Employee) error {
	db := getDB()
	err := db.Model(&employee).Where("id = ?", employee.ID).Updates(employee).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteEmployee(ctx context.Context, id int) error {
	db := getDB()
	var employee model.Employee
	if err := db.Delete(&employee, id).Error; err != nil {
		return err
	}

	return nil
}

var db *gorm.DB

func getDB() *gorm.DB {
	if db != nil {
		return db
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Employee{
		ID:       1,
		Name:     "Nafisa",
		TeamName: "Pikachu",
	})

	return db
}
