package repository

import "github.com/jinzhu/gorm"

type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{
			DB:        db.New(),
			committed: false,
			readOnly:  true,
		}
	} else {
		return &UnitOfWork{
			DB:        db.New().Begin(),
			committed: false,
			readOnly:  false,
		}
	}
}

func (uow *UnitOfWork) Commit() {
	if uow.readOnly == false {
		uow.DB.Commit()
	}
	uow.committed = true
}

func (uow *UnitOfWork) Complete() {
	if uow.readOnly == false && uow.committed == false {
		uow.DB.Rollback()
	}
}
