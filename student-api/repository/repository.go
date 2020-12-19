package repository

type gormRepository struct {
}

func NewRepository() Repository {
	return &gormRepository{}
}

type Repository interface {
	Get(uow *UnitOfWork, out interface{}) error
	Add(uow *UnitOfWork, entity interface{}) error
	GetFirst(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, entity interface{}) error
	Delete(uow *UnitOfWork, entity interface{}) error
}

func (*gormRepository) Get(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

func (*gormRepository) GetFirst(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	if err := db.Debug().First(out).Error; err != nil {
		return err
	}
	return nil
}

func (*gormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
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

func (*gormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()
	if err := db.Error; err != nil {
		return err
	}
	if err := db.Debug().Model(entity).Update(entity).Error; err != nil {
		return err
	}
	return nil
}

func (*gormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()
	if err := db.Error; err != nil {
		return err
	}
	if err := db.Debug().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}
