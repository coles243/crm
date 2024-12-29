package router

type Customer struct {
	ID        int     `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Role      *string `json:"role,omitempty"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Contacted *bool   `json:"contacted,omitempty"`
}
