package main

import (
	"flag"
	"fmt"
	"github.com/glauber-silva/farmers-markets/internal/database"
	"github.com/glauber-silva/farmers-markets/internal/markets"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

var (
	filename *string
)

func init() {
	filename = flag.String("filename",  "", "Name of the file seed")
}

func main() {
	flag.Parse()
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	path :=  basepath + "/files/" + *filename
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var entries = []markets.Market{}
	err = gocsv.Unmarshal(file, &entries)

	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)

	}
	log.Info("Database Connected. Starting to insert data")
	result := db.Create(entries)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Infof("Database populated with %d new row(s)" , result.RowsAffected)
}