package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
	DBError int
}

type Models struct {
	DB DBModel	
}

func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},		
	}
}