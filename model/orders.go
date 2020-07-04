package model

import (
	"database/sql"
)


/*
DB Table Details
-------------------------------------


CREATE TABLE orders
(
	system_id integer
		constraint orders_pk
			primary key autoincrement,
	customer_name text,
	file_name text,
	department text,
	material_id text,
	maker text,
	process text,
	create_time integer,
	deadline_time integer
, order_status integer, admin_status integer)

JSON Sample
-------------------------------------
{    "department": "OpHuuLvCVJubZdAedviSlYeuV",    "maker": "IWxIepDMcNgxQJifVoIwYXceZ",    "deadline_time": 45,    "system_id": 15,    "file_name": "WHJZFdHBWKLantdhYHDBfKwZV",    "process": "DmeKfboNTtqjqdHyRhmfSQLeZ",    "create_time": 18,    "order_status": 97,    "admin_status": 66,    "customer_name": "FoWXEIldbXurTcrGDrTTLFnSM",    "material_id": "JmagLmaNPnwgvIHpyOZDebdsk"}



*/

// Orders struct is a row record of the orders table in the main database
type Orders struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int32 `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] customer_name                                  text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	CustomerName sql.NullString `gorm:"column:customer_name;type:text;" json:"customer_name"`
	//[ 2] file_name                                      text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	FileName sql.NullString `gorm:"column:file_name;type:text;" json:"file_name"`
	//[ 3] department                                     text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Department sql.NullString `gorm:"column:department;type:text;" json:"department"`
	//[ 4] material_id                                    text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	MaterialID sql.NullString `gorm:"column:material_id;type:text;" json:"material_id"`
	//[ 5] maker                                          text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Maker sql.NullString `gorm:"column:maker;type:text;" json:"maker"`
	//[ 6] process                                        text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Process sql.NullString `gorm:"column:process;type:text;" json:"process"`
	//[ 7] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime sql.NullInt64 `gorm:"column:create_time;type:integer;" json:"create_time"`
	//[ 8] deadline_time                                  integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	DeadlineTime sql.NullInt64 `gorm:"column:deadline_time;type:integer;" json:"deadline_time"`
	//[ 9] order_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	OrderStatus sql.NullInt64 `gorm:"column:order_status;type:integer;" json:"order_status"`
	//[10] admin_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	AdminStatus sql.NullInt64 `gorm:"column:admin_status;type:integer;" json:"admin_status"`
}