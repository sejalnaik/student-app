package repository

import (
	"github.com/jinzhu/gorm"
)

type gormRepository struct {
}

func NewRepository() Repository {
	return &gormRepository{}
}

type Repository interface {
	Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(value interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where("id = ?", value)
		return db, nil
	}
}

func Model(out interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Model(out)
		return db, nil
	}
}

func (r *gormRepository) Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB
	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			uow.Complete()
		}
	}()
	if err := db.Error; err != nil {
		return err
	}
	if err := db.Debug().Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) Update(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			uow.Complete()
		}
	}()
	if err := db.Error; err != nil {
		return err
	}
	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Update(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			uow.Complete()
		}
	}()
	if err := db.Error; err != nil {
		return err
	}
	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}
