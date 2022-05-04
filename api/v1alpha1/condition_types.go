/*
Copyright 2021 The Flux authors

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

package v1alpha1

const (
	// HealthyCondition represents the last recorded
	// health assessment result.
	HealthyCondition string = "Healthy"
)

const (
	// PruneFailedReason represents the fact that the
	// pruning of the KluctlDeployment failed.
	PruneFailedReason string = "PruneFailed"

	// ArtifactFailedReason represents the fact that the
	// source artifact download failed.
	ArtifactFailedReason string = "ArtifactFailed"

	// PrepareFailedReason represents failure in the kluctl preparation phase
	PrepareFailedReason string = "PrepareFailed"

	// DeployFailedReason represents the fact that the
	// kluctl deploy command failed.
	DeployFailedReason string = "DeployFailed"

	// DependencyNotReadyReason represents the fact that
	// one of the dependencies is not ready.
	DependencyNotReadyReason string = "DependencyNotReady"

	// ReconciliationSucceededReason represents the fact that
	// the reconciliation succeeded.
	ReconciliationSucceededReason string = "ReconciliationSucceeded"

	// ReconciliationSkippedReason represents the fact that
	// the reconciliation was skipped due to an unchanged target.
	ReconciliationSkippedReason string = "ReconciliationSkipped"
)
