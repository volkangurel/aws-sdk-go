// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLambda_AddPermission() {
	svc := lambda.New(nil)

	params := &lambda.AddPermissionInput{
		Action:        aws.StringPtr("Action"),       // Required
		FunctionName:  aws.StringPtr("FunctionName"), // Required
		Principal:     aws.StringPtr("Principal"),    // Required
		StatementID:   aws.StringPtr("StatementId"),  // Required
		SourceARN:     aws.StringPtr("Arn"),
		SourceAccount: aws.StringPtr("SourceOwner"),
	}
	resp, err := svc.AddPermission(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_CreateEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.CreateEventSourceMappingInput{
		EventSourceARN:   aws.StringPtr("Arn"),                 // Required
		FunctionName:     aws.StringPtr("FunctionName"),        // Required
		StartingPosition: aws.StringPtr("EventSourcePosition"), // Required
		BatchSize:        aws.Int64Ptr(1),
		Enabled:          aws.BoolPtr(true),
	}
	resp, err := svc.CreateEventSourceMapping(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_CreateFunction() {
	svc := lambda.New(nil)

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{ // Required
			S3Bucket:        aws.StringPtr("S3Bucket"),
			S3Key:           aws.StringPtr("S3Key"),
			S3ObjectVersion: aws.StringPtr("S3ObjectVersion"),
			ZipFile:         []byte("PAYLOAD"),
		},
		FunctionName: aws.StringPtr("FunctionName"), // Required
		Handler:      aws.StringPtr("Handler"),      // Required
		Role:         aws.StringPtr("RoleArn"),      // Required
		Runtime:      aws.StringPtr("Runtime"),      // Required
		Description:  aws.StringPtr("Description"),
		MemorySize:   aws.Int64Ptr(1),
		Timeout:      aws.Int64Ptr(1),
	}
	resp, err := svc.CreateFunction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_DeleteEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.DeleteEventSourceMappingInput{
		UUID: aws.StringPtr("String"), // Required
	}
	resp, err := svc.DeleteEventSourceMapping(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_DeleteFunction() {
	svc := lambda.New(nil)

	params := &lambda.DeleteFunctionInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
	}
	resp, err := svc.DeleteFunction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.GetEventSourceMappingInput{
		UUID: aws.StringPtr("String"), // Required
	}
	resp, err := svc.GetEventSourceMapping(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetFunction() {
	svc := lambda.New(nil)

	params := &lambda.GetFunctionInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
	}
	resp, err := svc.GetFunction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetFunctionConfiguration() {
	svc := lambda.New(nil)

	params := &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
	}
	resp, err := svc.GetFunctionConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_GetPolicy() {
	svc := lambda.New(nil)

	params := &lambda.GetPolicyInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
	}
	resp, err := svc.GetPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_Invoke() {
	svc := lambda.New(nil)

	params := &lambda.InvokeInput{
		FunctionName:   aws.StringPtr("FunctionName"), // Required
		ClientContext:  aws.StringPtr("String"),
		InvocationType: aws.StringPtr("InvocationType"),
		LogType:        aws.StringPtr("LogType"),
		Payload:        []byte("PAYLOAD"),
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_InvokeAsync() {
	svc := lambda.New(nil)

	params := &lambda.InvokeAsyncInput{
		FunctionName: aws.StringPtr("FunctionName"),      // Required
		InvokeArgs:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.InvokeAsync(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_ListEventSourceMappings() {
	svc := lambda.New(nil)

	params := &lambda.ListEventSourceMappingsInput{
		EventSourceARN: aws.StringPtr("Arn"),
		FunctionName:   aws.StringPtr("FunctionName"),
		Marker:         aws.StringPtr("String"),
		MaxItems:       aws.Int64Ptr(1),
	}
	resp, err := svc.ListEventSourceMappings(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_ListFunctions() {
	svc := lambda.New(nil)

	params := &lambda.ListFunctionsInput{
		Marker:   aws.StringPtr("String"),
		MaxItems: aws.Int64Ptr(1),
	}
	resp, err := svc.ListFunctions(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_RemovePermission() {
	svc := lambda.New(nil)

	params := &lambda.RemovePermissionInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
		StatementID:  aws.StringPtr("StatementId"),  // Required
	}
	resp, err := svc.RemovePermission(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateEventSourceMapping() {
	svc := lambda.New(nil)

	params := &lambda.UpdateEventSourceMappingInput{
		UUID:         aws.StringPtr("String"), // Required
		BatchSize:    aws.Int64Ptr(1),
		Enabled:      aws.BoolPtr(true),
		FunctionName: aws.StringPtr("FunctionName"),
	}
	resp, err := svc.UpdateEventSourceMapping(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateFunctionCode() {
	svc := lambda.New(nil)

	params := &lambda.UpdateFunctionCodeInput{
		FunctionName:    aws.StringPtr("FunctionName"), // Required
		S3Bucket:        aws.StringPtr("S3Bucket"),
		S3Key:           aws.StringPtr("S3Key"),
		S3ObjectVersion: aws.StringPtr("S3ObjectVersion"),
		ZipFile:         []byte("PAYLOAD"),
	}
	resp, err := svc.UpdateFunctionCode(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleLambda_UpdateFunctionConfiguration() {
	svc := lambda.New(nil)

	params := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.StringPtr("FunctionName"), // Required
		Description:  aws.StringPtr("Description"),
		Handler:      aws.StringPtr("Handler"),
		MemorySize:   aws.Int64Ptr(1),
		Role:         aws.StringPtr("RoleArn"),
		Timeout:      aws.Int64Ptr(1),
	}
	resp, err := svc.UpdateFunctionConfiguration(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
