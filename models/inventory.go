package models

import (
	"encoding/json"
	"errors"
	"os"
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
)

type Object struct {
	ID			int
	Name		string
	Price		int
	CreatedAt	time.Time
}
type Inventory []Object

var nextID int

func (inv *Inventory) Add(name string, price int) {
	item := Object{
		ID : 		nextID,
		Name :		name, 
		Price :		price,
		CreatedAt:	time.Now(),
	}
	nextID++
	*inv = append(*inv, item)
}
func (inv *Inventory) Update(id int, name string, price int) error {
	ls := *inv
	index := inv.getIndexByID(id)
	if index == -1 {
		return errors.New("invalid ID")
	}
	if len(name) != 0 {
		ls[index].Name = name
	}
	if price >= 0 {
		ls[index].Price = price
	}
	
	return nil
}
func (inv *Inventory) Delete(id int) error{
	ls := *inv
	index := inv.getIndexByID(id)
	if index == -1 {
		return errors.New("invalid ID")
	}

	*inv = append(ls[:index], ls[index+1:]...)

	return nil
}
func (inv *Inventory) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("no data found")
	}
	err = json.Unmarshal(data, inv)
	if err != nil {
		return err
	}
	if len(*inv) > 0 {
		maxID := (*inv)[0].ID
		for _, object := range *inv {
			if object.ID > maxID {
				maxID = object.ID
			}
			nextID = maxID + 1
		}
	}
	return nil
}
func (inv *Inventory) Store(filename string) error {
	data, err := json.Marshal(inv)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func (inv *Inventory) Print(){
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Object"},
			{Align: simpletable.AlignCenter, Text: "Price"},
			{Align: simpletable.AlignCenter, Text: "Created at"},
		},
	}
	var cells [][]*simpletable.Cell
	for _, item := range *inv {
		task := item.Name
		price := fmt.Sprintf("%d", item.Price)
		createdAt := item.CreatedAt.Format("02-01-2006 15:04")

		cells = append(cells, []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", item.ID)},
			{Align: simpletable.AlignCenter, Text: task},
			{Align: simpletable.AlignCenter, Text: price},
			{Align: simpletable.AlignCenter, Text: createdAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: ""},
			{Align: simpletable.AlignLeft, Span: 3, Text: fmt.Sprintf("Total items: %d", len(*inv))},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
func (inv *Inventory)getIndexByID(id int)int{
	index := -1
	for i, object := range *inv {
		if object.ID == id {
			index = i
			break
		}
	}
	return index
}