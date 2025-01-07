// Паттерн "External Configuration Store" предполагает хранение конфигурации приложения вне самого приложения, в
// централизованном хранилище. Это позволяет динамически изменять конфигурацию без необходимости перезапуска приложения,
// что особенно полезно в распределенных системах и облачных средах.

// Когда использовать External Configuration Store
//Динамическая конфигурация: Когда необходимо изменять конфигурацию приложения без его перезапуска.
//Централизованное управление: Когда требуется централизованное управление конфигурацией для нескольких приложений или сервисов.
//Безопасность: Когда конфиденциальные данные (например, ключи API) должны храниться в защищенном месте.
//Масштабируемость: Когда необходимо обеспечить согласованность конфигурации в масштабируемых системах.

//Плюсы External Configuration Store
//Гибкость: Позволяет изменять конфигурацию без перезапуска приложения.
//Централизованное управление: Обеспечивает единое место для управления конфигурацией нескольких приложений.
//Безопасность: Позволяет хранить конфиденциальные данные в защищенном хранилище.
//Согласованность: Обеспечивает согласованность конфигурации в распределенных системах.

//Минусы External Configuration Store
//Сложность настройки: Требует настройки и управления внешним хранилищем конфигурации.
//Зависимость от сети: Приложение зависит от доступности внешнего хранилища.
//Задержки: Может возникнуть задержка при получении конфигурации из внешнего хранилища.
//Безопасность: Необходимо обеспечить безопасность доступа к хранилищу конфигурации.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Config - структура для хранения конфигурации
type Config struct {
	DatabaseURL string `json:"database_url"`
	APIKey      string `json:"api_key"`
}

// LoadConfigFromFile - загружает конфигурацию из файла
func LoadConfigFromFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// LoadConfigFromURL - загружает конфигурацию из удаленного сервиса
func LoadConfigFromURL(url string) (*Config, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	// Пример загрузки конфигурации из файла
	configFile := "config.json"
	config, err := LoadConfigFromFile(configFile)
	if err != nil {
		fmt.Println("Error loading config from file:", err)
		os.Exit(1)
	}
	fmt.Println("Config loaded from file:", config)

	// Пример загрузки конфигурации из URL
	configURL := "http://example.com/config"
	config, err = LoadConfigFromURL(configURL)
	if err != nil {
		fmt.Println("Error loading config from URL:", err)
		os.Exit(1)
	}
	fmt.Println("Config loaded from URL:", config)
}
