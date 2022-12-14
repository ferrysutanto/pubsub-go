// Code generated by MockGen. DO NOT EDIT.
// Source: gcp.pub.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pubsub "cloud.google.com/go/pubsub"
	gomock "github.com/golang/mock/gomock"
)

// MockGcpClient is a mock of GcpClient interface.
type MockGcpClient struct {
	ctrl     *gomock.Controller
	recorder *MockGcpClientMockRecorder
}

// MockGcpClientMockRecorder is the mock recorder for MockGcpClient.
type MockGcpClientMockRecorder struct {
	mock *MockGcpClient
}

// NewMockGcpClient creates a new mock instance.
func NewMockGcpClient(ctrl *gomock.Controller) *MockGcpClient {
	mock := &MockGcpClient{ctrl: ctrl}
	mock.recorder = &MockGcpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGcpClient) EXPECT() *MockGcpClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGcpClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGcpClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGcpClient)(nil).Close))
}

// CreateSubscription mocks base method.
func (m *MockGcpClient) CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, id, cfg)
	ret0, _ := ret[0].(*pubsub.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockGcpClientMockRecorder) CreateSubscription(ctx, id, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockGcpClient)(nil).CreateSubscription), ctx, id, cfg)
}

// CreateTopic mocks base method.
func (m *MockGcpClient) CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", ctx, topicID)
	ret0, _ := ret[0].(*pubsub.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopic indicates an expected call of CreateTopic.
func (mr *MockGcpClientMockRecorder) CreateTopic(ctx, topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockGcpClient)(nil).CreateTopic), ctx, topicID)
}

// CreateTopicWithConfig mocks base method.
func (m *MockGcpClient) CreateTopicWithConfig(ctx context.Context, topicID string, tc *pubsub.TopicConfig) (*pubsub.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopicWithConfig", ctx, topicID, tc)
	ret0, _ := ret[0].(*pubsub.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopicWithConfig indicates an expected call of CreateTopicWithConfig.
func (mr *MockGcpClientMockRecorder) CreateTopicWithConfig(ctx, topicID, tc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopicWithConfig", reflect.TypeOf((*MockGcpClient)(nil).CreateTopicWithConfig), ctx, topicID, tc)
}

// DetachSubscription mocks base method.
func (m *MockGcpClient) DetachSubscription(ctx context.Context, sub string) (*pubsub.DetachSubscriptionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachSubscription", ctx, sub)
	ret0, _ := ret[0].(*pubsub.DetachSubscriptionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachSubscription indicates an expected call of DetachSubscription.
func (mr *MockGcpClientMockRecorder) DetachSubscription(ctx, sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachSubscription", reflect.TypeOf((*MockGcpClient)(nil).DetachSubscription), ctx, sub)
}

// Snapshot mocks base method.
func (m *MockGcpClient) Snapshot(id string) *pubsub.Snapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", id)
	ret0, _ := ret[0].(*pubsub.Snapshot)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockGcpClientMockRecorder) Snapshot(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockGcpClient)(nil).Snapshot), id)
}

// Snapshots mocks base method.
func (m *MockGcpClient) Snapshots(ctx context.Context) *pubsub.SnapshotConfigIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", ctx)
	ret0, _ := ret[0].(*pubsub.SnapshotConfigIterator)
	return ret0
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockGcpClientMockRecorder) Snapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockGcpClient)(nil).Snapshots), ctx)
}

// Subscription mocks base method.
func (m *MockGcpClient) Subscription(id string) *pubsub.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscription", id)
	ret0, _ := ret[0].(*pubsub.Subscription)
	return ret0
}

// Subscription indicates an expected call of Subscription.
func (mr *MockGcpClientMockRecorder) Subscription(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscription", reflect.TypeOf((*MockGcpClient)(nil).Subscription), id)
}

// SubscriptionInProject mocks base method.
func (m *MockGcpClient) SubscriptionInProject(id, projectID string) *pubsub.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionInProject", id, projectID)
	ret0, _ := ret[0].(*pubsub.Subscription)
	return ret0
}

// SubscriptionInProject indicates an expected call of SubscriptionInProject.
func (mr *MockGcpClientMockRecorder) SubscriptionInProject(id, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionInProject", reflect.TypeOf((*MockGcpClient)(nil).SubscriptionInProject), id, projectID)
}

// Subscriptions mocks base method.
func (m *MockGcpClient) Subscriptions(ctx context.Context) *pubsub.SubscriptionIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscriptions", ctx)
	ret0, _ := ret[0].(*pubsub.SubscriptionIterator)
	return ret0
}

// Subscriptions indicates an expected call of Subscriptions.
func (mr *MockGcpClientMockRecorder) Subscriptions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscriptions", reflect.TypeOf((*MockGcpClient)(nil).Subscriptions), ctx)
}

// Topic mocks base method.
func (m *MockGcpClient) Topic(id string) *pubsub.Topic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic", id)
	ret0, _ := ret[0].(*pubsub.Topic)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockGcpClientMockRecorder) Topic(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockGcpClient)(nil).Topic), id)
}

// TopicInProject mocks base method.
func (m *MockGcpClient) TopicInProject(id, projectID string) *pubsub.Topic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopicInProject", id, projectID)
	ret0, _ := ret[0].(*pubsub.Topic)
	return ret0
}

// TopicInProject indicates an expected call of TopicInProject.
func (mr *MockGcpClientMockRecorder) TopicInProject(id, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopicInProject", reflect.TypeOf((*MockGcpClient)(nil).TopicInProject), id, projectID)
}

// Topics mocks base method.
func (m *MockGcpClient) Topics(ctx context.Context) *pubsub.TopicIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topics", ctx)
	ret0, _ := ret[0].(*pubsub.TopicIterator)
	return ret0
}

// Topics indicates an expected call of Topics.
func (mr *MockGcpClientMockRecorder) Topics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topics", reflect.TypeOf((*MockGcpClient)(nil).Topics), ctx)
}
