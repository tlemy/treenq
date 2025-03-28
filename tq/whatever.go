package tq

import (
	tqsdk "github.com/treenq/treenq/pkg/sdk"
)

func Build() (tqsdk.Space, error) {
	return tqsdk.Space{
		Key:    "treenq-poc",
		Region: "fra1",
		Service: tqsdk.Service{
			DockerfilePath: "Dockerfile",
			SizeSlug:       "basic-xxs",
			Name:           "treenq-poc-service",
			HttpPort:       8000,
			Replicas:       1,
		},
	}, nil
}
