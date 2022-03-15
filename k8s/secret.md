## 目录

- [secret 的工作原理](#secret的工作原理)
- [如何从 secret 中导出 TLS 证书](#如何从secret中导出TLS证书)

## secret 的工作原理

## 如何从 secret 中导出 TLS 证书
```shell
kubectl get secrets/{{ $secretName }} --output=jsonpath={.data.tls-ca\\.crt} | base64 --decode > ca_crt.crt
```