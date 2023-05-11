package dynamicclient

import (
	"context"
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type foo struct {
	dynamicClient *dynamic.DynamicClient
}

func New(configFile string) (*foo, error) {
	// Create a new Kubernetes API client
	config, err := clientcmd.BuildConfigFromFlags("", configFile)

	if err != nil {
		return nil, err
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &foo{
		dynamicClient: client,
	}, nil
}

func (f *foo) List(gvr schema.GroupVersionResource, namespace string) (*unstructured.UnstructuredList, error) {
	// Get the list of all pods in the cluster when namespace is ""
	objs, err := f.dynamicClient.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list %s in namespace %s: %w", gvr.String(), namespace, err)
	}

	return objs, nil
}

func (f *foo) Get(gvr schema.GroupVersionResource, namespace, name string) (*unstructured.Unstructured, error) {
	obj, err := f.dynamicClient.Resource(gvr).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get %s %s/%s: %w", gvr.String(), namespace, name, err)
	}

	return obj, nil
}

func (f *foo) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	gvr := schema.GroupVersionResource{Group: appsv1.SchemeGroupVersion.Group, Version: appsv1.SchemeGroupVersion.Version, Resource: "deployments"}
	obj, err := f.Get(gvr, namespace, name)
	if err != nil {
		return nil, err
	}

	deploy := &appsv1.Deployment{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), deploy)
	if err != nil {
		return nil, err
	}

	return deploy, nil
}

// Update update the given object,
// obj is a k8s resource or a custom resource which implements the Object and Type interfaces
func (f *foo) Update(obj interface{}) error {
	objectMeta, ok := obj.(metav1.Object)
	if !ok {
		return fmt.Errorf("obj has no ObjectMeta")
	}

	typeMeta, ok := obj.(schema.ObjectKind)
	if !ok {
		return fmt.Errorf("obj has no TypeMeta")
	}

	gvk := typeMeta.GroupVersionKind()
	gvr := kindToResource(gvk.GroupVersion(), gvk.Kind)

	data, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return err
	}

	_, err = f.dynamicClient.Resource(gvr).Namespace(objectMeta.GetNamespace()).Update(context.TODO(), &unstructured.Unstructured{Object: data}, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func kindToResource(gv schema.GroupVersion, kind string) schema.GroupVersionResource {
	resource := gv.WithResource(strings.ToLower(kind) + "s")
	return resource
}
