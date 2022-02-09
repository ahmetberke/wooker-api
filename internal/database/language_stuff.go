package database

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func ImplementLanguages(name string, database *gorm.DB)  {

	languagesFile, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = languagesFile.Close()
	}()

	var languages []*models.Language

	byteValue, err := ioutil.ReadAll(languagesFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &languages)
	if err != nil {
		panic(err)
	}

	for _, language := range languages{
		err := database.Create(&language).Error
		if err != nil {
			log.Printf("error on %v - error : %v", language.Name, err.Error())
		}
	}

}