package server

import (
	"context"
	"testing"

	"go.uber.org/zap/zaptest"

	"github.com/golang/protobuf/ptypes/wrappers"

	jobidl "github.com/ahmagdy/k8s-pod-scheduler/job/idl"

	"github.com/ahmagdy/k8s-pod-scheduler/scheduler"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name           string
		request        *jobidl.AddJobRequest
		expectedResult *jobidl.AddJobResponse
		expectedError  error
	}{
		{
			name: "new job is being registered",
			request: &jobidl.AddJobRequest{
				Job: &jobidl.Job{
					Name: &wrappers.StringValue{Value: "foobar"},
					Cron: &wrappers.StringValue{Value: "1 * * * * *"},
				},
			},
			expectedResult: &jobidl.AddJobResponse{
				Id: "foobar",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockedScheduler := scheduler.NewMockScheduler(ctrl)
			mockedScheduler.
				EXPECT().
				Add(gomock.Any()).
				Return(tc.request.Job.Name.Value, tc.expectedError).
				AnyTimes()

			grpcServer := newGRPCServer(zaptest.NewLogger(t), mockedScheduler)
			res, err := grpcServer.Add(context.Background(), tc.request)

			if tc.expectedError != nil {
				require.Error(t, tc.expectedError)
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, res.Id, tc.expectedResult.Id)
		})
	}
}
