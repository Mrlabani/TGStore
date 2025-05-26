// (c) Jisin0
//
// config/numerals.go creates user configured numbers.

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	stringTrue        = "true"
	defaultBatchLimit = 100000
)

var (
	DBChannel           int64   // id of database channel
	Admins              []int64 // admins list
	AutoDelete          int64   // minutes after which batch content is autodeleted
	FsubChannels        []int64 // list of force subscribe channels
	BatchSizeLimit      int64   // maximum messages allowed in a batch
	DisableNotification bool    // messages will be sent without a notification
	DisableAdminLogs    bool    // batch wont be logged if created by admins
	ProtectContent      bool    // disable forwarding content from bot
	AllowPublic         bool    // indicates wether anyone can use the bot
	LogChannel          int64   // logs new message batches for moderation purposes
)

func init() {
	err := godotenv.Load(".env", "config.env") // config.env is deprecated and only for backward compatibiity
	if err == nil {
		fmt.Println("configs loaded from .env file")
	}

	DBChannel = int64Environ("DB_CHANNEL")
	Admins = int64ListEnviron("ADMINS")
	FsubChannels = int64ListEnviron("FSUB")
	AutoDelete = int64Environ("AUTO_DELETE")
	LogChannel = int64Environ("LOG_CHANNEL")
	BatchSizeLimit = int64Environ("BATCH_SIZE_LIMIT", defaultBatchLimit)
	DisableNotification = strings.ToLower(os.Getenv("DISABLE_NOTIFICATION")) == stringTrue
	DisableAdminLogs = strings.ToLower(os.Getenv("DISABLE_ADMIN_LOGS")) == stringTrue
	ProtectContent = strings.ToLower(os.Getenv("PROTECT_CONTENT")) == stringTrue
	AllowPublic = strings.ToLower(os.Getenv("ALLOW_PUBLIC")) == stringTrue || len(Admins) < 1
}

// int64Environ gets a environment variable as an int64.
func int64Environ(envVar string, defaultVal ...int64) int64 {
	s := os.Getenv(envVar)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}

		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("config.int64environ: failed to parse %s: %v\n", envVar, err)
	}

	return i
}

// int64ListEnviron returns a slice of int64 for an environment variable.
func int64ListEnviron(envVar string) (result []int64) {
	s := os.Getenv(envVar)
	if s == "" {
		return result
	}

	for _, str := range strings.Split(s, " ") {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Printf("config.int64listenviron: failed to parse %s: %v\n", envVar, err)
		}

		result = append(result, i)
	}

	return result
}
