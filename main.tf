resource "aws_iam_policy" "policy" {
  name = var.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action   = var.actions
      Effect   = "Allow"
      Sid      = ""
      Resource = var.resources
      }
    ]
  })

}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = var.role
  policy_arn = aws_iam_policy.policy.arn
}