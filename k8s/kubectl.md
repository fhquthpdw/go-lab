### 如何 unset users, clusters, contexts
```shell
kubectl config unset users.USERNAME
kubectl config unset contexts.CONTEXT-NAME
kubectl config unset clusters.CLUSTER-NAME
```

### 如何 delete users, clusters, contexts
```shell
kubectl config delete-user USERNAME
kubectl config delete-cluster CLUSTER-NAME
kubectl config delete-context CONTEXT-NAME
```

### 如何合并多个 kubeconfig 文件内容
下面的命令是一行，把 "/PATH/TO/NEW/configFile" 换成新的 config 文件路径
```shell
cp ~/.kube/config ~/.kube/config.bak && KUBECONFIG=~/.kube/config:/PATH/TO/NEW/configFile kubectl config view --flatten > /tmp/config && mv /tmp/config ~/.kube/config
```
