package fastly

import (
	"testing"
)

func TestNGWAFAlertDatadogIntegrationConfigCompatibility(t *testing.T) {
	resource := resourceFastlyNGWAFAlertDatadogIntegration()
	
	// Test 1: Configuration with explicit site value (existing behavior)
	d1 := resource.TestResourceData()
	d1.Set("key", "123456789")
	d1.Set("site", "us1")
	d1.Set("workspace_id", "workspace_123")
	d1.Set("description", "Some Description")
	
	// Validate that explicit site works
	siteValue1 := d1.Get("site").(string)
	if siteValue1 != "us1" {
		t.Errorf("Expected site value 'us1', got '%s'", siteValue1)
	}
	
	// Test 2: Configuration without site value (should use default)
	d2 := resource.TestResourceData()
	d2.Set("key", "123456789")
	d2.Set("workspace_id", "workspace_123")
	d2.Set("description", "Some Description")
	// Note: site not set, should use default
	
	// The default may not be automatically applied in TestResourceData,
	// but we can verify that the schema allows the field to be optional
	siteSchema := resource.Schema["site"]
	if siteSchema.Required {
		t.Error("site field should be optional to allow default behavior")
	}
	if !siteSchema.Optional {
		t.Error("site field should be optional")
	}
	if siteSchema.Default != "us1" {
		t.Errorf("Expected default value 'us1', got '%v'", siteSchema.Default)
	}
	
	// Test 3: Verify that different valid 3-character sites work
	testSites := []string{"us1", "us3", "us5", "eu1", "ap1"}
	for _, site := range testSites {
		_, errors := siteSchema.ValidateFunc(site, "site")
		if len(errors) > 0 {
			t.Errorf("Site value '%s' should be valid, got errors: %v", site, errors)
		}
	}
}