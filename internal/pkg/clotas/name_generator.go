package clotas

import (
	"fmt"
	"time"
)

func GenerateName(name string) string {
	return fmt.Sprintf("%s%s0001%s%s.%s",
		time.Now().Format(dateLayout()),
		DefaultSeparator,
		DefaultSeparator,
		name,
		DefaultFileType)
}

func dateLayout() string {
	return DefaultDateLayout
}
