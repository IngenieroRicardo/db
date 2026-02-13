package main

import (
	"fmt"
	"time"
	db "github.com/IngenieroRicardo/db/go"
)

func main() {

	conn, err := db.LoadSQL("sqlite3", ":memory:", 670, 100, 10*time.Minute, 10*time.Minute)
	if err == nil {
		
		db.SQLrunonLoad(conn, "CREATE TABLE usuario(id INTEGER PRIMARY KEY AUTOINCREMENT, nickname VARCHAR(30) NOT NULL UNIQUE, picture BLOB);")

		argumento1 := "Ricardo"
		argumento2 := "blob::iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAIAAABLbSncAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAArSURBVBhXY/iPA0AlGBgwGFAKlwQmAKrAIgcVRZODCsI5cAAVgVDo4P9/AHe4m2U/OJCWAAAAAElFTkSuQmCC"
		db.SQLrunonLoad(conn, "INSERT INTO usuario(nickname, picture) VALUES (?, ?);", argumento1, argumento2)

		result1 := db.SQLrunonLoad(conn, "SELECT * FROM usuario")

		fmt.Printf("JSON: %s\n", result1.Json)

		db.CloseSQL(conn)
	}	
}