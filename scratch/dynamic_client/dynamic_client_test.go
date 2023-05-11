package dynamicclient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGet(t *testing.T) {
	t.Parallel()

	f, err := New("/root/.kube/config")
	require.NoError(t, err)

	namespace := "mec-2ccc9cbebd1b46a2bf188477808ff631"
	name := "deploy"
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	obj, err := f.Get(gvr, namespace, name)
	require.NoError(t, err)
	require.NotNil(t, obj)
}

func TestGetDeployment(t *testing.T) {
	t.Parallel()

	f, err := New("/root/.kube/config")
	require.NoError(t, err)

	namespace := "mec-2ccc9cbebd1b46a2bf188477808ff631"
	name := "deploy"

	deploy, err := f.GetDeployment(namespace, name)
	require.NoError(t, err)
	require.IsType(t, &appsv1.Deployment{}, deploy)
}

func TestList(t *testing.T) {
	t.Parallel()

	f, err := New("/root/.kube/config")
	require.NoError(t, err)

	gvr := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	namespace := "kube-system"

	objs, err := f.List(gvr, namespace)
	require.NoError(t, err)
	require.IsType(t, &unstructured.UnstructuredList{}, objs)

	// Print the name of each pod
	for _, pod := range objs.Items {
		fmt.Println(pod.GetName())
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	f, err := New("/root/.kube/config")
	require.NoError(t, err)

	namespace := "mec-2ccc9cbebd1b46a2bf188477808ff631"
	name := "deploy"

	deploy, err := f.GetDeployment(namespace, name)
	require.NoError(t, err)

	var replicas int32 = 1
	deploy.Spec.Replicas = &replicas

	err = f.Update(deploy)
	require.NoError(t, err)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	f, err := New("/root/.kube/config")
	require.NoError(t, err)

	data, err := os.ReadFile("/tmp/deploy.json")
	require.NoError(t, err)

	deploy := &appsv1.Deployment{}

	err = json.Unmarshal(data, deploy)
	require.NoError(t, err)

	err = f.Create(deploy)
	require.NoError(t, err)
}
