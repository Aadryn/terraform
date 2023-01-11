package default_testing

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Return_The_Identifier_When_NamingOptions_Is_Empty(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "identifier")
	assert.NotEmptyf(t, output, "identifier should not be empty")
}

func Test_Should_Return_The_Identifier_When_NamingOptions_Identifier_Is_Passed(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"identifier": "my-identifier",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "identifier")
	assert.Equal(t, "my-identifier", output)
}

func Test_Should_Return_The_Identifier_When_NamingOptions_Identifier_Is_Passed_Empty(t *testing.T) {
	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"identifier": "",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "identifier")
	assert.Equal(t, "", output)
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Prefix(t *testing.T) {
	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"prefix": "azerty",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, fmt.Sprintf("azerty-%s", id), output, "identifier should contain the prefix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Suffix(t *testing.T) {
	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"suffix": "qsdfgh",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, fmt.Sprintf("qsdfgh-%s", id), output, "identifier should contain the suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Resource_Name(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, fmt.Sprintf("az-%s", id), output, "name should contain the resource_name")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Suffix_And_Prefix(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"prefix": "azerty",
				"suffix": "qsdfgh",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, output, fmt.Sprintf("azerty-qsdfgh-%s", id), "identifier should contain the prefix and suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Suffix_Resource_Name_And_Prefix(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
				"prefix":        "azerty",
				"suffix":        "qsdfgh",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, output, fmt.Sprintf("azerty-az-qsdfgh-%s", id), "identifier should contain the prefix and suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Separator(t *testing.T) {

	identifier := uuid.NewV4().String()

	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
				"prefix":        "azerty",
				"suffix":        "qsdfgh",
				"separator":     "_",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, output, fmt.Sprintf("azerty_az_qsdfgh_%s", id), "identifier should contain the prefix and suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Lower(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
				"prefix":        "azerty",
				"suffix":        "AZERTY",
				"separator":     "_",
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, output, strings.ToLower(fmt.Sprintf("azerty_az_azerty_%s", id)), "identifier should contain the prefix and suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Upper(t *testing.T) {

	identifier := uuid.NewV4().String()
	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
				"prefix":        "azerty",
				"suffix":        "AZERTY",
				"separator":     "_",
				"lower":         false,
			},
		},
	})
	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.Containsf(t, output, strings.ToUpper(fmt.Sprintf("AZERTY_AZ_AZERTY_%s", id)), "identifier should contain the prefix and suffix")
}

func Test_Should_Return_A_Valid_Name_When_NamingOptions_Is_Passed_With_Length(t *testing.T) {

	identifier := uuid.NewV4().String()

	// Arranges
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		MaxRetries:   10,
		Vars: map[string]interface{}{
			"naming_options": map[string]interface{}{
				"resource_name": "az",
				"prefix":        "azerty",
				"suffix":        "AZERTYYYYYYYYYYYYYYYYY",
				"separator":     "_",
				"lower":         false,
				"length":        24,
			},
		},
	})

	defer deferedDeletion(t, terraformOptions, identifier)
	terraform.Init(t, terraformOptions)
	terraform.WorkspaceSelectOrNew(t, terraformOptions, identifier)

	// Acts
	terraform.Apply(t, terraformOptions)

	// Asserts
	output := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "identifier")
	assert.LessOrEqual(t, len(output), 24, "identifier should contain the prefix and suffix")
	assert.Containsf(t, output, strings.ToUpper(fmt.Sprintf("AZ_%s", id)), "identifier should contain the prefix and suffix")
}

func deferedDeletion(t *testing.T, terraformOptions *terraform.Options, identifier string) {
	terraform.Destroy(t, terraformOptions)
	terraform.WorkspaceDelete(t, terraformOptions, identifier)
}
