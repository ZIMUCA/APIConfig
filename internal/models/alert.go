package models

import "time"

type Alert struct {
	UID           string    `json:"uid"`
	ChangedFields []string  `json:"changed_fields"`
	DetectedAt    time.Time `json:"detected_at"`
	Summary       string    `json:"summary"`

	OldDtStart time.Time `json:"old_DtStart"`
	OldDtEnd   time.Time `json:"old_DtEnd"`
	NewDtStart time.Time `json:"new_DtStart"`
	NewDtEnd   time.Time `json:"new_DtEnd"`

	Description string `json:"description"`
}
