package services

import "fmt"

// NextMotorService is the reference service API for Next Motor.
// Other services (CarInspection, MCOnline, MotorIIA) are modeled after this.
type NextMotorService struct {
	BaseURL string
}

// NewNextMotorService creates a new NextMotorService instance.
func NewNextMotorService(baseURL string) *NextMotorService {
	return &NextMotorService{BaseURL: baseURL}
}

// ServiceName returns the name of the service.
func (s *NextMotorService) ServiceName() string {
	return "NextMotor"
}

// Execute sends the Next Motor service request and returns a response.
func (s *NextMotorService) Execute(req ServiceRequest) (*ServiceResponse, error) {
	if req.PolicyNumber == "" {
		return nil, fmt.Errorf("policy_number is required")
	}
	if req.ProductType == "" {
		return nil, fmt.Errorf("product_type is required")
	}

	return &ServiceResponse{
		Status:  "success",
		Message: fmt.Sprintf("NextMotor service executed for policy %s", req.PolicyNumber),
		RefID:   fmt.Sprintf("NM-%s", req.PolicyNumber),
	}, nil
}
