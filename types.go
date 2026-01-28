package weheat

// DeviceState represents the device lifecycle state.
type DeviceState int

const (
	DeviceStateProductionObsolete DeviceState = 0
	DeviceStateInStock            DeviceState = 1
	DeviceStateSold               DeviceState = 2
	DeviceStateActive             DeviceState = 3
	DeviceStateInactive           DeviceState = 4
	DeviceStateBroken             DeviceState = 5
	DeviceStateTest               DeviceState = 6
)

// HeatPumpModel identifies the model line.
type HeatPumpModel int

const (
	HeatPumpModelBlackBirdP80    HeatPumpModel = 0
	HeatPumpModelBlackBirdP60    HeatPumpModel = 1
	HeatPumpModelSparrowP60Brown HeatPumpModel = 2
	HeatPumpModelSparrowP60Green HeatPumpModel = 3
	HeatPumpModelSparrowP60Grey  HeatPumpModel = 4
	HeatPumpModelFlintP40        HeatPumpModel = 5
)

// HeatPumpStatus describes the last logged status of the heat pump.
type HeatPumpStatus int

const (
	HeatPumpStatusOffline       HeatPumpStatus = 1
	HeatPumpStatusStandby       HeatPumpStatus = 40
	HeatPumpStatusHeating       HeatPumpStatus = 70
	HeatPumpStatusDefrost       HeatPumpStatus = 90
	HeatPumpStatusCooling       HeatPumpStatus = 130
	HeatPumpStatusDHW           HeatPumpStatus = 150
	HeatPumpStatusLegionella    HeatPumpStatus = 160
	HeatPumpStatusSelfTest      HeatPumpStatus = 170
	HeatPumpStatusManualControl HeatPumpStatus = 180
)

// DhwType describes domestic hot water availability.
type DhwType int

const (
	DhwTypeUnknown     DhwType = 0
	DhwTypeAvailable   DhwType = 1
	DhwTypeUnavailable DhwType = 2
)

// BoilerType describes boiler type in the installation.
type BoilerType int

const (
	BoilerTypeUnknown   BoilerType = 0
	BoilerTypeNone      BoilerType = 1
	BoilerTypeOnOff     BoilerType = 2
	BoilerTypeOpenTherm BoilerType = 3
)

// Role describes user roles.
type Role int

const (
	RoleAdmin              Role = 0
	RoleSupport            Role = 1
	RoleFactory            Role = 2
	RoleSales              Role = 3
	RoleDataScientist      Role = 4
	RoleProductionObsolete Role = 5
	RoleInstaller          Role = 6
	RoleConsumer           Role = 7
	RoleDistributor        Role = 8
)

// HeatPumpState represents a derived operational state from log data.
type HeatPumpState string

const (
	HeatPumpStateOffline       HeatPumpState = "offline"
	HeatPumpStateStandby       HeatPumpState = "standby"
	HeatPumpStateWaterCheck    HeatPumpState = "water_check"
	HeatPumpStateHeating       HeatPumpState = "heating"
	HeatPumpStateCooling       HeatPumpState = "cooling"
	HeatPumpStateDHW           HeatPumpState = "dhw"
	HeatPumpStateLegionella    HeatPumpState = "legionella_prevention"
	HeatPumpStateDefrosting    HeatPumpState = "defrosting"
	HeatPumpStateSelfTest      HeatPumpState = "self_test"
	HeatPumpStateManualControl HeatPumpState = "manual_control"
)

// HeatPumpModelName returns a human-friendly model label.
func HeatPumpModelName(model HeatPumpModel) string {
	switch model {
	case HeatPumpModelBlackBirdP80:
		return "Blackbird P80 heat pump"
	case HeatPumpModelBlackBirdP60:
		return "Blackbird P60 heat pump"
	case HeatPumpModelSparrowP60Brown, HeatPumpModelSparrowP60Green, HeatPumpModelSparrowP60Grey:
		return "Sparrow P60 heat pump"
	case HeatPumpModelFlintP40:
		return "Flint P40 heat pump"
	default:
		return "Unknown"
	}
}
