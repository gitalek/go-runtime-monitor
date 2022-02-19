package testing

import (
	"fmt"
	"strings"
)

func GenerateUpdateGaugePath(name, value string) string {
	return GenerateUpdatePath("gauge", name, value)
}

func GenerateUpdateCounterPath(name, value string) string {
	return GenerateUpdatePath("counter", name, value)
}

func GenerateUpdatePath(kind, name, value string) string {
	return fmt.Sprintf("/update/%s/%s/%s", kind, name, value)
}

func TrimRespBodyString(body string) string {
	return strings.Trim(body, "\n ")
}
