package passport

import (
	"time"
)

type Record struct {
	Type      string    `json:type`
	Path      string    `json:path`
	Timestamp time.Time `json:timestamp`
}
