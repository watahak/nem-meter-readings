package models

import (
	"time"

	"github.com/google/uuid"
)

type MeterReading struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	NMI         string    `gorm:"type:varchar(10);not null;unique_index:idx_nmi_timestamp"`
	Timestamp   time.Time `gorm:"type:timestamp;not null;unique_index:idx_nmi_timestamp"`
	Consumption float32   `gorm:"type:numeric;not null"`
}

func (MeterReading) TableName() string {
	return "meter_readings"
}
