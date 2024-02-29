package main

import (
	"fmt"
	"reflect"
	"strings"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	// pengecekan tipe data variabel reflectValue
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// method Interface() mengembalikan nilai interface kosong atau interface{}
	// nilai aslinya bisa diakses dengan meng-casting interface kosong tersebut.
	var nilai = reflectValue.Interface().(int)
	var jumlah = 2 + nilai

	fmt.Printf("penjumlahan = %d\n", jumlah)

	// =================================
	var name = "danil"
	var reflectValue2 = reflect.ValueOf(name)

	fmt.Println("tipe variabel = ", reflectValue2.Type())

	// pengecekan tipe data variabel reflectValue2
	if reflectValue2.Kind() == reflect.String {
		fmt.Println("nilai variabel :", reflectValue2.String())
	}

	// accessing value using interface method
	fmt.Println("Nilai variabel2 : ", reflectValue2.Interface())

	fmt.Println(strings.Repeat("#", 50))

	for _, v := range []any{"hi", 29, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

	fmt.Println(strings.Repeat("#", 50))

	// identifying method information
	var s1 = &student{Name: "Jhon Wick", Grade: 2}
	fmt.Println("nama", s1.Name)

	var reflectValue3 = reflect.ValueOf(s1)
	var method = reflectValue3.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama", s1.Name)
}
