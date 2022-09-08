package util

import (
	"bufio"
	"os"
	"strings"
)

type Database struct {
	filePath string
	file     *os.File
	data     map[string]string
	sperator string
}

func GetDatabase(filePath string, sperator string) Database {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	db := Database{filePath, file, make(map[string]string), sperator}
	db.ReadData()
	return db
}

func (db Database) ReadData() {
	scanner := bufio.NewScanner(db.file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, db.sperator) {
			continue
		}
		strings := strings.Split(line, db.sperator)
		db.data[strings[0]] = strings[1]
	}
}

func (db Database) WriteData() {
	writer := bufio.NewWriter(db.file)
	for key, value := range db.data {
		writer.WriteString(key + db.sperator + value + "\n")
	}
	writer.Flush()
}

func (db Database) Close() {
	db.file.Close()
}

func (db Database) Get(key string) string {
	if value, ok := db.data[key]; ok {
		return value
	} else {
		return ""
	}
}

func (db Database) Set(key string, value string) {
	db.data[key] = value
}
