package services

import "fmt"

// CarInspectionService handles car inspection requests for Class 1 (MV).
type CarInspectionService struct {
	BaseURL string
}

// NewCarInspectionService creates a new CarInspectionService instance.
func NewCarInspectionService(baseURL string) *CarInspectionService {
	return &CarInspectionService{BaseURL: baseURL}
}

// ServiceName returns the name of the service.
func (s *CarInspectionService) ServiceName() string {
	return "CarInspection"
}

// Execute sends the car inspection request for Class 1 (MV) and returns a response.
func (s *CarInspectionService) Execute(req ServiceRequest) (*ServiceResponse, error) {
	if req.PolicyNumber == "" {
		return nil, fmt.Errorf("policy_number is required")
	}
	if req.ProductType != MV {
		return nil, fmt.Errorf("car_inspection service only supports product type MV, got %s", req.ProductType)
	}

	return &ServiceResponse{
		Status:  "success",
		Message: fmt.Sprintf("CarInspection service executed for policy %s (Class 1 MV)", req.PolicyNumber),
		RefID:   fmt.Sprintf("CI-%s", req.PolicyNumber),
	}, nil
}
