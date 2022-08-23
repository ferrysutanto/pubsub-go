package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type GcpClient interface {
	Close() error
	CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
	CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error)
	CreateTopicWithConfig(ctx context.Context, topicID string, tc *pubsub.TopicConfig) (*pubsub.Topic, error)
	DetachSubscription(ctx context.Context, sub string) (*pubsub.DetachSubscriptionResult, error)
	Snapshot(id string) *pubsub.Snapshot
	Snapshots(ctx context.Context) *pubsub.SnapshotConfigIterator
	Subscription(id string) *pubsub.Subscription
	SubscriptionInProject(id, projectID string) *pubsub.Subscription
	Subscriptions(ctx context.Context) *pubsub.SubscriptionIterator
	Topic(id string) *pubsub.Topic
	TopicInProject(id, projectID string) *pubsub.Topic
	Topics(ctx context.Context) *pubsub.TopicIterator
}
