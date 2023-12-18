package fil

import "fmt"

type HashTable struct {
	table map[int]string
}

// TODO: this is a new todo
func Hmain() {
	h := newHashTable()

	h.Add(1, "one")
	h.Add(2, "two")
	h.Add(3, "tree")
	h.Add(4, "tree")

	fmt.Println(h.Get(1))
	fmt.Println(h.Get(2))
	fmt.Println(h.Get(3))

}
func newHashTable() *HashTable {
	return &HashTable{
		table: make(map[int]string),
	}
}
func (h *HashTable) Add(key int, value string) {
	h.table[key] = value
}
func (h *HashTable) Get(key int) string {
	return h.table[key]
}
