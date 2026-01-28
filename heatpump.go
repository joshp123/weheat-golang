package weheat

import "context"

// HeatPump provides convenience accessors for heat pump telemetry.
type HeatPump struct {
	client          *Client
	id              string
	lastLog         *RawHeatPumpLog
	energyTotals    *TotalEnergyAggregate
	nominalMaxPower *float64
	model           *HeatPumpModel
}

// NewHeatPump builds a helper for a single heat pump.
func NewHeatPump(client *Client, id string) *HeatPump {
	return &HeatPump{client: client, id: id}
}

// ID returns the heat pump ID.
func (h *HeatPump) ID() string {
	return h.id
}

// Log returns the most recently fetched log entry.
func (h *HeatPump) Log() *RawHeatPumpLog {
	return h.lastLog
}

// EnergyTotals returns the most recently fetched energy totals.
func (h *HeatPump) EnergyTotals() *TotalEnergyAggregate {
	return h.energyTotals
}

// NominalMaxPower returns the nominal max power if known.
func (h *HeatPump) NominalMaxPower() *float64 {
	return h.nominalMaxPower
}

// RefreshStatus pulls both logs and energy totals.
func (h *HeatPump) RefreshStatus(ctx context.Context, opts RequestOptions) error {
	if err := h.RefreshLogs(ctx, opts); err != nil {
		return err
	}
	return h.RefreshEnergy(ctx, opts)
}

// RefreshLogs fetches the latest log entry.
func (h *HeatPump) RefreshLogs(ctx context.Context, opts RequestOptions) error {
	if h.client == nil {
		return ErrClientMissing
	}
	if h.nominalMaxPower == nil {
		_ = h.loadNominalMaxPower(ctx, opts)
	}
	log, err := h.client.GetLatestLog(ctx, h.id, opts)
	if err != nil {
		return err
	}
	h.lastLog = log
	return nil
}

// RefreshEnergy fetches the latest energy totals.
func (h *HeatPump) RefreshEnergy(ctx context.Context, opts RequestOptions) error {
	if h.client == nil {
		return ErrClientMissing
	}
	totals, err := h.client.GetEnergyTotals(ctx, h.id, opts)
	if err != nil {
		return err
	}
	h.energyTotals = totals
	return nil
}

func (h *HeatPump) loadNominalMaxPower(ctx context.Context, opts RequestOptions) error {
	details, err := h.client.GetHeatPump(ctx, h.id, opts)
	if err != nil {
		return err
	}
	if details.Model != nil {
		model := *details.Model
		h.model = &model
		value := nominalMaxPowerForModel(model)
		h.nominalMaxPower = &value
	}
	return nil
}

func nominalMaxPowerForModel(model HeatPumpModel) float64 {
	switch model {
	case HeatPumpModelBlackBirdP60:
		return 5520
	case HeatPumpModelSparrowP60Brown, HeatPumpModelSparrowP60Green, HeatPumpModelSparrowP60Grey:
		return 5520
	case HeatPumpModelFlintP40:
		return 5400
	default:
		return 4500
	}
}

func (h *HeatPump) log() *RawHeatPumpLog {
	if h == nil {
		return nil
	}
	return h.lastLog
}

func (h *HeatPump) energy() *TotalEnergyAggregate {
	if h == nil {
		return nil
	}
	return h.energyTotals
}

func (h *HeatPump) WaterInletTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.TWaterIn
	}
	return nil
}

func (h *HeatPump) WaterOutletTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.TWaterOut
	}
	return nil
}

func (h *HeatPump) WaterHouseInTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.TWaterHouseIn
	}
	return nil
}

func (h *HeatPump) AirInletTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.TAirIn
	}
	return nil
}

func (h *HeatPump) AirOutletTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.TAirOut
	}
	return nil
}

func (h *HeatPump) ThermostatWaterSetpoint() *float64 {
	if log := h.log(); log != nil {
		return log.TThermostatSetpoint
	}
	return nil
}

func (h *HeatPump) ThermostatRoomTemperature() *float64 {
	if log := h.log(); log != nil {
		return validFloat(log.TRoom)
	}
	return nil
}

func (h *HeatPump) ThermostatRoomTemperatureSetpoint() *float64 {
	if log := h.log(); log != nil {
		return validFloat(log.TRoomTarget)
	}
	return nil
}

func (h *HeatPump) ThermostatOnOffState() *int {
	if log := h.log(); log != nil {
		return log.OnOffThermostatState
	}
	return nil
}

func (h *HeatPump) PowerInput() *float64 {
	if log := h.log(); log != nil {
		return intToFloat(log.CMMassPowerIn)
	}
	return nil
}

func (h *HeatPump) PowerOutput() *float64 {
	if log := h.log(); log != nil {
		return intToFloat(log.CMMassPowerOut)
	}
	return nil
}

func (h *HeatPump) DHWTopTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.T1
	}
	return nil
}

func (h *HeatPump) DHWBottomTemperature() *float64 {
	if log := h.log(); log != nil {
		return log.T2
	}
	return nil
}

func (h *HeatPump) COP() *float64 {
	input := h.PowerInput()
	output := h.PowerOutput()
	if input == nil || output == nil {
		return nil
	}
	if *input > 0 {
		value := *output / *input
		return &value
	}
	value := 0.0
	return &value
}

func (h *HeatPump) IndoorUnitWaterPumpState() *bool {
	if log := h.log(); log != nil {
		return log.ControlBridgeStatusDecodedWaterPump
	}
	return nil
}

func (h *HeatPump) IndoorUnitAuxiliaryPumpState() *bool {
	if log := h.log(); log != nil {
		return log.ControlBridgeStatusDecodedWaterPump2
	}
	return nil
}

