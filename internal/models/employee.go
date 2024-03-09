package models

import (
	"fmt"
)

type Employee struct {
	Id    int
	Name  string
	Phone string
}

func (e *Employee) Validate() error {
	if e.Id < 0 {
		return fmt.Errorf("неверное значение Id (id должен быть больше или равен 0): %d", e.Id)
	}
	if e.Name == "" {
		return fmt.Errorf("поле Name не может быть пустым")
	}
	if e.Phone == "" {
		return fmt.Errorf("поле Phone не может быть пустым")
	}
	return nil
}
