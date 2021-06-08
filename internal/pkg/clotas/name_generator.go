package clotas

import (
	"fmt"
	"time"
)

func GenerateName(scriptName string, t time.Time, n int) string {
	return fmt.Sprintf("%s%s%03d%s%s.%s",
		t.Format(dateLayout()),
		DefaultSeparator,
		n,
		DefaultSeparator,
		scriptName,
		DefaultFileType)
}

func dateLayout() string {
	return DefaultDateLayout
}
