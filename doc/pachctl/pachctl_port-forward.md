## ./pachctl port-forward

Forward a port on the local machine to pachd. This command blocks.

### Synopsis


Forward a port on the local machine to pachd. This command blocks.

```
./pachctl port-forward
```

### Options

```
  -k, --kubectlflags string   Any kubectl flags to proxy, e.g. --kubectlflags='--kubeconfig /some/path/kubeconfig'
  -p, --port int              The local port to bind to. (default 30650)
  -x, --proxy-port int        The local port to bind to. (default 30081)
  -u, --ui-port int           The local port to bind to. (default 30080)
```

### Options inherited from parent commands

```
      --no-metrics   Don't report user metrics for this command
  -v, --verbose      Output verbose logs
```

### SEE ALSO
* [./pachctl](./pachctl.md)	 - 

###### Auto generated by spf13/cobra on 11-Jan-2018
