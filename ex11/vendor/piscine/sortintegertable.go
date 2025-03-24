package piscine

func lentable(table []int) int {
	i := 0
	for range table {
		i++
	}
	return i
}

func swap(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func SortIntegerTable(table []int) {
	i := 0
	change := true
	len := lentable(table);
	for change != false {
		change = false
		for i < len - 1 {
			if (table[i] > table[i + 1]) {
				swap(&table[i], &table[i + 1])
				change = true
			}
			i++
		}
		i = 0
	}
}