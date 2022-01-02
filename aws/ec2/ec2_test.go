package ec2

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/aws/aws-sdk-go/service/ec2"
// )

// func TestEC2_CreateSecurityGroup(t *testing.T) {
// 	type fields struct {
// 		region string
// 		svc    *ec2.EC2
// 	}
// 	type args struct {
// 		data SecurityGroupData
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *ec2.CreateSecurityGroupOutput
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			e := &EC2{
// 				region: tt.fields.region,
// 				svc:    tt.fields.svc,
// 			}
// 			got, err := e.CreateSecurityGroup(tt.args.data)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("EC2.CreateSecurityGroup() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("EC2.CreateSecurityGroup() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestEC2_CreateKeyPair(t *testing.T) {
// 	type fields struct {
// 		region string
// 		svc    *ec2.EC2
// 	}
// 	type args struct {
// 		keyName string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *ec2.CreateKeyPairOutput
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			e := &EC2{
// 				region: tt.fields.region,
// 				svc:    tt.fields.svc,
// 			}
// 			got, err := e.CreateKeyPair(tt.args.keyName)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("EC2.CreateKeyPair() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("EC2.CreateKeyPair() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
