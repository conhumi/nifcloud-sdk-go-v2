// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdbiface provides an interface to enable mocking the NIFCLOUD RDB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdbiface

import (
	"github.com/alice02/nifcloud-sdk-go-v2/service/rdb"
)

// RdbAPI provides an interface to enable mocking the
// rdb.Rdb service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // NIFCLOUD RDB.
//    func myFunc(svc rdbiface.RdbAPI) bool {
//        // Make svc.AddSourceIdentifierToSubscription request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := rdb.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRdbClient struct {
//        rdbiface.RdbAPI
//    }
//    func (m *mockRdbClient) AddSourceIdentifierToSubscription(input *rdb.AddSourceIdentifierToSubscriptionInput) (*rdb.AddSourceIdentifierToSubscriptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRdbClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RdbAPI interface {
	AddSourceIdentifierToSubscriptionRequest(*rdb.AddSourceIdentifierToSubscriptionInput) rdb.AddSourceIdentifierToSubscriptionRequest

	AuthorizeDBSecurityGroupIngressRequest(*rdb.AuthorizeDBSecurityGroupIngressInput) rdb.AuthorizeDBSecurityGroupIngressRequest

	CopyDBSnapshotRequest(*rdb.CopyDBSnapshotInput) rdb.CopyDBSnapshotRequest

	CreateDBInstanceRequest(*rdb.CreateDBInstanceInput) rdb.CreateDBInstanceRequest

	CreateDBInstanceReadReplicaRequest(*rdb.CreateDBInstanceReadReplicaInput) rdb.CreateDBInstanceReadReplicaRequest

	CreateDBParameterGroupRequest(*rdb.CreateDBParameterGroupInput) rdb.CreateDBParameterGroupRequest

	CreateDBSecurityGroupRequest(*rdb.CreateDBSecurityGroupInput) rdb.CreateDBSecurityGroupRequest

	CreateDBSnapshotRequest(*rdb.CreateDBSnapshotInput) rdb.CreateDBSnapshotRequest

	CreateEventSubscriptionRequest(*rdb.CreateEventSubscriptionInput) rdb.CreateEventSubscriptionRequest

	DeleteDBInstanceRequest(*rdb.DeleteDBInstanceInput) rdb.DeleteDBInstanceRequest

	DeleteDBParameterGroupRequest(*rdb.DeleteDBParameterGroupInput) rdb.DeleteDBParameterGroupRequest

	DeleteDBSecurityGroupRequest(*rdb.DeleteDBSecurityGroupInput) rdb.DeleteDBSecurityGroupRequest

	DeleteDBSnapshotRequest(*rdb.DeleteDBSnapshotInput) rdb.DeleteDBSnapshotRequest

	DeleteEventSubscriptionRequest(*rdb.DeleteEventSubscriptionInput) rdb.DeleteEventSubscriptionRequest

	DescribeDBEngineVersionsRequest(*rdb.DescribeDBEngineVersionsInput) rdb.DescribeDBEngineVersionsRequest

	DescribeDBInstancesRequest(*rdb.DescribeDBInstancesInput) rdb.DescribeDBInstancesRequest

	DescribeDBLogFilesRequest(*rdb.DescribeDBLogFilesInput) rdb.DescribeDBLogFilesRequest

	DescribeDBParameterGroupsRequest(*rdb.DescribeDBParameterGroupsInput) rdb.DescribeDBParameterGroupsRequest

	DescribeDBParametersRequest(*rdb.DescribeDBParametersInput) rdb.DescribeDBParametersRequest

	DescribeDBSecurityGroupsRequest(*rdb.DescribeDBSecurityGroupsInput) rdb.DescribeDBSecurityGroupsRequest

	DescribeDBSnapshotsRequest(*rdb.DescribeDBSnapshotsInput) rdb.DescribeDBSnapshotsRequest

	DescribeEngineDefaultParametersRequest(*rdb.DescribeEngineDefaultParametersInput) rdb.DescribeEngineDefaultParametersRequest

	DescribeEventCategoriesRequest(*rdb.DescribeEventCategoriesInput) rdb.DescribeEventCategoriesRequest

	DescribeEventSubscriptionsRequest(*rdb.DescribeEventSubscriptionsInput) rdb.DescribeEventSubscriptionsRequest

	DescribeEventsRequest(*rdb.DescribeEventsInput) rdb.DescribeEventsRequest

	DescribeOrderableDBInstanceOptionsRequest(*rdb.DescribeOrderableDBInstanceOptionsInput) rdb.DescribeOrderableDBInstanceOptionsRequest

	DownloadDBLogFilePortionRequest(*rdb.DownloadDBLogFilePortionInput) rdb.DownloadDBLogFilePortionRequest

	ModifyDBInstanceRequest(*rdb.ModifyDBInstanceInput) rdb.ModifyDBInstanceRequest

	ModifyDBParameterGroupRequest(*rdb.ModifyDBParameterGroupInput) rdb.ModifyDBParameterGroupRequest

	ModifyEventSubscriptionRequest(*rdb.ModifyEventSubscriptionInput) rdb.ModifyEventSubscriptionRequest

	NiftyFailoverDBInstanceRequest(*rdb.NiftyFailoverDBInstanceInput) rdb.NiftyFailoverDBInstanceRequest

	NiftyGetMetricStatisticsRequest(*rdb.NiftyGetMetricStatisticsInput) rdb.NiftyGetMetricStatisticsRequest

	RebootDBInstanceRequest(*rdb.RebootDBInstanceInput) rdb.RebootDBInstanceRequest

	RemoveSourceIdentifierFromSubscriptionRequest(*rdb.RemoveSourceIdentifierFromSubscriptionInput) rdb.RemoveSourceIdentifierFromSubscriptionRequest

	ResetDBParameterGroupRequest(*rdb.ResetDBParameterGroupInput) rdb.ResetDBParameterGroupRequest

	RestoreDBInstanceFromDBSnapshotRequest(*rdb.RestoreDBInstanceFromDBSnapshotInput) rdb.RestoreDBInstanceFromDBSnapshotRequest

	RestoreDBInstanceToPointInTimeRequest(*rdb.RestoreDBInstanceToPointInTimeInput) rdb.RestoreDBInstanceToPointInTimeRequest

	RevokeDBSecurityGroupIngressRequest(*rdb.RevokeDBSecurityGroupIngressInput) rdb.RevokeDBSecurityGroupIngressRequest
}

var _ RdbAPI = (*rdb.Rdb)(nil)