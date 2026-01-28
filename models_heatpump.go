package weheat

import "time"

// ReadHeatPump mirrors ReadHeatPumpDto.
type ReadHeatPump struct {
	ControlBoardID string          `json:"controlBoardId"`
	Name           *string         `json:"name,omitempty"`
	Model          *HeatPumpModel  `json:"model,omitempty"`
	DHWType        *DhwType        `json:"dhwType,omitempty"`
	BoilerType     *BoilerType     `json:"boilerType,omitempty"`
	Status         *HeatPumpStatus `json:"status,omitempty"`
	CommissionedAt *time.Time      `json:"commissionedAt,omitempty"`
	SerialNumber   string          `json:"serialNumber"`
	PartNumber     *string         `json:"partNumber,omitempty"`
	State          DeviceState     `json:"state"`
	ID             string          `json:"id"`
}

// ReadAllHeatPump mirrors ReadAllHeatPumpDto.
type ReadAllHeatPump struct {
	ControlBoardID  string          `json:"controlBoardId"`
	Name            *string         `json:"name,omitempty"`
	Model           *HeatPumpModel  `json:"model,omitempty"`
	DHWType         *DhwType        `json:"dhwType,omitempty"`
	BoilerType      *BoilerType     `json:"boilerType,omitempty"`
	Status          *HeatPumpStatus `json:"status,omitempty"`
	CommissionedAt  *time.Time      `json:"commissionedAt,omitempty"`
	SerialNumber    string          `json:"serialNumber"`
	PartNumber      *string         `json:"partNumber,omitempty"`
	State           DeviceState     `json:"state"`
	ID              string          `json:"id"`
	FirmwareVersion *string         `json:"firmwareVersion,omitempty"`
}

// PaginationMetadata describes paging metadata.
type PaginationMetadata struct {
	TotalCount  *int `json:"totalCount,omitempty"`
	PageSize    *int `json:"pageSize,omitempty"`
	CurrentPage *int `json:"currentPage,omitempty"`
	TotalPages  *int `json:"totalPages,omitempty"`
}

// ReadAllHeatPumpPagedResponse mirrors ReadAllHeatPumpDtoPagedResponse.
type ReadAllHeatPumpPagedResponse struct {
	Metadata *PaginationMetadata `json:"metadata,omitempty"`
	Data     []ReadAllHeatPump   `json:"data,omitempty"`
}
