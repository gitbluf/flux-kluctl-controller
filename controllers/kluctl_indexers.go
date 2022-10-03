/*
Copyright 2020 The Flux authors

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

package controllers

import (
	"context"
	"fmt"
	kluctlv1 "github.com/kluctl/flux-kluctl-controller/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
)

func (r *KluctlDeploymentReconciler) requestsForRevisionChangeOf(indexKey string) func(obj client.Object) []reconcile.Request {
	return func(obj client.Object) []reconcile.Request {
		repo, ok := obj.(interface {
			GetArtifact() *sourcev1.Artifact
		})
		if !ok {
			panic(fmt.Sprintf("Expected an object conformed with GetArtifact() method, but got a %T", obj))
		}
		// If we do not have an artifact, we have no requests to make
		if repo.GetArtifact() == nil {
			return nil
		}

		ctx := context.Background()
		list := &kluctlv1.KluctlDeploymentList{}

		if err := r.List(ctx, list, client.MatchingFields{
			indexKey: client.ObjectKeyFromObject(obj).String(),
		}); err != nil {
			return nil
		}
		var reqs []reconcile.Request
		for _, d := range list.Items {
			// If the revision of the artifact equals to the last attempted revision,
			// we should not make a request for this Kustomization
			if repo.GetArtifact().Revision == d.Status.LastAttemptedRevision {
				continue
			}
			reqs = append(reqs, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: d.GetNamespace(),
					Name:      d.GetName(),
				},
			})
		}
		return reqs
	}
}

func (r *KluctlDeploymentReconciler) indexBy(kind string) func(o client.Object) []string {
	return func(o client.Object) []string {
		k, ok := o.(*kluctlv1.KluctlDeployment)
		if !ok {
			return nil
		}

		if k.Spec.SourceRef.Kind == kind {
			namespace := k.GetNamespace()
			if k.Spec.SourceRef.Namespace != "" {
				namespace = k.Spec.SourceRef.Namespace
			}
			return []string{fmt.Sprintf("%s/%s", namespace, k.Spec.SourceRef.Name)}
		}

		return nil
	}
}
