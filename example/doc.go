// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:generate sh -c "bash $GARDENER_HACK_DIR/generate-crds.sh -p 20-crd- extensions.gardener.cloud druid.gardener.cloud resources.gardener.cloud"
//go:generate sh -c "bash $GARDENER_HACK_DIR/generate-extension.sh --name=provider-aws --provider-type=aws --component-name=provider-extension --extension-oci-repository=europe-docker.pkg.dev/gardener-project/releases/charts/gardener/extensions/provider-aws:$(cat ../VERSION) --admission-runtime-oci-repository=europe-docker.pkg.dev/gardener-project/releases/charts/gardener/extensions/admission-aws-runtime:$(cat ../VERSION) --admission-application-oci-repository=europe-docker.pkg.dev/gardener-project/releases/charts/gardener/extensions/admission-aws-application:$(cat ../VERSION) --destination=\"$REPO_ROOT/example/extension/extension.yaml\""
//go:generate kustomize build "$REPO_ROOT/example/extension/" -o "$REPO_ROOT/example/extension.yaml"

// Package example contains generated manifests for all CRDs and other examples.
// Useful for development purposes.
package example
