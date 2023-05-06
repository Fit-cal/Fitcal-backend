package config

import (
	"fitcal-backend/domain"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config = ConfigDB{}

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func ConnectDB() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
func (c ConfigDB) GetUserName() (*domain.User, error) {
	var result *domain.User
	result = &domain.User{
		Name:  "sabir",
		Email: "sabirbarahi41@gmail.com",
	}
	return result, nil
}
