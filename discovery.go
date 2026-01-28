package weheat

import "context"

// HeatPumpInfo contains discovery metadata for a heat pump.
type HeatPumpInfo struct {
	ID           string
	DeviceName   string
	Model        *HeatPumpModel
	ModelName    string
	SerialNumber string
	HasDHW       bool
}

// ReadableName returns a friendly name for the pump.
func (h HeatPumpInfo) ReadableName() string {
	if h.DeviceName != "" {
		return h.DeviceName
	}
	if h.ModelName != "" {
		return h.ModelName
	}
	return h.ID
}

// DiscoverActiveHeatPumps lists active heat pumps available to the account.
func (c *Client) DiscoverActiveHeatPumps(ctx context.Context) ([]HeatPumpInfo, error) {
	page := 1
	pageSize := 1000
	state := DeviceStateActive

	var out []HeatPumpInfo
	for {
		resp, err := c.ListHeatPumps(ctx, ListHeatPumpsParams{
			Page:     &page,
			PageSize: &pageSize,
			State:    &state,
		})
		if err != nil {
			return nil, err
		}
		if resp != nil {
			for _, pump := range resp.Data {
				info := HeatPumpInfo{
					ID:           pump.ID,
					SerialNumber: pump.SerialNumber,
				}
				if pump.Name != nil {
					info.DeviceName = *pump.Name
				}
				if pump.Model != nil {
					model := *pump.Model
					info.Model = &model
					info.ModelName = HeatPumpModelName(model)
				} else {
					info.ModelName = "Unknown"
				}
				if pump.DHWType != nil && *pump.DHWType == DhwTypeAvailable {
					info.HasDHW = true
				}
				out = append(out, info)
			}
		}

		if resp == nil || resp.Metadata == nil || resp.Metadata.TotalPages == nil {
			break
		}
		totalPages := *resp.Metadata.TotalPages
		if page >= totalPages {
			break
		}
		page++
	}

	return out, nil
}
