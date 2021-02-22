package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderID    int
	customerID int
}

type employee struct {
	id      int
	name    string
	salary  int
	country string
}

func createQuery(o interface{}) {
	t := reflect.TypeOf(o)
	query := fmt.Sprintf("insert into %s values(", t.Name())

	if t.Kind() == reflect.Struct {
		v := reflect.ValueOf(o)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			switch field.Kind() {
			case reflect.Int:
				query = fmt.Sprintf("%s%d,", query, field.Int())
			case reflect.String:
				query = fmt.Sprintf("%s'%s',", query, field.String())
			default:
				fmt.Println("type is not supported!")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)

	} else {
		fmt.Println("type is not supported!")
	}
}

func main() {
	o := order{
		orderID:    453,
		customerID: 23,
	}

	e := employee{
		id:      34,
		name:    "Vuong",
		salary:  1000,
		country: "Vietnam",
	}

	i := 10

	createQuery(o)
	createQuery(e)
	createQuery(i)
}
