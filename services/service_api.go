package services

// ProductType represents the type of insurance product.
type ProductType string

const (
	// MV represents Motor Vehicle (Class 1) product type.
	MV ProductType = "MV"
	// MC represents Motorcycle product type.
	MC ProductType = "MC"
)

// ServiceRequest represents a common request structure for service APIs.
type ServiceRequest struct {
	PolicyNumber  string      `json:"policy_number"`
	ProductType   ProductType `json:"product_type"`
	EffectiveDate string      `json:"effective_date"`
	ExpireDate    string      `json:"expire_date"`
}

// ServiceResponse represents a common response structure for service APIs.
type ServiceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	RefID   string `json:"ref_id"`
}

// ServiceAPI defines the interface that all service APIs must implement.
type ServiceAPI interface {
	// Execute sends the service request and returns a response.
	Execute(req ServiceRequest) (*ServiceResponse, error)
	// ServiceName returns the name of the service.
	ServiceName() string
}
