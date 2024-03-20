package utils

import (
	"github.com/spf13/viper"
)

// Config adalah struct yang akan digunakan untuk menyimpan konfigurasi aplikasi.
type Config struct {
	// Tambahkan field konfigurasi sesuai dengan kebutuhan Anda.
	// Misalnya:
	KafkaHost        string
	KafkaPort        int
	KafkaServiceName string
}

// InitConfig inisialisasi konfigurasi menggunakan Viper.
func InitConfig() (*Config, error) {

	// // Set path file konfigurasi untuk Viper
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Baca konfigurasi dari file (opsional)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Buat instance Config dan isi sesuai dengan konfigurasi yang dibaca
	config := &Config{
		KafkaHost:        viper.GetString("KAFKA_HOST"),
		KafkaPort:        viper.GetInt("KAFKA_PORT"),
		KafkaServiceName: viper.GetString("KAFKA_SERVICE"),
	}

	return config, nil
}
