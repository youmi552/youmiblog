package util

func CheckPageNumber(pageNumber int) int {
	if pageNumber < 1 {
		pageNumber = 1
	}
	return pageNumber
}
