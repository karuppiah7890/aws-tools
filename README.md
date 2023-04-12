# aws-tools

## allow-ip

To use this action, you could create a IAM user with just enough access to do this action

In my case I created an IAM user and attached policy directly with the following actions (permissions)
- `ec2:AuthorizeSecurityGroupIngress`
- `ec2:CreateTags` - This was shown as a dependent action in warning when adding `ec2:AuthorizeSecurityGroupIngress`

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateTags"
            ],
            "Resource": "arn:aws:ec2:us-east-1:123456789012:security-group/sg-11223344556677889"
        }
    ]
}
```
