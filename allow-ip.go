package main

import (
	"context"
	"fmt"
	"log"

	awsconf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/karuppiah7890/aws-tools/config"
	"github.com/urfave/cli/v2"
)

func NewAllowIpCommand() *cli.Command {
	command := cli.Command{
		Name:      "allow-ip",
		Usage:     "allow IP in inbound rules in a security group",
		UsageText: "aws-tools allow-ip <security-group-id> <rule-name> <ip-v4-address>",
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() != 3 {
				return fmt.Errorf("expected exactly 3 arguments, got %d", cCtx.Args().Len())
			}

			securityGroupId := cCtx.Args().Get(0)
			ruleName := cCtx.Args().Get(1)
			ipV4Address := cCtx.Args().Get(2)
			log.Println(securityGroupId, ruleName, ipV4Address)
			return allowIp(securityGroupId, ruleName, ipV4Address)
		},
	}

	return &command
}

func allowIp(securityGroupId string, ruleName string, ipV4Address string) error {

	// Create a new ingress inbound rule?
	// Or Update the existing ingress inbound rule if one exists? How? List all security group's inbound rules and check
	// the rule description and match it with rule name. If none exists, create a security group inbound rule with the rule name

	_, err := config.NewConfigFromEnvVars()
	if err != nil {
		return fmt.Errorf("error occurred while getting configuration from environment variables: %v", err)
	}

	awsconfig, err := awsconf.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("error occurred while loading aws configuration: %v", err)
	}

	ec2Client := ec2.NewFromConfig(awsconfig)

	var sshPort int32 = 22
	tcpProtocol := "tcp"
	cidrIpV4 := fmt.Sprintf("%s/32", ipV4Address)

	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: &securityGroupId,
		IpPermissions: []types.IpPermission{
			{
				IpProtocol: &tcpProtocol,
				FromPort:   &sshPort,
				ToPort:     &sshPort,
				IpRanges: []types.IpRange{
					{
						CidrIp:      &cidrIpV4,
						Description: &ruleName,
					},
				},
			},
		},
	}

	output, err := ec2Client.AuthorizeSecurityGroupIngress(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("error while authorizing security group ingress: %v", err)
	}

	if len(output.SecurityGroupRules) != 1 {
		return fmt.Errorf("expected to add one security group rule but %d were added", len(output.SecurityGroupRules))
	}

	securityGroupRule := output.SecurityGroupRules[0]

	gotCidrIpV4 := *securityGroupRule.CidrIpv4

	if gotCidrIpV4 != cidrIpV4 {
		return fmt.Errorf("expected '%s' CIDR IP V4 but got '%s'", cidrIpV4, gotCidrIpV4)
	}

	if *securityGroupRule.IsEgress {
		return fmt.Errorf("expected to add ingress / inbound rule but an outbound rule was added")
	}

	gotDescription := *securityGroupRule.Description

	if gotDescription != ruleName {
		return fmt.Errorf("expected '%s' rule description / name but got '%s'", ruleName, gotDescription)
	}

	gotFromPort := *securityGroupRule.FromPort

	if gotFromPort != sshPort {
		return fmt.Errorf("expected the from-port to be %d but got %d", sshPort, gotFromPort)
	}

	gotToPort := *securityGroupRule.ToPort

	if gotToPort != sshPort {
		return fmt.Errorf("expected the to-port to be %d but got %d", sshPort, gotToPort)
	}

	gotProtocol := *securityGroupRule.IpProtocol

	if gotProtocol != tcpProtocol {
		return fmt.Errorf("expected the protocol to be '%s' but got '%s'", tcpProtocol, gotProtocol)
	}

	return nil
}

// TODO: Checkout other related methods and use it in other features
// ec2Client.DescribeSecurityGroupRules()
// ec2Client.DescribeSecurityGroups()
// ec2Client.UpdateSecurityGroupRuleDescriptionsIngress()
// ec2Client.ModifySecurityGroupRules()
// ec2Client.RevokeSecurityGroupIngress()
