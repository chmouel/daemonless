module github.com/chmouel/daemonless

go 1.14

replace (
	github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78
	github.com/containerd/containerd => github.com/containerd/containerd v0.2.10-0.20170808145631-06b9cb351610
	github.com/containerd/containerd v1.3.2 => github.com/containerd/containerd v0.0.0-20191014053712-acdcf13d5eaf
	github.com/docker/docker => github.com/docker/docker v0.0.0-20190404075923-dbe4a30928d4
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190708153700-3bdd9d9f5532
	google.golang.org/grpc => google.golang.org/grpc v1.12.0
	k8s.io/api => k8s.io/api v0.18.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.0
	k8s.io/apiserver => k8s.io/apiserver v0.18.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.0
	k8s.io/client-go => k8s.io/client-go v0.18.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.0
	k8s.io/code-generator => k8s.io/code-generator v0.18.0
	k8s.io/component-base => k8s.io/component-base v0.18.0
	k8s.io/cri-api => k8s.io/cri-api v0.18.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.0
	k8s.io/kubectl => k8s.io/kubectl v0.18.0
	k8s.io/kubelet => k8s.io/kubelet v0.18.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.0
	k8s.io/metrics => k8s.io/metrics v0.18.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.0
)

require (
	github.com/containers/buildah v1.14.9
	github.com/containers/image/v5 v5.4.3
	github.com/containers/storage v1.19.2
	github.com/fsouza/go-dockerclient v1.6.5
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/mtrmac/gpgme v0.1.2
	github.com/opencontainers/runtime-spec v1.0.2
	github.com/openshift/api v0.0.0-20200512152615-944c57cb1477
	github.com/openshift/builder v4.0.0+incompatible
	github.com/openshift/client-go v0.0.0-20200326155132-2a6cd50aedd0
	github.com/openshift/source-to-image v1.3.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9
	k8s.io/component-base v0.18.0
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.18.2
)
