# kubewatch

Monitoring kubenetes cluster and showing pods/containers status

Fork from client-go [workqueue example](https://github.com/kubernetes/client-go/tree/master/examples/workqueue)

# Usage

```shell
kubewatch -h
Usage of kubewatch:
  -all-namespaces
        all namespaces
  -alsologtostderr
        log to standard error as well as files
  -kubeconfig string
        absolute path to the kubeconfig file (default "/Users/johnlin/.kube/config")
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -master string
        master url
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```
