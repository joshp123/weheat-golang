package weheat

import "time"

// RawHeatPumpLog mirrors RawHeatpumpLogAndIsOnlineDto/RawHeatPumpLogDto.
type RawHeatPumpLog struct {
	HeatPumpID                               string    `json:"heatPumpId"`
	Timestamp                                time.Time `json:"timestamp"`
	State                                    *int      `json:"state,omitempty"`
	ControlBridgeStatus                      *int      `json:"controlBridgeStatus,omitempty"`
	ControlBridgeStatusDecodedWaterPump      *bool     `json:"controlBridgeStatusDecodedWaterPump,omitempty"`
	ControlBridgeStatusDecodedGasBoiler      *bool     `json:"controlBridgeStatusDecodedGasBoiler,omitempty"`
	ControlBridgeStatusDecodedElectricHeater *bool     `json:"controlBridgeStatusDecodedElectricHeater,omitempty"`
	ControlBridgeStatusDecodedWaterPump2     *bool     `json:"controlBridgeStatusDecodedWaterPump2,omitempty"`
	T1                                       *float64  `json:"t1,omitempty"`
	T2                                       *float64  `json:"t2,omitempty"`
	TAirIn                                   *float64  `json:"tAirIn,omitempty"`
	TAirOut                                  *float64  `json:"tAirOut,omitempty"`
	TWaterIn                                 *float64  `json:"tWaterIn,omitempty"`
	TWaterOut                                *float64  `json:"tWaterOut,omitempty"`
	TWaterHouseIn                            *float64  `json:"tWaterHouseIn,omitempty"`
	RPM                                      *float64  `json:"rpm,omitempty"`
	OnOffThermostatState                     *int      `json:"onOffThermostatState,omitempty"`
	TRoom                                    *float64  `json:"tRoom,omitempty"`
	TRoomTarget                              *float64  `json:"tRoomTarget,omitempty"`
	TThermostatSetpoint                      *float64  `json:"tThermostatSetpoint,omitempty"`
	OTBoilerFeedTemperature                  *float64  `json:"otBoilerFeedTemperature,omitempty"`
	OTBoilerReturnTemperature                *float64  `json:"otBoilerReturnTemperature,omitempty"`
	CentralHeatingFlow                       *int      `json:"centralHeatingFlow,omitempty"`
	DHWFlow                                  *int      `json:"dhwFlow,omitempty"`
	Interval                                 int       `json:"interval"`
	InputStatus                              *int      `json:"inputStatus,omitempty"`
	CurrentControlMethod                     *int      `json:"currentControlMethod,omitempty"`
	SignalStrength                           *int      `json:"signalStrength,omitempty"`
	RPMLimiter                               *float64  `json:"rpmLimiter,omitempty"`
	RPMLimiterType                           *int      `json:"rpmLimiterType,omitempty"`
	PCompressorIn                            *float64  `json:"pCompressorIn,omitempty"`
	PCompressorOut                           *float64  `json:"pCompressorOut,omitempty"`
	PCompressorInTarget                      *float64  `json:"pCompressorInTarget,omitempty"`
	TCompressorIn                            *float64  `json:"tCompressorIn,omitempty"`
	TCompressorOut                           *float64  `json:"tCompressorOut,omitempty"`
	TCompressorInTransient                   *float64  `json:"tCompressorInTransient,omitempty"`
	TCompressorOutTransient                  *float64  `json:"tCompressorOutTransient,omitempty"`
	DeltaTCompressorInSuperheat              *float64  `json:"deltaTCompressorInSuperheat,omitempty"`
	Fan                                      *float64  `json:"fan,omitempty"`
	FanPower                                 *float64  `json:"fanPower,omitempty"`
	TemperatureErrorIntegral                 *float64  `json:"temperatureErrorIntegral,omitempty"`
	ThermostatStatus                         *int      `json:"thermostatStatus,omitempty"`
	OTBoilerStatus                           *int      `json:"otBoilerStatus,omitempty"`
	CentralHeatingPWMRequestedDutyCycle      *int      `json:"centralHeatingPwmRequestedDutyCycle,omitempty"`
	DHWPWMRequestedDutyCycle                 *int      `json:"dhwPwmRequestedDutyCycle,omitempty"`
	SINR                                     *int      `json:"sinr,omitempty"`
	Error                                    *int      `json:"error,omitempty"`
	ErrorDecodedDtcNone                      *bool     `json:"errorDecodedDtcNone,omitempty"`
	ErrorDecodedDtcContinue                  *bool     `json:"errorDecodedDtcContinue,omitempty"`
	ErrorDecodedDtcCompressorOff             *bool     `json:"errorDecodedDtcCompressorOff,omitempty"`
	ErrorDecodedDtcDefrostForbidden          *bool     `json:"errorDecodedDtcDefrostForbidden,omitempty"`
	ErrorDecodedDtcRequestService            *bool     `json:"errorDecodedDtcRequestService,omitempty"`
	ErrorDecodedDtcUseHeatingCurve           *bool     `json:"errorDecodedDtcUseHeatingCurve,omitempty"`
	ErrorDecodedDtcDHWForbidden              *bool     `json:"errorDecodedDtcDhwForbidden,omitempty"`
	ErrorDecodedDtcError                     *bool     `json:"errorDecodedDtcError,omitempty"`
	ErrorDecodedDtcInactive                  *bool     `json:"errorDecodedDtcInactive,omitempty"`
	ControlBridgeStatusDecodedDHWValve       *bool     `json:"controlBridgeStatusDecodedDhwValve,omitempty"`
	TBoard                                   *float64  `json:"tBoard,omitempty"`
	TInverter                                *float64  `json:"tInverter,omitempty"`
	CompressorPowerLowAccuracy               *float64  `json:"compressorPowerLowAccuracy,omitempty"`
	Valve                                    *float64  `json:"valve,omitempty"`
	InverterInputVoltage                     *float64  `json:"inverterInputVoltage,omitempty"`
	IndoorUnitHeaterTemperature              *float64  `json:"indoorUnitHeaterTemperature,omitempty"`
	IndoorUnitInputCurrent                   *float64  `json:"indoorUnitInputCurrent,omitempty"`
	CoolingStatus                            *int      `json:"coolingStatus,omitempty"`
	CMMassPowerIn                            *int      `json:"cmMassPowerIn,omitempty"`
	CMMassPowerOut                           *int      `json:"cmMassPowerOut,omitempty"`
	DebugVariable1                           *float64  `json:"debugVariable1,omitempty"`
	DebugVariable2                           *float64  `json:"debugVariable2,omitempty"`
	DebugVariable3                           *float64  `json:"debugVariable3,omitempty"`
	DebugVariable4                           *float64  `json:"debugVariable4,omitempty"`
	DebugVariable5                           *float64  `json:"debugVariable5,omitempty"`
	IsOnline                                 *bool     `json:"isOnline,omitempty"`
}

// HeatPumpLogView stores aggregated log output without filtering fields.
type HeatPumpLogView map[string]any
