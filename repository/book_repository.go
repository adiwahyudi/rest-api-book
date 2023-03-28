package repository

import (
	"chap2-project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) Get() ([]model.Book, error) {
	books := make([]model.Book, 0)

	tx := br.db.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}
func (br *BookRepository) GetOne(id uint) (model.Book, error) {
	book := model.Book{}

	tx := br.db.First(&book, id)
	if tx.Error != nil {
		return model.Book{}, tx.Error
	}
	return book, nil

}

func (br *BookRepository) Save(newBook model.Book) (model.Book, error) {
	tx := br.db.Create(&newBook)
	if tx.Error != nil {
		return model.Book{}, tx.Error
	}
	return newBook, nil
}

func (br *BookRepository) Update(updateBook model.Book, id uint) (model.Book, error) {
	tx := br.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
		},
		).
		Where("id = ?", id).
		Updates(&updateBook)
	if tx.Error != nil {
		return model.Book{}, tx.Error
	}
	return updateBook, nil
}

func (br *BookRepository) Delete(deletedBook model.Book, id uint) error {
	tx := br.db.Delete(&deletedBook, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
