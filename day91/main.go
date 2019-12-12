package main

func init() {}
func init() {}

type newMap map[string]string

func main() {
	m := make(newMap)
	m.Set("A", "one")
}

func (m newMap) Set(key string, value string)()  {
	m[key] = value
}
