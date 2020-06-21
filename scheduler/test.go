package scheduler

import (
	"testing"

	"go.uber.org/zap"

	"github.com/Ahmad-Magdy/k8s-pod-scheduler/k8s"

	"go.uber.org/zap/zaptest"
	"k8s.io/client-go/kubernetes/fake"
)

func newTestScheduler(t *testing.T) Scheduler {
	log := zaptest.NewLogger(t)
	k8s, err := k8s.New(log, fake.NewSimpleClientset())
	if err != nil {
		log.Fatal("Error has occurred while initializing newTestScheduler", zap.Error(err))
	}
	return New(log, k8s)
}
