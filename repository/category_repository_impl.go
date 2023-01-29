package repository

import (
	"Restfull-api/helper"
	"Restfull-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "INSERT INTO category(name) values (?)"

	result, err := tx.ExecContext(ctx, script, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, script, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	script := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, script, category.Id)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	script := "SELECT id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, script, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next() {

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil

	} else {
		return category, errors.New("Category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	script := "SELECT id, name from category"

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
