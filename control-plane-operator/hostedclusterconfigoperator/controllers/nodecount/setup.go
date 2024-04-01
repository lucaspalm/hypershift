package nodecount

import (
	"context"
	"time"

	hypershiftclient "github.com/openshift/hypershift/client/clientset/clientset"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/openshift/hypershift/control-plane-operator/hostedclusterconfigoperator/operator"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

const ControllerName = "nodecount"

func Setup(ctx context.Context, opts *operator.HostedClusterConfigOperatorConfig) error {
	hypershiftClient, err := hypershiftclient.NewForConfig(opts.CPCluster.GetConfig())
	if err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(opts.Manager).
		Named(ControllerName).
		For(&corev1.Node{}).
		WithOptions(controller.Options{
			RateLimiter: workqueue.NewItemExponentialFailureRateLimiter(1*time.Second, 10*time.Second),
		}).Complete(&reconciler{
		hcpName:            opts.HCPName,
		hcpNamespace:       opts.Namespace,
		client:             hypershiftClient,
		lister:             opts.CPCluster.GetClient(),
		guestClusterClient: opts.Manager.GetClient(),
	})
}
