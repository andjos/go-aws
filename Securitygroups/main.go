package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	ec2svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1")}))

	allGrp := getAllSecGrps(ec2svc)
	//fmt.Println(allGrp)

	//fmt.Println(allGrp.SecurityGroups[2].IpPermissions)
	// Print all Description
	for _, d := range allGrp.SecurityGroups {
		fmt.Printf("Name: %v\n\n%v\n", *d.Description, d.IpPermissions)

	}

}

func getAllSecGrps(ec2svc *ec2.EC2) *ec2.DescribeSecurityGroupsOutput {
	r, err := ec2svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{})
	if err != nil {
		fmt.Println("ops: ", err)
	}
	return r
}
