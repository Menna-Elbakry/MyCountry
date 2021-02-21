package table

import (
	"task/model"
)

//Users table
type Country struct {
	tableName struct{} `sql:"country"`
	Name      string   `sql:"name"`
	Location  string   `sql:"location"`
	Wiki      string   `sql:"wiki"`
}

func (c *Country) MapToModule() model.Country {
	return model.Country{
		Name:     c.Name,
		Location: c.Location,
		Wiki:     c.Wiki,
	}
}

func (c Country) Fill(con *model.Country) *Country {
	return &Country{
		Name:     con.Name,
		Location: con.Location,
		Wiki:     con.Wiki,
	}
}
