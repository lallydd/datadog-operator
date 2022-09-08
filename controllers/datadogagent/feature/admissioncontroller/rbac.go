// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package admissioncontroller

import (
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/DataDog/datadog-operator/pkg/extendeddaemonset"
	"github.com/DataDog/datadog-operator/pkg/kubernetes/rbac"
)

func getRBACClusterPolicyRules() []rbacv1.PolicyRule {
	return []rbacv1.PolicyRule{
		// MutatingWebhooksConfigs
		{
			APIGroups: []string{rbac.AdmissionAPIGroup},
			Resources: []string{rbac.MutatingConfigResource},
			Verbs: []string{
				rbac.GetVerb,
				rbac.ListVerb,
				rbac.WatchVerb,
				rbac.CreateVerb,
				rbac.UpdateVerb,
			},
		},
		// ExtendedDaemonsetReplicaSets
		{
			APIGroups: []string{extendeddaemonset.GroupVersion.Group},
			Resources: []string{
				rbac.ExtendedDaemonSetReplicaSetResource,
			},
			Verbs: []string{rbac.GetVerb},
		},
		// Deployments, Replicasets, Statefulsets, Daemonsets
		{
			APIGroups: []string{rbac.AppsAPIGroup},
			Resources: []string{
				rbac.DeploymentsResource,
				rbac.ReplicasetsResource,
				rbac.StatefulsetsResource,
				rbac.DaemonsetsResource,
			},
			Verbs: []string{rbac.GetVerb},
		},
		// Jobs
		{
			APIGroups: []string{rbac.BatchAPIGroup},
			Resources: []string{rbac.JobsResource},
			Verbs: []string{
				rbac.ListVerb,
				rbac.WatchVerb,
				rbac.GetVerb,
			},
		},
		// CronJobs
		{
			APIGroups: []string{rbac.BatchAPIGroup},
			Resources: []string{rbac.CronjobsResource},
			Verbs: []string{
				rbac.ListVerb,
				rbac.WatchVerb,
				rbac.GetVerb,
			},
		},
	}
}

func getRBACPolicyRules() []rbacv1.PolicyRule {
	return []rbacv1.PolicyRule{
		// Secrets
		{
			APIGroups: []string{rbac.CoreAPIGroup},
			Resources: []string{rbac.SecretsResource},
			Verbs: []string{
				rbac.GetVerb,
				rbac.ListVerb,
				rbac.WatchVerb,
				rbac.CreateVerb,
				rbac.UpdateVerb,
			},
		},
	}
}
