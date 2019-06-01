# KubeSelect

This is a little tool that allows you to select the KUBECTL
configuration and context that you want to work with right now. It
basically wraps the kubectl command.

## Usage

Select a kubeconfig file (stored inside ~/.kube) and a context within
each file:

```
$ eval $(kubeselect select)
```

On my own laptop I've added the following shell alias to make this a
bit more convenient:

```
$ alias ks='eval $(kubeselect select)'
```

To find out, what environment you've currently selected, run
`kubeselect status`:

```
$ kubeselect status
do-h10n.yaml // do-fra1-do-h10n
```

I've tried to compress the output as much as possible so that it can
also be used inside a shell prompt.

The `kubeselect run` command basically just wraps kubectl but uses the
specified environment variables for setting the `--context` flag.

Again, I've defined an alias to make my life easier:

```
$ alias k='kubeselect run -- '
```
