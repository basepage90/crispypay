package infra

import (
	"os"
)

// 로컬실행 과 도커실행 모두에 대응하기 위한 장치
func GetHost() string {
	host := os.Getenv("PARTNER_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}
