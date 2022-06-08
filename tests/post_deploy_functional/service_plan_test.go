package tests

import (
	"os"
	"path"
	"testing"

	"github.com/gruntwork-io/go-commons/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TerraTestSuite struct {
	suite.Suite
	TerraformOptions *terraform.Options
}

// setup to do before any test runs
func (suite *TerraTestSuite) SetupSuite() {
	tmpDir := test_structure.CopyTerraformFolderToTemp(suite.T(), "../..", ".")
	_ = files.CopyFile(path.Join("..", "..", ".tool-versions"), path.Join(tmpDir, ".tool-versions"))
	cwd, err := os.Getwd()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.TerraformOptions = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformDir: tmpDir,
		VarFiles:     []string{path.Join(cwd, "..", "test.tfvars")},
	})
	terraform.InitAndApplyAndIdempotent(suite.T(), suite.TerraformOptions)
}

// TearDownAllSuite has a TearDownSuite method, which will run after all the tests in the suite have been run.
func (suite *TerraTestSuite) TearDownSuite() {
	terraform.Destroy(suite.T(), suite.TerraformOptions)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TerraTestSuite))
}

func (suite *TerraTestSuite) TestServicePlan() {
	// NOTE: "subscriptionID" is overridden by the environment variable "ARM_SUBSCRIPTION_ID". <>
	// subscriptionID := ""

	service_plan_id := terraform.Output(suite.T(), suite.TerraformOptions, "service_plan_id")
	service_plan_kind := terraform.Output(suite.T(), suite.TerraformOptions, "service_plan_kind")
	os_type := terraform.Output(suite.T(), suite.TerraformOptions, "os_type")
	is_reserved := terraform.Output(suite.T(), suite.TerraformOptions, "is_reserved")
	worker_count := terraform.Output(suite.T(), suite.TerraformOptions, "worker_count")
	expected_service_plan_kind := "linux"
	expected_os_type := "Linux"
	expected_is_reserved := "true"
	expected_worker_count := "2"

	suite.NotEmpty(service_plan_id, "Service Plan Id cannot be null")
	suite.Equal(expected_service_plan_kind, service_plan_kind, "Plan Kind should match")
	suite.Equal(expected_os_type, os_type, "OS Type should match")
	suite.Equal(expected_is_reserved, is_reserved, "Is_Reserved should match")
	suite.Equal(expected_worker_count, worker_count, "Worker Count should match")

}
