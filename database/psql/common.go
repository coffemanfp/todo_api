package psql

func parsePagination(page int) (limit, offset int) {
	// TODO: Make this configurable
	return (page + 1) * 20, page * 20
}
