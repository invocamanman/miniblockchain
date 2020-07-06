package main

func main() {
	db := openDb()
	defer db.Close()

	server(db)
}
