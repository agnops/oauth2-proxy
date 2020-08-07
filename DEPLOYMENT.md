oauth2-proxy
=================

A reverse proxy and static file server that provides authentication using Providers (GitHub, GitLab and others)
to validate accounts by email, domain or group.

## Configuration

The following table lists the configurable parameters of the oauth2-proxy chart and their default values.

|              Parameter               |                             Description                             |                       Default                       |
| ------------------------------------ | ------------------------------------------------------------------- | --------------------------------------------------- |
| `image.repository`                   | Container image name                                                | `agnops/kube-ops-manager`                      |
| `image.tag`                          | Container image tag                                                 | `latest`                                            |
| `image.pullPolicy`                   | Container pull policy                                               | `Always`                                            |
| `provider`                           | Auth Provider [github/gitlab/bitbucket]                             | `github`                                            |
| `clientId`                           | OAuth Client ID                                                     | `""`                                                |
| `clientSecret`                       | OAuth Client Secret                                                 | `""`                                                |
| `cookieSecret`                       | OAuth Cookie Secret                                                 | `""`                                                |
| `nodeSelector`                       | Pod's Node assignment                                               | `{}`                                                |
| `tolerations`                        | Pod tolerations                                                     | `[]`                                                |
| `affinity`                           | Pod affinity policy                                                 | `{}`                                                |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

```
helm upgrade --install oauth2-proxy chart/oauth2-proxy --namespace ci-cd-tools --wait
```