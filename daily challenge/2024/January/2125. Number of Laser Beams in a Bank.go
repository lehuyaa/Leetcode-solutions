package main

func countLaze(row string) int {
	count := 0
	for i := 0; i < len(row); i++ {
		if row[i] == '1' {
			count++
		}
	}
	return count
}

func numberOfBeams(bank []string) int {
	previousRow := 0
	numberBeams := 0
	for i := 0; i < len(bank); i++ {
		currentRow := countLaze(bank[i])
		if currentRow == 0 {
			continue
		}
		numberBeams += previousRow * currentRow
		previousRow = currentRow
	}
	return numberBeams
}
