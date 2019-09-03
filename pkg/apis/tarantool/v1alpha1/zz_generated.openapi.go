// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.Cluster":                  schema_pkg_apis_tarantool_v1alpha1_Cluster(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterSpec":              schema_pkg_apis_tarantool_v1alpha1_ClusterSpec(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterStatus":            schema_pkg_apis_tarantool_v1alpha1_ClusterStatus(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplate":       schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplate(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplateSpec":   schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplateSpec(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplateStatus": schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplateStatus(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.Role":                     schema_pkg_apis_tarantool_v1alpha1_Role(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleSpec":                 schema_pkg_apis_tarantool_v1alpha1_RoleSpec(ref),
		"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleStatus":               schema_pkg_apis_tarantool_v1alpha1_RoleStatus(ref),
	}
}

func schema_pkg_apis_tarantool_v1alpha1_Cluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Cluster is the Schema for the clusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterSpec", "github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_ClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterSpec defines the desired state of Cluster",
				Properties: map[string]spec.Schema{
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_ClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterStatus defines the observed state of Cluster",
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicasetTemplate is the Schema for the replicasettemplates API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/apps/v1.StatefulSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplateStatus", "k8s.io/api/apps/v1.StatefulSetSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicasetTemplateSpec defines the desired state of ReplicasetTemplate",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_ReplicasetTemplateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicasetTemplateStatus defines the observed state of ReplicasetTemplate",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_Role(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Role is the Schema for the roles API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleSpec", "github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.RoleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_RoleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RoleSpec defines the desired state of Role",
				Properties: map[string]spec.Schema{
					"numReplicasets": {
						SchemaProps: spec.SchemaProps{
							Description: "NumReplicasets is a number of StatefulSets (Tarantol replicasets) created under this Role",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"storageTemplate": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplate"),
						},
					},
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "Selector is a LabelSelector to find ReplicasetTemplate resources from which StatefulSet created",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tarantool/tarantool-operator/pkg/apis/tarantool/v1alpha1.ReplicasetTemplate", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_tarantool_v1alpha1_RoleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RoleStatus defines the observed state of Role",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
