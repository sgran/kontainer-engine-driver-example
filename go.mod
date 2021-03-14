module github.com/sgran/kontainer-engine-driver-example

go 1.15

replace (
	github.com/heptio/authenticator => sigs.k8s.io/aws-iam-authenticator v0.5.2
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
)

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/imdario/mergo v0.3.7
	github.com/kr/pretty v0.2.0 // indirect
	github.com/rancher/kontainer-engine v0.0.4-dev.0.20201223224019-89626b028c6a
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.36.0 // indirect
)
