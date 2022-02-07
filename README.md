# nostdglobals

nostdglobals is a simple Go linter that checks for usages of global variables defined in the go standard library

## Install

go `>= 1.16`

```shell
go install github.com/katsadim/nostdglobals
```

## Usage
To lint all the packages in a program:
```shell
> nostdglobals ./...
```

## std stands for standard library

After having a look at Seth Vargo's excellent [blog post](https://www.sethvargo.com/what-id-like-to-see-in-go-2/), I figured 
that it was about time to take matters into my own hands and create this tool. Here is an excerpt from the post:

> As just one example, both http.DefaultClient and http.DefaultTransport are global variables with shared state. 
> http.DefaultClient has no configured timeout, which makes it trivial to DOS your own service and create bottlenecks. 
> Many packages mutate http.DefaultClient and http.DefaultTransport, which can waste days of developer resources 
> tracking down bugs.
> 
> ...
> 
> I also worry about this class of issues from a software supply chain standpoint. If I can develop a useful package 
> that secretly modifies the http.DefaultTransport to use a custom RoundTripper that funnels all your traffic through 
> my servers, that would make for a very bad time.

## Support

For now this linter only reports http.DefaultClient and DefaultTransport. More to come soon!

## Scan popular projects

### Kubernetes

```shell
katsadim > ~/go/bin/nostdglobals ./...
~/dev/kubernetes/staging/src/k8s.io/client-go/transport/cache.go:87:10: should not make use of 'http.DefaultTransport'
~/dev/kubernetes/staging/src/k8s.io/client-go/rest/request.go:680:12: should not make use of 'http.DefaultClient'
~/dev/kubernetes/staging/src/k8s.io/client-go/rest/request.go:816:12: should not make use of 'http.DefaultClient'
~/dev/kubernetes/staging/src/k8s.io/client-go/rest/request.go:946:12: should not make use of 'http.DefaultClient'
~/dev/kubernetes/staging/src/k8s.io/client-go/rest/transport.go:38:18: should not make use of 'http.DefaultTransport'
~/dev/kubernetes/staging/src/k8s.io/client-go/rest/transport.go:44:16: should not make use of 'http.DefaultClient'
```

### google-api-go-client
```shell
katsadim > ~/go/bin/nostdglobals ./...     
~/dev/google-api-go-client/internal/gensupport/send.go:35:12: should not make use of 'http.DefaultClient'
~/dev/google-api-go-client/internal/gensupport/send.go:69:12: should not make use of 'http.DefaultClient'
~/dev/google-api-go-client/googleapi/transport/apikey.go:34:8: should not make use of 'http.DefaultTransport'
~/dev/google-api-go-client/transport/http/dial.go:166:27: should not make use of 'http.DefaultTransport'
~/dev/google-api-go-client/examples/main.go:72:29: should not make use of 'http.DefaultTransport'
~/dev/google-api-go-client/google-api-go-generator/gen.go:363:14: should not make use of 'http.DefaultClient'
~/dev/google-api-go-client/idtoken/idtoken.go:57:41: should not make use of 'http.DefaultTransport'
~/dev/google-api-go-client/idtoken/validate.go:33:57: should not make use of 'http.DefaultClient'
```

### aws-sdk-go
Disclaimer: this behaviour is documented
```shell
katsadim > ~/go/bin/nostdglobals ./...  
~/dev/aws-sdk-go/aws/corehandlers/handlers.go:126:15: should not make use of 'http.DefaultTransport'
~/dev/aws-sdk-go/aws/defaults/defaults.go:59:18: should not make use of 'http.DefaultClient'
~/dev/aws-sdk-go/example/aws/request/httptrace/config.go:45:11: should not make use of 'http.DefaultTransport'
```

### istio
```shell
katsadim > ~/go/bin/nostdglobals ./...  
~/dev/istio/pkg/kube/client.go:729:15: should not make use of 'http.DefaultClient'
~/dev/istio/pilot/cmd/pilot-agent/status/server.go:591:15: should not make use of 'http.DefaultClient'
```

## Future work
* Lint Vendor directory
* Add more sketchy global variables
* Introduce configuration support which could contain globals to report