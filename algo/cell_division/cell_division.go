package celldivision

// cellDivision return the number of cell after 'hours'
func cellDivision(hours int) int {
	if hours < 0 {
		return -1
	}

	switch hours {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	default:
		return 2*cellDivision(hours-1) - cellDivision(hours-4)
	}
}
