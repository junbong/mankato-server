package data

import "fmt"

type Data struct {
	key     string
	value   string
	expire  uint16
}


func New(key string, value string, expire uint16) (*Data) {
	return &Data{key:key, value:value, expire:expire}
}


func (d Data) String() string {
	return fmt.Sprintf("Data{ key:%s, value:%s, expire:%d }", d.Key(), d.Value(), d.Expire())
}


func (d Data) Key() string {
	return d.key
}


func (d Data) Value() string {
	return d.value
}


func (d Data) Expire() uint16 {
	return d.expire
}
