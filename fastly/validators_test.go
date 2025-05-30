package fastly

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-cty/cty"

	gofastly "github.com/fastly/go-fastly/v10/fastly"
)

func TestValidateLoggingFormatVersion(t *testing.T) {
	for name, testcase := range map[string]struct {
		value          int
		expectedWarns  int
		expectedErrors int
	}{
		"0": {0, 0, 1},
		"1": {1, 0, 0},
		"2": {2, 0, 0},
		"3": {3, 0, 1},
		"4": {4, 0, 1},
		"5": {5, 0, 1},
	} {
		t.Run(name, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateLoggingFormatVersion()(testcase.value, cty.GetAttrPath("format_version")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateLoggingMessageType(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"classic", 0, 0},
		{"loggly", 0, 0},
		{"logplex", 0, 0},
		{"blank", 0, 0},
		{"CLASSIC", 0, 1},
		{"LOGGLY", 0, 1},
		{"LOGPLEX", 0, 1},
		{"BLANK", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateLoggingMessageType()(testcase.value, cty.GetAttrPath("message_type")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateLoggingPlacement(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"none", 0, 0},
		{"NONE", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateLoggingPlacement()(testcase.value, cty.GetAttrPath("placement")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateLoggingServerSideEncryption(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"AES256", 0, 0},
		{"aws:kms", 0, 0},
		{"aes256", 0, 1},
		{"AWS:KMS", 0, 1},
		{"aws:KMS", 0, 1},
		{"AWS:kms", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateLoggingServerSideEncryption()(testcase.value, cty.GetAttrPath("server_side_encryption")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateDirectorQuorum(t *testing.T) {
	for name, testcase := range map[string]struct {
		value          int
		expectedWarns  int
		expectedErrors int
	}{
		"0":   {0, 0, 0},
		"55":  {55, 0, 0},
		"100": {100, 0, 0},
		"-1":  {-1, 0, 1},
		"101": {101, 0, 1},
		"150": {150, 0, 1},
	} {
		t.Run(name, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateDirectorQuorum()(testcase.value, cty.GetAttrPath("quorum")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateDirectorType(t *testing.T) {
	for name, testcase := range map[string]struct {
		value          int
		expectedWarns  int
		expectedErrors int
	}{
		"0": {0, 0, 1},
		"1": {1, 0, 0},
		"2": {2, 0, 1},
		"3": {3, 0, 0},
		"4": {4, 0, 0},
		"5": {5, 0, 1},
	} {
		t.Run(name, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateDirectorType()(testcase.value, cty.GetAttrPath("type")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateConditionType(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"REQUEST", 0, 0},
		{"RESPONSE", 0, 0},
		{"CACHE", 0, 0},
		{"PREFETCH", 0, 0},
		{"request", 0, 1},
		{"response", 0, 1},
		{"cache", 0, 1},
		{"prefetch", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateConditionType()(testcase.value, cty.GetAttrPath("type")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateHeaderAction(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"set", 0, 0},
		{"append", 0, 0},
		{"delete", 0, 0},
		{"regex", 0, 0},
		{"regex_repeat", 0, 0},
		{"SET", 0, 1},
		{"APPEND", 0, 1},
		{"DELETE", 0, 1},
		{"REGEX", 0, 1},
		{"REGEX_REPEAT", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateHeaderAction()(testcase.value, cty.GetAttrPath("action")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateHeaderType(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"request", 0, 0},
		{"fetch", 0, 0},
		{"cache", 0, 0},
		{"response", 0, 0},
		{"REQUEST", 0, 1},
		{"FETCH", 0, 1},
		{"CACHE", 0, 1},
		{"RESPONSE", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateHeaderType()(testcase.value, cty.GetAttrPath("type")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateSnippetType(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"init", 0, 0},
		{"recv", 0, 0},
		{"hit", 0, 0},
		{"miss", 0, 0},
		{"pass", 0, 0},
		{"hash", 0, 0},
		{"fetch", 0, 0},
		{"error", 0, 0},
		{"deliver", 0, 0},
		{"log", 0, 0},
		{"none", 0, 0},
		{"INIT", 0, 1},
		{"RECV", 0, 1},
		{"HIT", 0, 1},
		{"MISS", 0, 1},
		{"PASS", 0, 1},
		{"FETCH", 0, 1},
		{"HASH", 0, 1},
		{"ERROR", 0, 1},
		{"DELIVER", 0, 1},
		{"LOG", 0, 1},
		{"NONE", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateSnippetType()(testcase.value, cty.GetAttrPath("type")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateDictionaryItemMaxSize(t *testing.T) {
	for name, testcase := range map[string]struct {
		value          map[string]any
		expectedWarns  int
		expectedErrors int
	}{
		"Ten hundred dictionary items":          {createTestDictionaryItems(10), 0, 0},
		"Ten thousand dictionary items":         {createTestDictionaryItems(gofastly.MaximumDictionarySize), 0, 0},
		"Ten thousand and one dictionary items": {createTestDictionaryItems(gofastly.MaximumDictionarySize + 1), 0, 1},
	} {
		t.Run(name, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateDictionaryItems()(testcase.value, cty.GetAttrPath("dictionary_items")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func createTestDictionaryItems(size int) map[string]any {
	dictionaryItems := make(map[string]any)

	for i := 0; i < size; i++ {
		dictionaryItems[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	return dictionaryItems
}

func TestValidateUserRole(t *testing.T) {
	for _, testcase := range []struct {
		value          string
		expectedWarns  int
		expectedErrors int
	}{
		{"user", 0, 0},
		{"billing", 0, 0},
		{"engineer", 0, 0},
		{"superuser", 0, 0},
		{"USER", 0, 1},
		{"BILLING", 0, 1},
		{"ENGINEER", 0, 1},
		{"SUPERUSER", 0, 1},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validateUserRole()(testcase.value, cty.GetAttrPath("role")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidatePEMCertificate(t *testing.T) {
	key, cert, ca, err := generateKeyAndCertWithCA()
	if err != nil {
		t.Fatal(err)
	}

	for _, testcase := range []struct {
		value           string
		expectedPemType string
		expectedWarns   int
		expectedErrors  int
	}{
		{key, "PRIVATE KEY", 0, 0},
		{cert, "CERTIFICATE", 0, 0},
		{ca, "CERTIFICATE", 0, 0},
		{key, "CERTIFICATE", 0, 1},
		{"-----BEGIN CERTIFICATE-----\ncafebabe-----END CERTIFICATE-----\n", "CERTIFICATE", 0, 1},
		{fmt.Sprintf("%s\n%s", cert, ca), "CERTIFICATE", 0, 1},
		{"", "", 0, 1},
	} {
		t.Run(fmt.Sprintf("%s - %s", testcase.expectedPemType, testcase.value), func(t *testing.T) {
			actualWarns, actualErrors := diagToWarnsAndErrs(validatePEMBlock(testcase.expectedPemType)(testcase.value, cty.GetAttrPath("certificate_blob")))
			if len(actualWarns) != testcase.expectedWarns {
				t.Errorf("expected %d warnings, actual %d ", testcase.expectedWarns, len(actualWarns))
			}
			if len(actualErrors) != testcase.expectedErrors {
				t.Errorf("expected %d errors, actual %d ", testcase.expectedErrors, len(actualErrors))
			}
		})
	}
}

func TestValidateMultiplePEMBlocks(t *testing.T) {
	key, cert, ca, err := generateKeyAndCertWithCA()
	if err != nil {
		t.Fatal(err)
	}

	for name, testCase := range map[string]struct {
		value            string
		expectedPemType  string
		expectedWarnings int
		expectedErrors   int
	}{
		"single cert":                  {cert, "CERTIFICATE", 0, 0},
		"empty string":                 {"", "", 0, 1},
		"two blocks same type":         {fmt.Sprintf("%s\n%s", cert, ca), "CERTIFICATE", 0, 0},
		"three blocks different types": {fmt.Sprintf("%s\n%s\n%s", cert, ca, key), "CERTIFICATE", 0, 1},
		"gibberish":                    {"jkljansdfj\nasldfjhadskjfh\nlshakdjf", "", 0, 1},
	} {
		t.Run(name, func(t *testing.T) {
			actualWarnings, actualErrors := diagToWarnsAndErrs(validatePEMBlocks(testCase.expectedPemType)(testCase.value, cty.GetAttrPath("intermediates_blob")))

			if len(actualWarnings) != testCase.expectedWarnings {
				t.Errorf("expected %d warnings, got %d", testCase.expectedWarnings, len(actualWarnings))
			}
			if len(actualErrors) != testCase.expectedErrors {
				t.Errorf("expected %d errors, got %d", testCase.expectedErrors, len(actualErrors))
			}
		})
	}
}
