//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gardenerscheduler

import (
	"context"
	"errors"
	"flag"
	gardenv1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/test-infra/pkg/hostscheduler"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	Name                                         = "gardener"
	CloudProviderAll gardenv1beta1.CloudProvider = "all"
)

var (
	kubeconfigPath    string
	cloudproviderName string
	id                string
)

var Register hostscheduler.Register = func(m hostscheduler.Registrations) {
	if m == nil {
		m = make(hostscheduler.Registrations)
	}
	m[Name] = &hostscheduler.Registration{
		Interface: registerScheduler,
		Flags:     registerFlags,
	}
}

var registerFlags hostscheduler.RegisterFlagsFunc = func(fs *flag.FlagSet) {
	fs.StringVar(&kubeconfigPath, "kubeconfig", "", "Path to the gardener cluster kubeconfigPath")
	fs.StringVar(&cloudproviderName, "cloudprovider", "", "Specify the cloudprovider of the shoot that should be taken from the pool")
	fs.StringVar(&id, "id", "", "Unique id to identify the cluster")
}

var registerScheduler hostscheduler.RegisterInterfaceFromArgsFunc = func(ctx context.Context, logger *logrus.Logger) (hostscheduler.Interface, error) {

	if kubeconfigPath == "" {
		return nil, errors.New("no kubeconfig is specified")
	}
	if cloudproviderName == "" {
		cloudproviderName = string(gardenv1beta1.CloudProviderGCP)
	}

	logger.Debugf("Kubeconfig path: %s", kubeconfigPath)
	logger.Debugf("CloudProvider: %s", cloudproviderName)
	logger.Debugf("ID: %s", id)

	return New(ctx, logger, id, kubeconfigPath, gardenv1beta1.CloudProvider(cloudproviderName))
}

func New(_ context.Context, logger *logrus.Logger, id, kubeconfigPath string, cloudprovider gardenv1beta1.CloudProvider) (*gardenerscheduler, error) {

	k8sClient, err := kubernetes.NewClientFromFile("", kubeconfigPath, client.Options{
		Scheme: kubernetes.GardenScheme,
	})
	if err != nil {
		return nil, err
	}

	namespace, err := getNamespaceOfKubeconfig(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return &gardenerscheduler{
		client:        k8sClient,
		logger:        logger,
		id:            id,
		namespace:     namespace,
		cloudprovider: cloudprovider,
	}, nil
}

var _ hostscheduler.Interface = &gardenerscheduler{}
