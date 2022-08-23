// Code generated by MockGen. DO NOT EDIT.
// Source: aws.pub.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	gomock "github.com/golang/mock/gomock"
)

// MockAwsClient is a mock of AwsClient interface.
type MockAwsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAwsClientMockRecorder
}

// MockAwsClientMockRecorder is the mock recorder for MockAwsClient.
type MockAwsClientMockRecorder struct {
	mock *MockAwsClient
}

// NewMockAwsClient creates a new mock instance.
func NewMockAwsClient(ctrl *gomock.Controller) *MockAwsClient {
	mock := &MockAwsClient{ctrl: ctrl}
	mock.recorder = &MockAwsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsClient) EXPECT() *MockAwsClientMockRecorder {
	return m.recorder
}

// AddPermission mocks base method.
func (m *MockAwsClient) AddPermission(ctx context.Context, params *sqs.AddPermissionInput, optFns ...func(*sqs.Options)) (*sqs.AddPermissionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPermission", varargs...)
	ret0, _ := ret[0].(*sqs.AddPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPermission indicates an expected call of AddPermission.
func (mr *MockAwsClientMockRecorder) AddPermission(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPermission", reflect.TypeOf((*MockAwsClient)(nil).AddPermission), varargs...)
}

// ChangeMessageVisibility mocks base method.
func (m *MockAwsClient) ChangeMessageVisibility(ctx context.Context, params *sqs.ChangeMessageVisibilityInput, optFns ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMessageVisibility", varargs...)
	ret0, _ := ret[0].(*sqs.ChangeMessageVisibilityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMessageVisibility indicates an expected call of ChangeMessageVisibility.
func (mr *MockAwsClientMockRecorder) ChangeMessageVisibility(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibility", reflect.TypeOf((*MockAwsClient)(nil).ChangeMessageVisibility), varargs...)
}

// ChangeMessageVisibilityBatch mocks base method.
func (m *MockAwsClient) ChangeMessageVisibilityBatch(ctx context.Context, params *sqs.ChangeMessageVisibilityBatchInput, optFns ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMessageVisibilityBatch", varargs...)
	ret0, _ := ret[0].(*sqs.ChangeMessageVisibilityBatchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMessageVisibilityBatch indicates an expected call of ChangeMessageVisibilityBatch.
func (mr *MockAwsClientMockRecorder) ChangeMessageVisibilityBatch(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibilityBatch", reflect.TypeOf((*MockAwsClient)(nil).ChangeMessageVisibilityBatch), varargs...)
}

// CreateQueue mocks base method.
func (m *MockAwsClient) CreateQueue(ctx context.Context, params *sqs.CreateQueueInput, optFns ...func(*sqs.Options)) (*sqs.CreateQueueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateQueue", varargs...)
	ret0, _ := ret[0].(*sqs.CreateQueueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQueue indicates an expected call of CreateQueue.
func (mr *MockAwsClientMockRecorder) CreateQueue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQueue", reflect.TypeOf((*MockAwsClient)(nil).CreateQueue), varargs...)
}

// DeleteMessage mocks base method.
func (m *MockAwsClient) DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMessage", varargs...)
	ret0, _ := ret[0].(*sqs.DeleteMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockAwsClientMockRecorder) DeleteMessage(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockAwsClient)(nil).DeleteMessage), varargs...)
}

// DeleteMessageBatch mocks base method.
func (m *MockAwsClient) DeleteMessageBatch(ctx context.Context, params *sqs.DeleteMessageBatchInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageBatchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMessageBatch", varargs...)
	ret0, _ := ret[0].(*sqs.DeleteMessageBatchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMessageBatch indicates an expected call of DeleteMessageBatch.
func (mr *MockAwsClientMockRecorder) DeleteMessageBatch(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageBatch", reflect.TypeOf((*MockAwsClient)(nil).DeleteMessageBatch), varargs...)
}

// DeleteQueue mocks base method.
func (m *MockAwsClient) DeleteQueue(ctx context.Context, params *sqs.DeleteQueueInput, optFns ...func(*sqs.Options)) (*sqs.DeleteQueueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteQueue", varargs...)
	ret0, _ := ret[0].(*sqs.DeleteQueueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteQueue indicates an expected call of DeleteQueue.
func (mr *MockAwsClientMockRecorder) DeleteQueue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQueue", reflect.TypeOf((*MockAwsClient)(nil).DeleteQueue), varargs...)
}

// GetQueueAttributes mocks base method.
func (m *MockAwsClient) GetQueueAttributes(ctx context.Context, params *sqs.GetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueueAttributes", varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueAttributes indicates an expected call of GetQueueAttributes.
func (mr *MockAwsClientMockRecorder) GetQueueAttributes(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAttributes", reflect.TypeOf((*MockAwsClient)(nil).GetQueueAttributes), varargs...)
}

// GetQueueUrl mocks base method.
func (m *MockAwsClient) GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueueUrl", varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueUrlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueUrl indicates an expected call of GetQueueUrl.
func (mr *MockAwsClientMockRecorder) GetQueueUrl(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueUrl", reflect.TypeOf((*MockAwsClient)(nil).GetQueueUrl), varargs...)
}

// ListDeadLetterSourceQueues mocks base method.
func (m *MockAwsClient) ListDeadLetterSourceQueues(ctx context.Context, params *sqs.ListDeadLetterSourceQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeadLetterSourceQueues", varargs...)
	ret0, _ := ret[0].(*sqs.ListDeadLetterSourceQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeadLetterSourceQueues indicates an expected call of ListDeadLetterSourceQueues.
func (mr *MockAwsClientMockRecorder) ListDeadLetterSourceQueues(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeadLetterSourceQueues", reflect.TypeOf((*MockAwsClient)(nil).ListDeadLetterSourceQueues), varargs...)
}

// ListQueueTags mocks base method.
func (m *MockAwsClient) ListQueueTags(ctx context.Context, params *sqs.ListQueueTagsInput, optFns ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueueTags", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueueTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueueTags indicates an expected call of ListQueueTags.
func (mr *MockAwsClientMockRecorder) ListQueueTags(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueueTags", reflect.TypeOf((*MockAwsClient)(nil).ListQueueTags), varargs...)
}

// ListQueues mocks base method.
func (m *MockAwsClient) ListQueues(ctx context.Context, params *sqs.ListQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueues", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueues indicates an expected call of ListQueues.
func (mr *MockAwsClientMockRecorder) ListQueues(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueues", reflect.TypeOf((*MockAwsClient)(nil).ListQueues), varargs...)
}

// PurgeQueue mocks base method.
func (m *MockAwsClient) PurgeQueue(ctx context.Context, params *sqs.PurgeQueueInput, optFns ...func(*sqs.Options)) (*sqs.PurgeQueueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PurgeQueue", varargs...)
	ret0, _ := ret[0].(*sqs.PurgeQueueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeQueue indicates an expected call of PurgeQueue.
func (mr *MockAwsClientMockRecorder) PurgeQueue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeQueue", reflect.TypeOf((*MockAwsClient)(nil).PurgeQueue), varargs...)
}

// ReceiveMessage mocks base method.
func (m *MockAwsClient) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReceiveMessage", varargs...)
	ret0, _ := ret[0].(*sqs.ReceiveMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockAwsClientMockRecorder) ReceiveMessage(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockAwsClient)(nil).ReceiveMessage), varargs...)
}

// RemovePermission mocks base method.
func (m *MockAwsClient) RemovePermission(ctx context.Context, params *sqs.RemovePermissionInput, optFns ...func(*sqs.Options)) (*sqs.RemovePermissionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemovePermission", varargs...)
	ret0, _ := ret[0].(*sqs.RemovePermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePermission indicates an expected call of RemovePermission.
func (mr *MockAwsClientMockRecorder) RemovePermission(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePermission", reflect.TypeOf((*MockAwsClient)(nil).RemovePermission), varargs...)
}

// SendMessage mocks base method.
func (m *MockAwsClient) SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(*sqs.SendMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockAwsClientMockRecorder) SendMessage(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockAwsClient)(nil).SendMessage), varargs...)
}

// SendMessageBatch mocks base method.
func (m *MockAwsClient) SendMessageBatch(ctx context.Context, params *sqs.SendMessageBatchInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageBatchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageBatch", varargs...)
	ret0, _ := ret[0].(*sqs.SendMessageBatchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageBatch indicates an expected call of SendMessageBatch.
func (mr *MockAwsClientMockRecorder) SendMessageBatch(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageBatch", reflect.TypeOf((*MockAwsClient)(nil).SendMessageBatch), varargs...)
}

// SetQueueAttributes mocks base method.
func (m *MockAwsClient) SetQueueAttributes(ctx context.Context, params *sqs.SetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.SetQueueAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetQueueAttributes", varargs...)
	ret0, _ := ret[0].(*sqs.SetQueueAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetQueueAttributes indicates an expected call of SetQueueAttributes.
func (mr *MockAwsClientMockRecorder) SetQueueAttributes(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetQueueAttributes", reflect.TypeOf((*MockAwsClient)(nil).SetQueueAttributes), varargs...)
}

// TagQueue mocks base method.
func (m *MockAwsClient) TagQueue(ctx context.Context, params *sqs.TagQueueInput, optFns ...func(*sqs.Options)) (*sqs.TagQueueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagQueue", varargs...)
	ret0, _ := ret[0].(*sqs.TagQueueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagQueue indicates an expected call of TagQueue.
func (mr *MockAwsClientMockRecorder) TagQueue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagQueue", reflect.TypeOf((*MockAwsClient)(nil).TagQueue), varargs...)
}

// UntagQueue mocks base method.
func (m *MockAwsClient) UntagQueue(ctx context.Context, params *sqs.UntagQueueInput, optFns ...func(*sqs.Options)) (*sqs.UntagQueueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagQueue", varargs...)
	ret0, _ := ret[0].(*sqs.UntagQueueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagQueue indicates an expected call of UntagQueue.
func (mr *MockAwsClientMockRecorder) UntagQueue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagQueue", reflect.TypeOf((*MockAwsClient)(nil).UntagQueue), varargs...)
}