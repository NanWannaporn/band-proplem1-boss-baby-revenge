package services

import "testing"

func TestCarInspectionService_Execute_Success(t *testing.T) {
	svc := NewCarInspectionService("http://localhost:8080")
	req := ServiceRequest{
		PolicyNumber:  "POL-001",
		ProductType:   MV,
		EffectiveDate: "2026-01-01",
		ExpireDate:    "2027-01-01",
	}

	resp, err := svc.Execute(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}
	if resp.RefID != "CI-POL-001" {
		t.Errorf("expected RefID 'CI-POL-001', got '%s'", resp.RefID)
	}
}

func TestCarInspectionService_Execute_MissingPolicy(t *testing.T) {
	svc := NewCarInspectionService("http://localhost:8080")
	req := ServiceRequest{ProductType: MV}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for missing policy number")
	}
}

func TestCarInspectionService_Execute_WrongProductType(t *testing.T) {
	svc := NewCarInspectionService("http://localhost:8080")
	req := ServiceRequest{PolicyNumber: "POL-001", ProductType: MC}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for wrong product type")
	}
}

func TestCarInspectionService_ServiceName(t *testing.T) {
	svc := NewCarInspectionService("http://localhost:8080")
	if svc.ServiceName() != "CarInspection" {
		t.Errorf("expected 'CarInspection', got '%s'", svc.ServiceName())
	}
}

func TestMCOnlineService_Execute_Success(t *testing.T) {
	svc := NewMCOnlineService("http://localhost:8080")
	req := ServiceRequest{
		PolicyNumber:  "POL-002",
		ProductType:   MC,
		EffectiveDate: "2026-01-01",
		ExpireDate:    "2027-01-01",
	}

	resp, err := svc.Execute(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}
	if resp.RefID != "MC-POL-002" {
		t.Errorf("expected RefID 'MC-POL-002', got '%s'", resp.RefID)
	}
}

func TestMCOnlineService_Execute_MissingPolicy(t *testing.T) {
	svc := NewMCOnlineService("http://localhost:8080")
	req := ServiceRequest{ProductType: MC}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for missing policy number")
	}
}

func TestMCOnlineService_Execute_WrongProductType(t *testing.T) {
	svc := NewMCOnlineService("http://localhost:8080")
	req := ServiceRequest{PolicyNumber: "POL-002", ProductType: MV}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for wrong product type")
	}
}

func TestMCOnlineService_ServiceName(t *testing.T) {
	svc := NewMCOnlineService("http://localhost:8080")
	if svc.ServiceName() != "MCOnline" {
		t.Errorf("expected 'MCOnline', got '%s'", svc.ServiceName())
	}
}

func TestMotorIIAService_Execute_Success(t *testing.T) {
	svc := NewMotorIIAService("http://localhost:8080")
	req := ServiceRequest{
		PolicyNumber:  "POL-003",
		ProductType:   MV,
		EffectiveDate: "2026-01-01",
		ExpireDate:    "2027-01-01",
	}

	resp, err := svc.Execute(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}
	if resp.RefID != "IIA-POL-003" {
		t.Errorf("expected RefID 'IIA-POL-003', got '%s'", resp.RefID)
	}
}

func TestMotorIIAService_Execute_MissingPolicy(t *testing.T) {
	svc := NewMotorIIAService("http://localhost:8080")
	req := ServiceRequest{ProductType: MV}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for missing policy number")
	}
}

func TestMotorIIAService_Execute_WrongProductType(t *testing.T) {
	svc := NewMotorIIAService("http://localhost:8080")
	req := ServiceRequest{PolicyNumber: "POL-003", ProductType: MC}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for wrong product type")
	}
}

func TestMotorIIAService_ServiceName(t *testing.T) {
	svc := NewMotorIIAService("http://localhost:8080")
	if svc.ServiceName() != "MotorIIA" {
		t.Errorf("expected 'MotorIIA', got '%s'", svc.ServiceName())
	}
}

func TestNextMotorService_Execute_Success(t *testing.T) {
	svc := NewNextMotorService("http://localhost:8080")
	req := ServiceRequest{
		PolicyNumber:  "POL-100",
		ProductType:   MV,
		EffectiveDate: "2026-01-01",
		ExpireDate:    "2027-01-01",
	}

	resp, err := svc.Execute(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}
	if resp.RefID != "NM-POL-100" {
		t.Errorf("expected RefID 'NM-POL-100', got '%s'", resp.RefID)
	}
}

func TestNextMotorService_Execute_MissingPolicy(t *testing.T) {
	svc := NewNextMotorService("http://localhost:8080")
	req := ServiceRequest{ProductType: MV}

	_, err := svc.Execute(req)
	if err == nil {
		t.Fatal("expected error for missing policy number")
	}
}

func TestServiceAPI_Interface(t *testing.T) {
	// Verify all services implement ServiceAPI interface
	var _ ServiceAPI = &NextMotorService{}
	var _ ServiceAPI = &CarInspectionService{}
	var _ ServiceAPI = &MCOnlineService{}
	var _ ServiceAPI = &MotorIIAService{}
}
