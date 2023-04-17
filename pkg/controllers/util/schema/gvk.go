package schema

import (
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/kubewharf/kubeadmiral/pkg/controllers/common"
)

func IsServiceGvk(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "" && gvk.Kind == common.ServiceKind
}

func IsServiceAccountGvk(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "" && gvk.Kind == common.ServiceAccountKind
}

func IsJobGvk(gvk schema.GroupVersionKind) bool {
	return gvk.Group == batchv1.GroupName && gvk.Kind == common.JobKind
}

func IsPersistentVolumeGvk(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "" && gvk.Kind == common.PersistentVolumeKind
}

func IsPersistentVolumeClaimGvk(gvk schema.GroupVersionKind) bool {
	return gvk.Group == "" && gvk.Kind == common.PersistentVolumeClaimKind
}
