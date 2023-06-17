package op

import (
	"fmt"
)

func StringAny(v any) string {
	switch v := v.(type) {
	case rune:
		return fmt.Sprintf("%q", v)
	case string:
		return fmt.Sprintf("%q", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
