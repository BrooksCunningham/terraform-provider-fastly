package fastly

import (
	"testing"
)

func TestNGWAFAlertDatadogIntegrationSchemaDefaults(t *testing.T) {
	resource := resourceFastlyNGWAFAlertDatadogIntegration()
	
	// Check that the site field exists and has the correct schema definition
	siteSchema, exists := resource.Schema["site"]
	if !exists {
		t.Fatal("site field does not exist in schema")
	}
	
	// Check that it's optional (allowing for default behavior)
	if siteSchema.Required {
		t.Error("site field should be optional to allow default value")
	}
	
	if !siteSchema.Optional {
		t.Error("site field should be optional")
	}
	
	// Check that it has the correct default value
	if siteSchema.Default != "us1" {
		t.Errorf("Expected default value 'us1', got '%v'", siteSchema.Default)
	}
	
	// Check that the validation function is still present
	if siteSchema.ValidateFunc == nil {
		t.Error("site field should still have length validation")
	}
	
	// Test the validation function with the default value
	_, errors := siteSchema.ValidateFunc("us1", "site")
	if len(errors) > 0 {
		t.Errorf("Default value 'us1' should pass validation, got errors: %v", errors)
	}
	
	// Test that empty string fails validation (the original problem)
	_, errors = siteSchema.ValidateFunc("", "site")
	if len(errors) == 0 {
		t.Error("Empty string should fail validation")
	}
}