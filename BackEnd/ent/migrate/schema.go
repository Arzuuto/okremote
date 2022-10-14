// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdvertisementsColumns holds the columns for the "advertisements" table.
	AdvertisementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeString},
		{Name: "contrat", Type: field.TypeString},
		{Name: "entreprise", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "localisation", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "remuneration", Type: field.TypeString},
		{Name: "sector", Type: field.TypeString},
		{Name: "users_advertisement", Type: field.TypeInt, Nullable: true},
	}
	// AdvertisementsTable holds the schema information for the "advertisements" table.
	AdvertisementsTable = &schema.Table{
		Name:       "advertisements",
		Columns:    AdvertisementsColumns,
		PrimaryKey: []*schema.Column{AdvertisementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "advertisements_users_advertisement",
				Columns:    []*schema.Column{AdvertisementsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fist_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdvertisementsTable,
		UsersTable,
	}
)

func init() {
	AdvertisementsTable.ForeignKeys[0].RefTable = UsersTable
}
