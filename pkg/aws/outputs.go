package aws

import (
	"context"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	code2cloudv1deploybktstackawsmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/storagebucket/stack/aws/model"
)

func Outputs(ctx context.Context, input *code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackInput) (*code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackOutputs, error) {
	return &code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackOutputs{}, nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackInput) *code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackOutputs{}
	}
	return &code2cloudv1deploybktstackawsmodel.StorageBucketAwsStackOutputs{}
}