func (h *HeatPump) IndoorUnitDHWValveOrPumpState() *bool {
	if log := h.log(); log != nil {
		return log.ControlBridgeStatusDecodedDHWValve
	}
	return nil
}

func (h *HeatPump) IndoorUnitGasBoilerState() *bool {
	if log := h.log(); log != nil {
		return log.ControlBridgeStatusDecodedGasBoiler
	}
	return nil
}

func (h *HeatPump) IndoorUnitElectricHeaterState() *bool {
	if log := h.log(); log != nil {
		return log.ControlBridgeStatusDecodedElectricHeater
	}
	return nil
}

func (h *HeatPump) CompressorPercentage() *int {
	if h.nominalMaxPower == nil {
		return nil
	}
	log := h.log()
	if log == nil || log.RPM == nil {
		return nil
	}
	value := int((100.0 / *h.nominalMaxPower) * *log.RPM)
	return &value
}

func (h *HeatPump) CompressorRPM() *float64 {
	if log := h.log(); log != nil {
		return log.RPM
	}
	return nil
}

func (h *HeatPump) HeatPumpState() *HeatPumpState {
	log := h.log()
	if log == nil || log.State == nil {
		return nil
	}
	return ParseHeatPumpState(*log.State)
}

func ParseHeatPumpState(code int) *HeatPumpState {
	var state HeatPumpState
	switch {
	case code == 1:
		state = HeatPumpStateOffline
	case code == 40:
		state = HeatPumpStateStandby
	case code == 70:
		state = HeatPumpStateHeating
	case code >= 130 && code < 140:
		state = HeatPumpStateCooling
	case code == 150:
		state = HeatPumpStateDHW
	case code == 160:
		state = HeatPumpStateLegionella
	case code == 170:
		state = HeatPumpStateSelfTest
	case code == 180:
		state = HeatPumpStateManualControl
	case code >= 200 && code <= 240:
		state = HeatPumpStateDefrosting
	default:
		return nil
	}
	return &state
}

func (h *HeatPump) DHWFlowVolume() *float64 {
	log := h.log()
	if log == nil || log.DHWFlow == nil {
		return nil
	}
	return pwmToVolume(float64(*log.DHWFlow), 2.1)
}

func (h *HeatPump) CentralHeatingFlowVolume() *float64 {
	log := h.log()
	if log == nil || log.CentralHeatingFlow == nil {
		return nil
	}
	return pwmToVolume(float64(*log.CentralHeatingFlow), 2.1)
}

func (h *HeatPump) EnergyInHeating() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEInHeating
	}
	return nil
}

func (h *HeatPump) EnergyInDHW() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEInDHW
	}
	return nil
}

func (h *HeatPump) EnergyInDefrost() *float64 {
	energy := h.energy()
	if energy == nil {
		return nil
	}
	return sumFloat(energy.TotalEInHeatingDefrost, energy.TotalEInDHWDefrost)
}

func (h *HeatPump) EnergyInDefrostDHW() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEInDHWDefrost
	}
	return nil
}

func (h *HeatPump) EnergyInDefrostCH() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEInHeatingDefrost
	}
	return nil
}

func (h *HeatPump) EnergyInCooling() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEInCooling
	}
	return nil
}

func (h *HeatPump) EnergyOutHeating() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEOutHeating
	}
	return nil
}

func (h *HeatPump) EnergyOutDHW() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEOutDHW
	}
	return nil
}

func (h *HeatPump) EnergyOutDefrost() *float64 {
	energy := h.energy()
	if energy == nil {
		return nil
	}
	return sumFloat(energy.TotalEOutDHWDefrost, energy.TotalEOutHeatingDefrost)
}

func (h *HeatPump) EnergyOutDefrostDHW() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEOutDHWDefrost
	}
	return nil
}

func (h *HeatPump) EnergyOutDefrostCH() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEOutHeatingDefrost
	}
	return nil
}

func (h *HeatPump) EnergyOutCooling() *float64 {
	if energy := h.energy(); energy != nil {
		return energy.TotalEOutCooling
	}
	return nil
}

func (h *HeatPump) EnergyTotal() *float64 {
	energy := h.energy()
	if energy == nil {
		return nil
	}
	return sumFloat(
		energy.TotalEInHeating,
		energy.TotalEInDHW,
		energy.TotalEInCooling,
		energy.TotalEInHeatingDefrost,
		energy.TotalEInDHWDefrost,
	)
}

func (h *HeatPump) EnergyOutput() *float64 {
	energy := h.energy()
	if energy == nil {
		return nil
	}
	if energy.TotalEOutHeating == nil || energy.TotalEOutDHW == nil || energy.TotalEOutHeatingDefrost == nil || energy.TotalEOutDHWDefrost == nil || energy.TotalEOutCooling == nil {
		return nil
	}
	value := *energy.TotalEOutHeating + *energy.TotalEOutDHW + (-*energy.TotalEOutHeatingDefrost) + (-*energy.TotalEOutDHWDefrost) + (-*energy.TotalEOutCooling)
	return &value
}

func intToFloat(value *int) *float64 {
	if value == nil {
		return nil
	}
	v := float64(*value)
	return &v
}

func validFloat(value *float64) *float64 {
	if value == nil {
		return nil
	}
	if *value == -1 {
		return nil
	}
	return value
}

func pwmToVolume(pwm float64, max float64) *float64 {
	if pwm < 1 || pwm > 75 {
		return nil
	}
	if pwm <= 5 {
		value := 0.0
		return &value
	}
	value := ((pwm - 5) / 70) * max
	return &value
}

func sumFloat(values ...*float64) *float64 {
	if len(values) == 0 {
		return nil
	}
	var total float64
	for _, value := range values {
		if value == nil {
			return nil
		}
		total += *value
	}
	return &total
}
