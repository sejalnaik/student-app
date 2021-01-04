package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
)

type gormRepository struct {
}

func NewRepository() Repository {
	return &gormRepository{}
}

type Repository interface {
	Get(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
	Add(uow *UnitOfWork, entity interface{}) error
	Update(uow *UnitOfWork, entity interface{}, entityMap map[string]interface{}, queryProcessors []QueryProcessor) error
	Delete(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) error
	Select(uow *UnitOfWork, condition string, entity interface{}) error
	//BookSpecial(uow *UnitOfWork, entity interface{}) error
	Scan(uow *UnitOfWork, out interface{}, entity interface{}, queryProcessors []QueryProcessor) error
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Where(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Model(out).Where(condition, args...)
		return db, nil
	}
}

func PreloadAssociations(preloadAssociations []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}

func Model() QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Model(out)
		return db, nil
	}
}

func Select(condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Select(condition)
		return db, nil
	}
}

func Joins(condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Joins(condition)
		return db, nil
	}
}

func Group(condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Group(condition)
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

func (r *gormRepository) Update(uow *UnitOfWork, entity interface{}, entityMap map[string]interface{}, queryProcessors []QueryProcessor) error {
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
	if err := db.Debug().Updates(entityMap).Error; err != nil {
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

func (r *gormRepository) Select(uow *UnitOfWork, condition string, entity interface{}) error {
	db := uow.DB

	if err := db.Debug().Model(&model.Student{}).Select(condition).Scan(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) Scan(uow *UnitOfWork, entity interface{}, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return err
			}
		}
	}

	if err := db.Debug().Scan(out).Error; err != nil {
		return err
	}
	return nil
}

/*func (r *gormRepository) BookSpecial(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB

	if err := db.Debug().Model(&model.Book{}).Select("books.ID as id, books.name as name, MIN(total_stock) as total_stock, (MIN(total_stock) - COUNT(book_issues.returned)) AS available").
		Joins("inner join book_issues on book_issues.book_id = books.id").
		Joins("inner join students on book_issues.student_id = students.id").
		Where("book_issues.returned = 0").
		Group("books.id").
		Group("books.name").
		Scan(entity).Error; err != nil {
		return err
	}
	return nil
}*/
