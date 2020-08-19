package server

import (
	"context"
	"testing"

	job "github.com/ahmagdy/k8s-pod-scheduler/job"

	"github.com/ahmagdy/k8s-pod-scheduler/k8s"

	"go.uber.org/zap/zaptest"

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
					Name: "foobar",
					Cron: "1 * * * * *",
				},
			},
			expectedResult: &jobidl.AddJobResponse{
				Id: "foobar",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			schedulerJob := job.SchedulerJobFromJob(tc.request.Job)
			ctrl := gomock.NewController(t)

			mockedScheduler := scheduler.NewMockScheduler(ctrl)
			k := k8s.NewMockK8S(ctrl)
			k.
				EXPECT().
				CreateCronJob(schedulerJob, gomock.Any()).
				Return(tc.request.Job.Name, tc.expectedError).
				AnyTimes()
			logger := zaptest.NewLogger(t)

			grpcServer := newGRPCServer(logger, mockedScheduler, k)
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
