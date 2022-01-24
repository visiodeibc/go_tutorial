package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName 는 이름을 가지고 greetings.Hello를 호출하여 정상적인 값을
// 반환하는지 확인합니다.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty 빈 스트링을 input으로 주어 에러 결과가 잘 나오는지 확인합니다.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
