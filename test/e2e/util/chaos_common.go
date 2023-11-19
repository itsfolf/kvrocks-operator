package util

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	k8sApiClient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ChaosMeshNamespace = "chaos-mesh"
)

type Experiment struct {
	chaosObject k8sApiClient.Object
	name        string
	namespace   string
}

func (env *KubernetesEnv) addChaosExperiment(experiment Experiment) {
	env.ChaosMeshExperiments = append(env.ChaosMeshExperiments, experiment)
}

func (env *KubernetesEnv) CreateExperiment(chaos k8sApiClient.Object) *Experiment {
	fmt.Fprintf(GinkgoWriter, "CreateExperiment name=%s\n", chaos.GetName())
	err := env.Client.Create(context.Background(), chaos)
	Expect(err).NotTo(HaveOccurred())

	// create chaos experiment
	experiment := Experiment{
		chaosObject: chaos,
		name:        chaos.GetName(),
		namespace:   chaos.GetNamespace(),
	}

	env.addChaosExperiment(experiment)

	return &experiment
}
