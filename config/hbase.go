package config

import (
	"errors"
	"fmt"
	"log"

	"os"
	"strings"

	_ "github.com/apache/calcite-avatica-go/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type HbaseConnection struct {
	Name string
	URL  string
	DB   *gorm.DB
}

func ConnectHbase() ([]HbaseConnection, error) {

	listUrl := strings.Split(os.Getenv("PQS_URL"), ",")
	log.Println(listUrl)

	var listConn []HbaseConnection
	for _, url := range listUrl {
		url = fmt.Sprintf(
			"%s/%s",
			url,
			CONFIG["SCHEMA_TABLE"],
		)
		log.Print(url)

		hbaseConn, err := gorm.Open("avatica", url)
		if err != nil {
			log.Printf("DB connection error: %s", err.Error())
			continue
		}
		hbaseConn.DB().SetMaxIdleConns(2)
		hbaseConn.DB().SetMaxOpenConns(1000)
		log.Printf("Hbase connnection success: %s", url)

		listConn = append(listConn, HbaseConnection{
			Name: url,
			URL:  url,
			DB:   hbaseConn,
		})
	}

	if len(listConn) == 0 {
		return nil, errors.New("no hbase connection available")
	}

	return listConn, nil
}
