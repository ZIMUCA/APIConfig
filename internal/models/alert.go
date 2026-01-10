package models

import "time"

type Alert struct {
	UID           string    `json:"uid"`
	ChangedFields []string  `json:"changed_fields"`
	DetectedAt    time.Time `json:"detected_at"`
}
