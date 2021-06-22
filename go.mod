module github.com/crossplane/provider-openstack

go 1.13

require (
	github.com/crossplane/crossplane-runtime v0.13.0
	github.com/crossplane/crossplane-tools v0.0.0-20201201125637-9ddc70edfd0d
	github.com/google/go-cmp v0.5.2
	github.com/gophercloud/gophercloud v0.1.0
	github.com/pkg/errors v0.9.1
	google.golang.org/api v0.20.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	sigs.k8s.io/controller-runtime v0.8.0
	sigs.k8s.io/controller-tools v0.3.0
)

replace github.com/crossplane/provider-openstack => /home/ricardo/ws/provider-openstack
