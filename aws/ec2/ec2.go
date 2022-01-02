package ec2

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// To create an ec2 instance, we need to
// 1. Create a key pair: https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.CreateKeyPair
// 2. Create security group: https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.CreateSecurityGroup

// Notes
// 1. Security groups can only be created one per region, if you wish to launch multi region insances,
// you will need to create a security group in each region.
// 2. Security groups should be created in the same region as the key pair that will be used for the ec2 instance.

var IS_TEST = os.Getenv("GO_ENV") == "test"

type EC2 struct {
	region string
	svc    *ec2.EC2
}

type EC2Config struct {
	Region string
}

type SecurityGroupData struct {
	GroupName   string
	Description string
	VPCId       string
}

type CreateInstanceInputData struct {
	ImageId      string
	InstanceType string
	KeyName      string
	MaxCount     int64
	MinCount     int64
}

// type KeyPairInput struct {
// 	KeyName string
// }

func NewEC2(config EC2Config) *EC2 {
	if IS_TEST {
		return &EC2{}
	}

	awsConfig := aws.NewConfig()
	if config.Region != "" {
		awsConfig.WithRegion(config.Region)
	}

	svc := ec2.New(session.New(), awsConfig)

	return &EC2{
		region: config.Region,
		svc:    svc,
	}
}

// TODO: look for a way to automatically download created private key
func (e *EC2) CreateKeyPair(keyName string) (*ec2.CreateKeyPairOutput, error) {
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	}

	return e.svc.CreateKeyPair(input)
}

func (e *EC2) CreateSecurityGroup(data SecurityGroupData) (*ec2.CreateSecurityGroupOutput, error) {
	input := &ec2.CreateSecurityGroupInput{
		Description: aws.String(data.Description),
		GroupName:   aws.String(data.GroupName),
		VpcId:       aws.String(data.VPCId),
	}

	return e.svc.CreateSecurityGroup(input)
}

// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.ModifySecurityGroupRules
// func (e *EC2) ModifySecurityGroupRules() {
// 	input := &ec2.ModifySecurityGroupRulesInput{

// 	}
// }

// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#example_EC2_RunInstances_shared00
func (e *EC2) CreateInstance(data CreateInstanceInputData) (*ec2.Reservation, error) {
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(data.ImageId),
		InstanceType: aws.String(data.InstanceType),
		KeyName:      aws.String(data.KeyName),
		MaxCount:     aws.Int64(data.MaxCount),
		MinCount:     aws.Int64(data.MinCount),
	}

	return e.svc.RunInstances(input)
}
