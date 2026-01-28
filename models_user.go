package weheat

import "time"

// ReadUserMe mirrors ReadUserMeDto.
type ReadUserMe struct {
	ID        string    `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Role      Role      `json:"role"`
	Email     *string   `json:"email,omitempty"`
	UpdatedOn time.Time `json:"updatedOn"`
	CreatedOn time.Time `json:"createdOn"`
	Language  *string   `json:"language,omitempty"`
}

// ReadUser mirrors ReadUserDto.
type ReadUser struct {
	ID        string    `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Role      Role      `json:"role"`
	Email     *string   `json:"email,omitempty"`
	UpdatedOn time.Time `json:"updatedOn"`
	CreatedOn time.Time `json:"createdOn"`
}
