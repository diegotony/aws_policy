# aws_policy
Custom module for AWS policies

<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_iam_policy_attachment.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy_attachment) | resource |
| [random_pet.this](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_policies"></a> [policies](#input\_policies) | n/a | `list(string)` | n/a | yes |
| <a name="input_groups"></a> [groups](#input\_groups) | groups list | `list(string)` | `[]` | no |
| <a name="input_name"></a> [name](#input\_name) | (optional) describe your variable | `string` | `null` | no |
| <a name="input_roles"></a> [roles](#input\_roles) | roles list | `list(string)` | `[]` | no |
| <a name="input_users"></a> [users](#input\_users) | users list | `list(string)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_policy_arn"></a> [policy\_arn](#output\_policy\_arn) | n/a |
<!-- END_TF_DOCS -->