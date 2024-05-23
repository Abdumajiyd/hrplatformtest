package config

import (
    "log"  
    "os"   

    "github.com/joho/godotenv" // .env fayllarini o'qish uchun
)

// Konfiguratsiya uchun struct (tuzilma)
type Config struct {
    HTTPPort    string 
    DatabaseURL string 
}

// Konfiguratsiyani yuklaydigan funksiya
func Load() *Config {
    // .env faylini yuklash uchun
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found") // Agar .env fayl topilmasa
    }

    // Konfiguratsiyani qaytaradi
    return &Config{
        HTTPPort:    getEnv("HTTP_PORT", "8080"), // HTTP_PORT ozgaruvchisini oladi, agar bo'lmasa 8080 qaytaradi
        DatabaseURL: getDatabaseURL(), // malumotlar bazasi URLini yaratadi
    }
}

// envdagi o'zgaruvchisini oladigan funksiya
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value // agar mavjud bo'lsa, qiymatini qaytaradi
    }
    return defaultValue // yoqsa, standart qiymatni qaytaradi
}

// Malumotlar bazasi URLini yaratadigan funksiya
func getDatabaseURL() string {
    // Malumotlar bazasi uchun kerakli o'zgaruvchilarni oladi yoki standart qiymatlardan foydalanadi
    host := getEnv("DB_HOST", "localhost")
    port := getEnv("DB_PORT", "5432")
    user := getEnv("DB_USER", "user")
    password := getEnv("DB_PASSWORD", "password")
    name := getEnv("DB_NAME", "dbname")

    return "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + name + "?sslmode=disable"
}

























































































// package config

// import (
// 	"fmt"
// 	"log"

// 	"github.com/jmoiron/sqlx"
// 	// "github.com/joho/godotenv"

// 	_ "github.com/lib/pq"
// 	// "github.com/spf13/cast"
// 	"github.com/spf13/viper"
// )

// // Config strukturasi konfiguratsiya parametrlarini saqlash uchun ishlatiladi
// type Config struct {
// 	HTTPPort         string `mapstructure:"HTTP_PORT"`      // HTTP server porti
// 	PostgresHost     string `mapstructure:"DB_HOST"`        // PostgreSQL host (server manzili)
// 	PostgresPort     int    `mapstructure:"DB_PORT"`        // PostgreSQL port
// 	PostgresUser     string `mapstructure:"DB_USER"`        // PostgreSQL foydalanuvchi nomi
// 	PostgresPassword string `mapstructure:"DB_PASSWORD"`    // PostgreSQL parol
// 	PostgresDatabase string `mapstructure:"DB_NAME"`        // PostgreSQL ma'lumotlar bazasi nomi
// 	DefaultOffset    int    `mapstructure:"DEFAULT_OFFSET"` // Sahifalash uchun boshlang'ich indeks (default: 0)
// 	DefaultLimit     int    `mapstructure:"DEFAULT_LIMIT"`  // Sahifalash uchun limit (default: 10)
// }

// // Load funksiyasi konfiguratsiyani yuklaydi
// func Load() Config {
// 	viper.AddConfigPath(".")    // .env faylini qidirish uchun "config" papkasini qo'shamiz
// 	viper.SetConfigName(".env") // Konfiguratsiya faylining nomi
// 	viper.SetConfigType("env")  // Konfiguratsiya faylining turi

// 	// .env faylini o'qish
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf(".env faylini o'qishda xatolik: %v", err)
// 	}

// 	// Konfiguratsiyani Config strukturasiga o'tkazish 
// 	var cfg Config
// 	if err := viper.Unmarshal(&cfg); err != nil {
// 		log.Fatalf(".env faylidan ma'lumotlarni structga joylashda xatolik: %v", err)
// 	}

// 	return cfg // Yuklangan konfiguratsiyani qaytarish

// 	// bu ikkinchi usuul
// 	// if err := godotenv.Load(); err != nil {
// 	// 	log.Fatal("Error loading .env file")
// 	// }
// 	// config := Config{
// 	// 	HTTPPort:         cast.ToString(getOrReturnDefaultValue("HTTP_PORT", 8080)),
// 	// 	PostgresHost:     cast.ToString(getOrReturnDefaultValue("DB_HOST", "localhost")),
// 	// 	PostgresPort:     cast.ToInt(getOrReturnDefaultValue("DB_PORT", 5432)),
// 	// 	PostgresUser:     cast.ToString(getOrReturnDefaultValue("DB_USER", "n10")),
// 	// 	PostgresPassword: cast.ToString(getOrReturnDefaultValue("DB_PASSWORD", "1234")),
// 	// 	PostgresDatabase: cast.ToString(getOrReturnDefaultValue("DB_NAME", "n10")),
// 	// 	DefaultOffset:    cast.ToInt(getOrReturnDefaultValue("DEFAULT_OFFSET", 0)),
// 	// 	DefaultLimit:     cast.ToInt(getOrReturnDefaultValue("DEFAULT_LIMIT", 10)),
// 	// }
// 	// return config

// }

// // func getOrReturnDefaultValue(key string, defaultValue any) any {
// // 	value, ok := os.LookupEnv(key)
// // 	if ok {
// // 		return value
// // 	}
// // 	return fmt.Sprintf("%v", defaultValue)
// // }

// // ConnectDB funksiyasi PostgreSQL ma'lumotlar bazasiga ulanishni ochadi
// func (cfg *Config) ConnectDB() (*sqlx.DB, error) {
// 	// Ma'lumotlar bazasiga ulanish uchun DSN (Data Source Name) yaratish
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

// 	// Ma'lumotlar bazasiga ulanish
// 	db, err := sqlx.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("ma'lumotlar bazasiga ulanishda xatolik: %w", err) // Xatolikni qaytarish
// 	}

// 	// Ulanishni tekshirish (ping)
// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("ma'lumotlar bazasiga ulanishni tekshirishda xatolik: %w", err)
// 	}

// 	log.Println("Ma'lumotlar bazasiga muvaffaqiyatli ulanildi")
// 	return db, nil // Ulanish obyektini qaytarish
// }

