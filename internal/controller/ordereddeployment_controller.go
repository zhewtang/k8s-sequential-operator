/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	appv1 "zhtang.github.io/k8s-sequential-operator/api/v1"
)

// OrderedDeploymentReconciler reconciles a OrderedDeployment object
type OrderedDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const (
	reconcileInterval = 20 * time.Second
)

//+kubebuilder:rbac:groups=app.zhtang.github.io,resources=ordereddeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.zhtang.github.io,resources=ordereddeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.zhtang.github.io,resources=ordereddeployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OrderedDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *OrderedDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	orderedDeployment := &appv1.OrderedDeployment{}
	err := r.Get(ctx, req.NamespacedName, orderedDeployment)
	if err != nil {
		logger.Error(err, "cannot get resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Reconcile logic: Iterate through the list of deployments and create/update them sequentially
	for _, deploymentName := range orderedDeployment.Spec.DeploymentOrder {
		// fetch the deployment resource
		deployment := &appsv1.Deployment{}
		err := r.Get(ctx, client.ObjectKey{Namespace: orderedDeployment.Namespace, Name: deploymentName}, deployment)
		if err != nil {
			logger.Info("info", "deployment", deployment.Name, "namespace", orderedDeployment.Name, "target", deploymentName)
			logger.Error(err, "unable to fetch deployment", "deployment", deploymentName)
			return ctrl.Result{}, err
		}

		// check if deployment state for the given deployment
		deploymentReady := false
		for _, condition := range deployment.Status.Conditions {
			if condition.Type == appsv1.DeploymentAvailable && condition.Status == corev1.ConditionTrue {
				deploymentReady = true
				break
			}
		}

		if !deploymentReady {
			// this deployment is not ready. check back later
			logger.Info("deployment not ready yet", "deployment", deploymentName)
			return ctrl.Result{RequeueAfter: reconcileInterval}, nil
		}

		// Update the Deployment's image if needed
		if deployment.Spec.Template.Spec.Containers[0].Image != orderedDeployment.Spec.ImageName {
			previous := deployment.Spec.Template.Spec.Containers[0].Image
			deployment.Spec.Template.Spec.Containers[0].Image = orderedDeployment.Spec.ImageName
			err := r.Update(ctx, deployment)
			if err != nil {
				logger.Error(err, "unable to update Deployment", "deployment", deploymentName)
				return ctrl.Result{}, err
			}

			logger.Info("Deployment reconciled", "deployment", deploymentName, "previous", previous, "now", orderedDeployment.Spec.ImageName)
		} else {
			logger.Info("Deployment already synced", "deployment", deploymentName)
		}
	}

	return ctrl.Result{RequeueAfter: reconcileInterval}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OrderedDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.OrderedDeployment{}).
		Complete(r)
}
