package database

import (
	"tokokecilkita-go/order"
	"tokokecilkita-go/organization"
	"tokokecilkita-go/product"
	"tokokecilkita-go/user"
)

type Table struct {
	Table interface{}
}

func RegisterTables() []Table {
	return []Table{
		{Table: user.User{}},
		{Table: organization.Organization{}},
		{Table: product.Product{}},
		{Table: order.Order{}},
		{Table: order.OrderProduct{}},
	}
}
