package weheat

import "time"

// TotalEnergyAggregate mirrors TotalEnergyAggregate.
type TotalEnergyAggregate struct {
	HeatPumpID              *string  `json:"heatPumpId,omitempty"`
	TotalEInHeating         *float64 `json:"totalEInHeating,omitempty"`
	TotalEInStandby         *float64 `json:"totalEInStandby,omitempty"`
	TotalEInDHW             *float64 `json:"totalEInDhw,omitempty"`
	TotalEInHeatingDefrost  *float64 `json:"totalEInHeatingDefrost,omitempty"`
	TotalEInDHWDefrost      *float64 `json:"totalEInDhwDefrost,omitempty"`
	TotalEInCooling         *float64 `json:"totalEInCooling,omitempty"`
	TotalEOutHeating        *float64 `json:"totalEOutHeating,omitempty"`
	TotalEOutDHW            *float64 `json:"totalEOutDhw,omitempty"`
	TotalEOutHeatingDefrost *float64 `json:"totalEOutHeatingDefrost,omitempty"`
	TotalEOutDHWDefrost     *float64 `json:"totalEOutDhwDefrost,omitempty"`
	TotalEOutCooling        *float64 `json:"totalEOutCooling,omitempty"`
}

// EnergyView mirrors EnergyViewDto.
type EnergyView struct {
	Interval                       *string    `json:"interval,omitempty"`
	TimeBucket                     *time.Time `json:"timeBucket,omitempty"`
	TotalEInHeating                float64    `json:"totalEInHeating"`
	TotalEInStandby                float64    `json:"totalEInStandby"`
	TotalEInDHW                    float64    `json:"totalEInDhw"`
	TotalEInHeatingDefrost         float64    `json:"totalEInHeatingDefrost"`
	TotalEInDHWDefrost             float64    `json:"totalEInDhwDefrost"`
	TotalEInCooling                float64    `json:"totalEInCooling"`
	TotalEOutHeating               float64    `json:"totalEOutHeating"`
	TotalEOutDHW                   float64    `json:"totalEOutDhw"`
	TotalEOutHeatingDefrost        float64    `json:"totalEOutHeatingDefrost"`
	TotalEOutDHWDefrost            float64    `json:"totalEOutDhwDefrost"`
	TotalEOutCooling               float64    `json:"totalEOutCooling"`
	AveragePowerEInHeating         float64    `json:"averagePowerEInHeating"`
	AveragePowerEInStandby         float64    `json:"averagePowerEInStandby"`
	AveragePowerEInDHW             float64    `json:"averagePowerEInDhw"`
	AveragePowerEInHeatingDefrost  float64    `json:"averagePowerEInHeatingDefrost"`
	AveragePowerEInDHWDefrost      float64    `json:"averagePowerEInDhwDefrost"`
	AveragePowerEInCooling         float64    `json:"averagePowerEInCooling"`
	AveragePowerEOutHeating        float64    `json:"averagePowerEOutHeating"`
	AveragePowerEOutDHW            float64    `json:"averagePowerEOutDhw"`
	AveragePowerEOutHeatingDefrost float64    `json:"averagePowerEOutHeatingDefrost"`
	AveragePowerEOutDHWDefrost     float64    `json:"averagePowerEOutDhwDefrost"`
	AveragePowerEOutCooling        float64    `json:"averagePowerEOutCooling"`
}
