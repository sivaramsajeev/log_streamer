package configs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	ConfigPropertiesFile = "CONFIG_PROPERTIES_FILE"
	ConfigTopicName      = "CONFIG_TOPIC_NAME"
	ConfigLogFilePath    = "LOG_FILE_PATH"
)

type IConfigs interface {
	Read() *kafka.ConfigMap
}

func GetConfig() (cfg IConfigs) {
	mode := os.Getenv("KAFKA_CONFIG_MODE")

	switch mode {
	case "", "FILE":
		cfg = NewFileConfigs()
	default:
		panic("Unsupported Config mode")
	}
	return
}

type FileConfigs struct {
	path string
}

func NewFileConfigs() *FileConfigs {
	path := MustReadEnv(ConfigPropertiesFile)
	return &FileConfigs{
		path: path,
	}
}

func (f *FileConfigs) Read() *kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(f.path)
	Must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	err = scanner.Err()
	Must(err)

	cm := kafka.ConfigMap(m)
	return &cm
}

type MessageConfig struct {
	Key        []byte
	Value      []byte
	SourcePath string
}

func NewMessageConfig() *MessageConfig {
	filePath := MustReadEnv(ConfigLogFilePath)
	host, err := os.Hostname()
	Must(err)

	key := host + "_" + path.Base(filePath)

	contents := ReadFile(filePath)

	return &MessageConfig{
		Key:        []byte(key),
		Value:      contents,
		SourcePath: filePath,
	}
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	Must(err)
	defer file.Close()

	content, err := io.ReadAll(file)
	Must(err)

	return content
}

func MustReadEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		Must(fmt.Errorf("%s is unset", key))
	}
	return val
}
