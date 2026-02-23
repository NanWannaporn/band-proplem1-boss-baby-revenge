package services

import "fmt"

// MCOnlineService handles MC Online requests for Motorcycle (MC).
type MCOnlineService struct {
	BaseURL string
}

// NewMCOnlineService creates a new MCOnlineService instance.
func NewMCOnlineService(baseURL string) *MCOnlineService {
	return &MCOnlineService{BaseURL: baseURL}
}

// ServiceName returns the name of the service.
func (s *MCOnlineService) ServiceName() string {
	return "MCOnline"
}

// Execute sends the MC Online request for Motorcycle (MC) and returns a response.
func (s *MCOnlineService) Execute(req ServiceRequest) (*ServiceResponse, error) {
	if req.PolicyNumber == "" {
		return nil, fmt.Errorf("policy_number is required")
	}
	if req.ProductType != MC {
		return nil, fmt.Errorf("mc_online service only supports product type MC, got %s", req.ProductType)
	}

	return &ServiceResponse{
		Status:  "success",
		Message: fmt.Sprintf("MCOnline service executed for policy %s (MC)", req.PolicyNumber),
		RefID:   fmt.Sprintf("MC-%s", req.PolicyNumber),
	}, nil
}
