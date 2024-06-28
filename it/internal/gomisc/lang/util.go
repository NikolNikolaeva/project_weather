package lang

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func CloseQuietly(c io.Closer) {
	_ = c.Close()
}

func Checksum(data []byte) string {
	result := sha256.Sum256(data)

	return hex.EncodeToString(result[:])
}

func Ref[T any](value T) *T {
	return &value
}

func Zero[T any]() (result T) {
	return
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func First[T any](first T, _ ...any) T {
	return first
}

func Second[T any](_ any, second T, _ ...any) T {
	return second
}

func Third[T any](_ any, _ any, third T, _ ...any) T {
	return third
}

func Ternary[T any](condition bool, first T, second T) T {
	if condition {
		return first
	}

	return second
}

func DoWithLock[T any](lock sync.Locker, callback func() T) T {
	lock.Lock()
	defer lock.Unlock()

	return callback()
}

func ToString(value interface{}) string {
	if value == nil {
		return ""
	}

	if _, ok := value.(string); !ok {
		return fmt.Sprint(value)
	}

	return value.(string)
}

func ToJSON(val interface{}) string {
	return string(First(json.Marshal(val)))
}

func FromJSON[T any](data []byte) (result T, err error) {
	return result, json.Unmarshal(data, &result)
}

func EqualTo[T comparable](this T) func(T) bool {
	return func(that T) bool {
		return this == that
	}
}

func GetEnv(variable string, defaultValue string) string {
	value, exists := os.LookupEnv(variable)
	if !exists {
		value = defaultValue
	}
	return value
}

func TruncateTemplateArgs(template string, args ...any) []any {
	return args[:min(len(args), strings.Count(template, "%"))]
}
