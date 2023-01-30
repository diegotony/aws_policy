package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPolicyCreation(t *testing.T) {

	t.Parallel()
	groups := []string{"developers"}
	aws_account_id := os.Getenv("AWS_ACCOUNT_ID")
	// arn:aws:iam::123456789:policy/test_policy
	policy := "arn:aws:iam::" + aws_account_id + ":policy/test_policy"
	policies := []string{}
	policies = append(policies, policy)

	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"groups":   groups,
			"policies": policies,
		},
	})

	// arn:aws:iam::237918751012:policy/policy-decrypt-secrets-jwt
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	policy_arn := terraform.Output(t, terraformOptions, "policy_arn")
	assert.Equal(t, policy, policy_arn)
}

// go test -v -run TestInstanceCreation -timeout 30m
