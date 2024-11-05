package core

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

/*
MustLoadConfig загружает конфигурацию из файла, путь к которому указывается в переменной среды CONFIG_PATH.

Возвращает ошибку в следующих случаях:

- Если переменная окружения CONFIG_PATH не задана или указанный путь не существует.

- Если возникает ошибка при открытии файла конфигурации.

- Если не удается декодировать содержимое файла в структуру Config.

- Если конфигурация не проходит валидацию.

Эта функция предназначена для обеспечения корректной загрузки параметров конфигурации, необходимых для работы приложения.
После успешной загрузки и проверки конфигурации, она возвращает указатель на структуру Config.
*/
func MustLoadConfig() (*Config, error) {
	// Получаем путь к файлу конфигурации
	configPath := getConfigPath()
	if configPath == "" {
		return nil, fmt.Errorf("нет пути к файлу конфига")
	}

	// Открываем файл конфигурации
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл конфигурации: %v", err)
	}
	defer file.Close()

	// Декодируем файл конфигурации в структуру Config
	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("не удалось декодировать файл конфигурации: %v", err)
	}

	// Проверяем валидность конфигурации
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

/*
Validate проверяет, валидные ли значения имеет структура Config.

Возвращает ошибку, если:

- `Env` не является одной из предопределенных констант.
*/
func (cfg *Config) Validate() error {
	switch cfg.Env {
	case Local, Dev, Prod:
		// Значение валидно, ничего не делаем
	default:
		return fmt.Errorf("значение поля Env: %s, не является валидным: [%s, %s, %s]", cfg.Env, Local, Dev, Prod)
	}

	// Если значение валидно, возвращаем nil (нет ошибок)
	return nil
}

/*
getConfigPath возвращает путь к файлу конфигурации.

Сначала он пытается прочитать путь из командной строки с помощью флага "config".
Если флаг не указан, возвращается значение переменной окружения CONFIG_PATH.
Если оба значения отсутствуют, возвращается пустая строка.
*/
func getConfigPath() string {
	var path string

	// Определяем флаг командной строки с именем "config" для указания пути к файлу конфигурации
	flag.StringVar(&path, "config", "", "путь к файлу конфигурации")
	flag.Parse()

	// Если путь все еще пустой, пытаемся получить его из переменной окружения
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
