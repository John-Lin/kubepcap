# kubepcap

Kubernetes pod capture.

Monitoring kubenetes cluster and showing pods/containers status with field selector.

Fork from client-go [workqueue example](https://github.com/kubernetes/client-go/tree/master/examples/workqueue)

## Install

```shell
go get -u github.com/John-Lin/kubepcap
```

## Usage

```
kubepcap -h
Usage of kubepcap:
  -all-namespaces
        All namespaces in cluster
  -field-selector string
        Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2).
  -kubeconfig string
        absolute path to the kubeconfig file (default "/Users/johnlin/.kube/config")
  -master string
        master url
  -namespace string
        If present, the namespace scope for this CLI request (default "default")
```

## Example 

```shell
$ kubepcap --all-namespaces --field-selector=status.phase!=Running

$ kubepcap --field-selector=spec.nodeName==docker-for-desktop

$ kubepcap --field-selector=metadata.name==ubuntu
```
