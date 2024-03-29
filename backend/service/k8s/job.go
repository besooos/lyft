package k8s

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	k8sapiv1 "github.com/lyft/clutch/backend/api/k8s/v1"
)

func (s *svc) DescribeJob(ctx context.Context, clientset, cluster, namespace, name string) (*k8sapiv1.Job, error) {
	cs, err := s.manager.GetK8sClientset(ctx, clientset, cluster, namespace)
	if err != nil {
		return nil, err
	}

	jobs, err := cs.BatchV1().Jobs(cs.Namespace()).List(ctx, metav1.ListOptions{
		FieldSelector: "metadata.name=" + name,
	})
	if err != nil {
		return nil, err
	}

	if len(jobs.Items) == 1 {
		return protoForJob(cs.Cluster(), &jobs.Items[0]), nil
	} else if len(jobs.Items) > 1 {
		return nil, status.Error(codes.FailedPrecondition, "located multiple jobs")
	}
	return nil, status.Error(codes.NotFound, "unable to locate specified job")
}

func (s *svc) DeleteJob(ctx context.Context, clientset, cluster, namespace, name string) error {
	cs, err := s.manager.GetK8sClientset(ctx, clientset, cluster, namespace)
	if err != nil {
		return err
	}

	opts := metav1.DeleteOptions{}

	return cs.BatchV1().Jobs(cs.Namespace()).Delete(ctx, name, opts)
}

func (s *svc) CreateJob(ctx context.Context, clientset, cluster, namespace string, job *v1.Job) (*k8sapiv1.Job, error) {
	cs, err := s.manager.GetK8sClientset(ctx, clientset, cluster, namespace)
	if err != nil {
		return nil, err
	}

	opts := metav1.CreateOptions{}

	resultJob, err := cs.BatchV1().Jobs(cs.Namespace()).Create(ctx, job, opts)
	if err != nil {
		return nil, err
	}

	return protoForJob(cs.Cluster(), resultJob), nil
}

func (s *svc) ListJobs(ctx context.Context, clientset, cluster, namespace string, listOptions *k8sapiv1.ListOptions) ([]*k8sapiv1.Job, error) {
	cs, err := s.manager.GetK8sClientset(ctx, clientset, cluster, namespace)
	if err != nil {
		return nil, err
	}

	opts, err := ApplyListOptions(listOptions)
	if err != nil {
		return nil, err
	}

	jobList, err := cs.BatchV1().Jobs(cs.Namespace()).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	var jobs []*k8sapiv1.Job
	for _, j := range jobList.Items {
		job := j
		jobs = append(jobs, protoForJob(cs.Cluster(), &job))
	}

	return jobs, nil
}

func protoForJob(cluster string, k8sJob *v1.Job) *k8sapiv1.Job {
	clusterName := GetKubeClusterName(k8sJob)
	if clusterName == "" {
		clusterName = cluster
	}

	return &k8sapiv1.Job{
		Cluster:     clusterName,
		Namespace:   k8sJob.Namespace,
		Name:        k8sJob.Name,
		Labels:      k8sJob.Labels,
		Annotations: k8sJob.Annotations,
	}
}
