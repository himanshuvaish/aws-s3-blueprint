package aws

import (
	"fmt"

	"github.com/pkg/errors"
	code2cloudv1deploybktstackawsmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/storagebucket/stack/aws/model"
	pulumiawsprovider "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/automation/provider/aws"
	puluminameoutputaws "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/name/provider/cloud/aws/output"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input   *code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackInput
	AwsTags map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	awsProvider, err := pulumiawsprovider.Get(ctx, s.Input.CredentialsInput.Aws)
	if err != nil {
		return errors.Wrap(err, "failed to setup aws provider")
	}

	if err := addBucket(ctx, s.Input.ResourceInput, s.AwsTags, awsProvider); err != nil {
		return errors.Wrap(err, "failed to add bucket")
	}
	return nil
}

func addBucket(ctx *pulumi.Context, input *code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackResourceInput,
	tags map[string]string, awsProvider *aws.Provider) error {
	addedBucket, err := s3.NewBucket(ctx, input.StorageBucket.Metadata.Name, &s3.BucketArgs{
		Acl:          pulumi.String(s3.CannedAclPrivate), // Set the bucket ACL to private
		Tags:         pulumi.ToStringMap(tags),
		Region:       pulumi.String(input.StorageBucket.Spec.Aws.Region),
		Bucket:       pulumi.String(input.StorageBucket.Metadata.Name),
		ForceDestroy: pulumi.Bool(true),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create bucket resource")
	}
	ctx.Export(getAwsRegionOutputName(), addedBucket.Region)
	if input.StorageBucket.Spec.IsPublic {
		_, err = s3.NewBucketPublicAccessBlock(ctx, fmt.Sprintf("%s-public", input.StorageBucket.Metadata.Name), &s3.BucketPublicAccessBlockArgs{
			Bucket:                addedBucket.Bucket,
			BlockPublicAcls:       pulumi.Bool(false),
			BlockPublicPolicy:     pulumi.Bool(false),
			IgnorePublicAcls:      pulumi.Bool(false),
			RestrictPublicBuckets: pulumi.Bool(false),
		}, pulumi.Parent(addedBucket))
		if err != nil {
			return errors.Wrap(err, "failed to add public access control rule")
		}
	}
	return nil
}

func getAwsRegionOutputName() string {
	return puluminameoutputaws.Name(s3.Bucket{}, "aws-region")
}
