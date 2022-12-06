package mariadb

import (
	galerav1 "github.com/openstack-k8s-operators/galera-operator/api/v1beta1"
	mariadbv1 "github.com/openstack-k8s-operators/mariadb-operator/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A shim to interact with other Openstack object until the MariaDB CR is replaced with Galera
func MariaDBShim(m *galerav1.Galera) *mariadbv1.MariaDB {
	mariadb := &mariadbv1.MariaDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
			Labels:    GetLabels(m.Name),
		},
		Spec: mariadbv1.MariaDBSpec{
			Secret:         m.Spec.Secret,
			StorageClass:   m.Spec.StorageClass,
			StorageRequest: m.Spec.StorageRequest,
			ContainerImage: m.Spec.ContainerImage,
		},
	}
	return mariadb
}
