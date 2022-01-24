package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("empty name")
	}
	//테스트 성공 케이스
	message := fmt.Sprintf(randomFormat(), name)
	// 테스트 실패 케이스
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos 함수는 여러명에 대한 인사를 반환합니다.
func Hellos(names []string) (map[string]string, error) {
	// map으로 n명의 이름과 n개의 인사말을 메치 시켜줍니다.
	// map의 정의 : map[KeyType]ValueType
	messages := make(map[string]string)
	//for loop을 사용하요 이름 개수 만큼 인사말을 배정해 줍니다.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 룹을 돌아 결정된 값을 map에서 설정해 줍니다.
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
