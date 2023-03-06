package client

import (
	"context"

	servicetype "github.com/bnb-chain/greenfield-storage-provider/service/types"
	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bnb-chain/greenfield-storage-provider/pkg/log"
	"github.com/bnb-chain/greenfield-storage-provider/service/stonenode/types"
)

// StoneNodeClient is an stone node gRPC service client wrapper
type StoneNodeClient struct {
	address string
	node    types.StoneNodeServiceClient
	conn    *grpc.ClientConn
}

// NewStoneNodeClient return a StoneNodeClient instance
func NewStoneNodeClient(address string) (*StoneNodeClient, error) {
	conn, err := grpc.DialContext(context.Background(), address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorw("failed to dail syncer", "error", err)
		return nil, err
	}
	client := &StoneNodeClient{
		address: address,
		conn:    conn,
		node:    types.NewStoneNodeServiceClient(conn),
	}
	return client, nil
}

// Close the stone node gPRC connection
func (client *StoneNodeClient) Close() error {
	return client.conn.Close()
}

// ReplicateObject async replicate an object payload to other storage provider and seal object
func (client *StoneNodeClient) ReplicateObject(ctx context.Context, object *storagetypes.ObjectInfo, opts ...grpc.CallOption) error {
	_, err := client.node.ReplicateObject(ctx, &types.ReplicateObjectRequest{ObjectInfo: object}, opts...)
	return err
}

// QueryReplicatingObject query a replicating object payload information by object id
func (client *StoneNodeClient) QueryReplicatingObject(ctx context.Context, objectID uint64) (*servicetype.ReplicateSegmentInfo, error) {
	resp, err := client.node.QueryReplicatingObject(ctx, &types.QueryReplicatingObjectRequest{ObjectId: objectID})
	if err != nil {
		return nil, err
	}
	return resp.GetReplicateSegmentInfo(), err
}