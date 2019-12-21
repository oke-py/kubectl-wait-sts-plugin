# kubectl-wait-sts-plugin
kubectl wait-sts plugin makes it possible to wait until statefulset gets ready.
This plugin is a workaround for [kubernetes/kubernetes#79606](https://github.com/kubernetes/kubernetes/issues/79606).

```
kubectl wait-sts -h
Wait until Statefulset gets ready

Usage:
  wait-sts [statefulset-name] [flags]

Examples:

        # wait for statefulset
        kubectl wait-sts <statefulset>

        # wait for statefulset in different namespace
        kubectl wait-sts <statefulset> -n/--namespace <ns>


Flags:
  -h, --help               help for wait-sts
  -n, --namespace string   override the namespace defined in the current context
      --timeout duration   The length of time to wait before giving up.  Zero means check once and don't wait, negative means wait for a week. (default 30s)
```

## Install the plugin

### Krew

Preparing...

### Binary

You can find the latest binaries in the [releases](https://github.com/oke-py/kubectl-wait-sts-plugin/releases) section.
To install it, place it somewhere in your `$PATH` for `kubectl` to pick it up.

### Build

```
# Clone this repository (or your fork)
git clone https://github.com/oke-py/kubectl-wait-sts-plugin.git
cd kubectl-wait-sts-plugin
make build
```

## License

This software is released under the MIT License.
