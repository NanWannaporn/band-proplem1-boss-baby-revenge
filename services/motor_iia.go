package services

import "fmt"

// MotorIIAService handles Motor IIA generation requests for Motor Vehicle (MV).
type MotorIIAService struct {
	BaseURL string
}

// NewMotorIIAService creates a new MotorIIAService instance.
func NewMotorIIAService(baseURL string) *MotorIIAService {
	return &MotorIIAService{BaseURL: baseURL}
}

// ServiceName returns the name of the service.
func (s *MotorIIAService) ServiceName() string {
	return "MotorIIA"
}

// Execute sends the Motor IIA generation request for MV and returns a response.
func (s *MotorIIAService) Execute(req ServiceRequest) (*ServiceResponse, error) {
	if req.PolicyNumber == "" {
		return nil, fmt.Errorf("policy_number is required")
	}
	if req.ProductType != MV {
		return nil, fmt.Errorf("motor_iia service only supports product type MV, got %s", req.ProductType)
	}

	return &ServiceResponse{
		Status:  "success",
		Message: fmt.Sprintf("MotorIIA service executed for policy %s (MV)", req.PolicyNumber),
		RefID:   fmt.Sprintf("IIA-%s", req.PolicyNumber),
	}, nil
}
