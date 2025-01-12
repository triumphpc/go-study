package main

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Базовый пример
type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

// Пример с тегами и вложенными структурами
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Auth   AuthConfig   `mapstructure:"auth"`
}

type ServerConfig struct {
	Port     int      `mapstructure:"port"`
	Host     string   `mapstructure:"host"`
	Protocol string   `mapstructure:"protocol"`
	Tags     []string `mapstructure:"tags"`
}

type AuthConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func main() {
	// Пример 1: Базовое использование
	input := map[string]interface{}{
		"name":   "John",
		"age":    30,
		"emails": []string{"john@example.com", "john@gmail.com"},
		"extra": map[string]string{
			"country": "USA",
			"city":    "New York",
		},
	}

	var person Person
	err := mapstructure.Decode(input, &person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Person: %+v\n", person)

	// Пример 2: Конфигурация с вложенными структурами
	configInput := map[string]interface{}{
		"server": map[string]interface{}{
			"port":     8080,
			"host":     "localhost",
			"protocol": "https",
			"tags":     []string{"production", "v1"},
		},
		"auth": map[string]interface{}{
			"username": "admin",
			"password": "secret",
		},
	}

	var config Config
	err = mapstructure.Decode(configInput, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", config)

	// Пример 3: Использование DecoderConfig для дополнительных настроек
	var result Person
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true, // Разрешить слабую типизацию
		Result:           &result,
		TagName:          "mapstructure", // Использовать другой тег
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		panic(err)
	}

	input2 := map[string]interface{}{
		"name":   "Jane",
		"age":    "25", // Строка вместо числа - сработает благодаря WeaklyTypedInput
		"emails": []interface{}{"jane@example.com"},
	}

	err = decoder.Decode(input2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result with WeaklyTypedInput: %+v\n", result)

	// Пример 4: Обработка пропущенных полей
	type PartialPerson struct {
		Name string
		Age  int
	}

	input3 := map[string]interface{}{
		"name": "Bob",
		// Age пропущен
	}

	var partial PartialPerson
	err = mapstructure.Decode(input3, &partial)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Partial: %+v\n", partial)
}

// Пример 5: Пользовательский декодер
type Duration struct {
	time.Duration
}

func (d *Duration) DecodeMapstructure(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("expected string, got %T", data)
	}

	duration, err := time.ParseDuration(str)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}

type ServiceConfig struct {
	Timeout Duration `mapstructure:"timeout"`
}
