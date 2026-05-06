package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful_api/helper"
	"restful_api/model/domain"
)

type CategoryRepositoryImpln struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpln{}
}

func (c *CategoryRepositoryImpln) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.ErrorT(err)

	id, err := result.LastInsertId()
	helper.ErrorT(err)
	category.Id = int(id)
	return category

}
func (c *CategoryRepositoryImpln) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name =? where id =?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.ErrorT(err)
	return category
}
func (c *CategoryRepositoryImpln) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.ErrorT(err)

}
func (c *CategoryRepositoryImpln) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id =?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.ErrorT(err)

	defer rows.Close() // <--- WAJIB TAMBAHKAN INI

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorT(err)
		return category, nil
	} else {
		return category, errors.New("Category is not found")
	}
}
func (c *CategoryRepositoryImpln) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from category"
	rows, err := tx.QueryContext(ctx, SQL)

	helper.ErrorT(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorT(err)
		categories = append(categories, category)

	}
	return categories
}
