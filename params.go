package weheat

import "time"

// RequestOptions sets optional request headers.
type RequestOptions struct {
	XVersion        string
	XBackendVersion string
}

// LogInterval defines log aggregation granularity.
type LogInterval string

const (
	LogIntervalMinute        LogInterval = "Minute"
	LogIntervalFiveMinute    LogInterval = "FiveMinute"
	LogIntervalFifteenMinute LogInterval = "FifteenMinute"
	LogIntervalHour          LogInterval = "Hour"
	LogIntervalDay           LogInterval = "Day"
	LogIntervalWeek          LogInterval = "Week"
	LogIntervalMonth         LogInterval = "Month"
	LogIntervalYear          LogInterval = "Year"
)

// EnergyInterval defines energy aggregation granularity.
type EnergyInterval string

const (
	EnergyIntervalHour  EnergyInterval = "Hour"
	EnergyIntervalDay   EnergyInterval = "Day"
	EnergyIntervalWeek  EnergyInterval = "Week"
	EnergyIntervalMonth EnergyInterval = "Month"
	EnergyIntervalYear  EnergyInterval = "Year"
)

// ListHeatPumpsParams controls the heat pump listing query.
type ListHeatPumpsParams struct {
	Page           *int
	PageSize       *int
	Models         []HeatPumpModel
	OrganisationID string
	Search         string
	State          *DeviceState
	RequestOptions
}

// LogQuery controls heat pump log queries.
type LogQuery struct {
	StartTime *time.Time
	EndTime   *time.Time
	Interval  LogInterval
	RequestOptions
}

// EnergyLogQuery controls energy log queries.
type EnergyLogQuery struct {
	StartTime *time.Time
	EndTime   *time.Time
	Interval  EnergyInterval
	RequestOptions
}
