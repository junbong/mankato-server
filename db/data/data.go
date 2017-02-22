package data

type Data struct {
	key     string
	value   string
	expire  uint16
}


func New(key string, value string, expire uint16) (*Data) {
	return &Data{key:key, value:value, expire:expire}
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
