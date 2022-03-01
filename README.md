# mini-kube-demo

Mini Kube Demo is a repository holding files and instructions required for a simple technical demo a simple [Golang](https://go.dev) web application deployed in loadbalanced fashion on [Kubernetes](https://kubernetes.io) (with use of [MiniKube](https://minikube.sigs.k8s.io/docs/)).

## Installation

The demo requires following software to be installed on the system:

* [MiniKube](https://minikube.sigs.k8s.io/docs/)
* [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)
* [Git](https://git-scm.com)

Additionally, following software could be installed for development of the demo:

* [Docker](https://www.docker.com)
* [Golang](https://go.dev)

Installing of the required software is not in the scope of this demo, therefore the users should handle the installation themselves, accoringly to the documentation from respective projects, or seek technical help.

Mini Kube Demo requires an access to a system terminal (ie. `Terminal.app` on MacOS Intel/AppleSilicon) and some terminal proficiency.

## Usage

### Demo time

Follow the steps below in the terminal.

1. Start a Mini Kube cluster

        $ minikube start
        😄  minikube v1.25.2 on Darwin 12.2.1 (arm64)
        (... stripped ...)
        🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
        $

2. Clone the Mini Kube Demo repository and enter to the folder

        $ git clone git@github.com:bartekrutkowski/mini-kube-demo.git && cd mini-kube-demo
        Cloning into 'mini-kube-demo'...
        (... stripped ...)
        $

3. Deploy the Deployment, Service and Ingress resources on the Mini Kube cluster using `kubectl` and the `mini-kube-demo-deployment.yaml` file

        $ kubectl apply -f mini-kube-demo-deployment.yaml
        namespace/hello-world created
        deployment.apps/hello-world created
        service/hello-world created
        ingress.networking.k8s.io/hello-world created
        $

4. Verify the deployment with `kubectl`

        $ kubectl -n hello-world get all
        NAME                               READY   STATUS    RESTARTS   AGE
        pod/hello-world-697cff6b6c-dtb69   1/1     Running   0          85s
        pod/hello-world-697cff6b6c-hkr8w   1/1     Running   0          85s

        NAME                  TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
        service/hello-world   LoadBalancer   10.101.83.79   <pending>     8080:30480/TCP   85s

        NAME                          READY   UP-TO-DATE   AVAILABLE   AGE
        deployment.apps/hello-world   2/2     2            2           85s

        NAME                                     DESIRED   CURRENT   READY   AGE
        replicaset.apps/hello-world-697cff6b6c   2         2         2       85s
        $

5. Launch [Mini Kube Tunnel](https://minikube.sigs.k8s.io/docs/handbook/accessing/#using-minikube-tunnel) service (you may be asked for your operating system password)

        $ minikube tunnel
        ✅  Tunnel successfully started

        📌  NOTE: Please do not close this terminal as this process must stay alive for the tunnel to be accessible ...

        🏃  Starting tunnel for service hello-world.
        ❗  The service/ingress hello-world requires privileged ports to be exposed: [80 443]
        🔑  sudo permission will be asked for it.
        🏃  Starting tunnel for service hello-world.
        Password:

    Note that `minikube tunnel` command wont return to the command prompt and needs to run until the end of the demo, otherwise there will be no connectivity to the application running on the cluster.

6. Verify that the application is running in loadbalanced fashion with `curl`

        $ for i in $(seq 0 10); do curl http://localhost:8080/; done
        Hello world from fVoV2t
        Hello world from fVoV2t
        Hello world from JppA1h
        Hello world from 6LvgZD
        Hello world from JppA1h
        Hello world from 6LvgZD
        Hello world from JppA1h
        Hello world from JppA1h
        Hello world from EVewSe
        Hello world from EVewSe
        Hello world from fVoV2t
        $

    Note, how the ID string of the pod changes randomly, depending on which pod out of the default 4 replicas serves the HTTP request.

    Alternatively you can open a web browser window and paste the `http://localhost:8080/` url to the address bar, refreshing few times - you should be able to observe the same changing names.

### Cleanup

After experimenting with the demo, follow the steps below in the terminal.

1. Stop the `minikube tunnel` with `^C` (CTL + C)

        ✅  Tunnel successfully started

        📌  NOTE: Please do not close this terminal as this process must stay alive for the tunnel to be accessible ...

        🏃  Starting tunnel for service hello-world.
        ❗  The service/ingress hello-world requires privileged ports to be exposed: [80 443]
        🔑  sudo permission will be asked for it.
        🏃  Starting tunnel for service hello-world.
        Password:
        ^C✋  Stopped tunnel for service hello-world.
        ✋  Stopped tunnel for service hello-world.
        $

2. Delete the resources from the Mini Kube cluster

        $ kubectl delete -f mini-kube-demo-deployment.yaml
        namespace "hello-world" deleted
        deployment.apps "hello-world" deleted
        service "hello-world" deleted
        ingress.networking.k8s.io "hello-world" deleted
        $

3. Verify with `kubectl` that the resources are deleted properly

        $ kubectl -n hello-world get all
        No resources found in hello-world namespace.
        $

4. Delete the MiniKube cluster with `minikube`

        $ minikube delete
        🔥  Deleting "minikube" in docker ...
        🔥  Deleting container "minikube" ...
        🔥  Removing /Users/r/.minikube/machines/minikube ...
        💀  Removed all traces of the "minikube" cluster.
        $

And that's it, the demo is concluded, the environment is cleaned up.

## Advanced usage

### Web app

The web application used in this demo is a very simple Golang application that doesn't take any configuration and doesn't have any other functionality than responding to HTTP requests coming in on port `8080` with `Hello world from eENVcp` string, where the last 6 characters of the response are randomly generated once every start of the application.

The app can be easily tested with `go run` command, assuming we have the repository cloned already from the above instructions.

        $ cd hello-world && go run ./cmd/web/
        2022/03/01 18:31:48 starting PpzK1B http server on port 8080

When sending requests with `curl` (or by using web browser) the app will log each time it has served a request

    2022/03/01 18:31:48 starting PpzK1B http server on port 8080
    2022/03/01 18:33:26 sent response to incoming request
    2022/03/01 18:33:26 sent response to incoming request
    2022/03/01 18:33:26 sent response to incoming request
    2022/03/01 18:33:26 sent response to incoming request
    2022/03/01 18:33:26 sent response to incoming request
    2022/03/01 18:33:26 sent response to incoming request

Note, in such case the response given by the app won't change, since there is only one instance running during such testing and the random string stays the same for each response until the app is restarted.

### Docker image

If the Docker image from the public repository published on [Docker Hub](https://hub.docker.com/r/bartekrutkowski/mini-kube-demo) doesn't work on the testing machine due to operating system or architecture differences, one can build it's own Docker image with `docker` command, assuming we have the repository cloned already from the above instructions.

        $ cd hello-world && docker build . -q -t mini-kube-demo:latest
        sha256:719f885a0c36e2f5979497f88ea604bde2d1cf623f075f2030444ba9e81bf933
        $

The image built this way needs to be then properly tagged and pushed to a container registry of choice. Once that's done, the `mini-kube-demo-deployment.yaml` file needs to be edited where the container `image` is set.

        (... stripped ...)
        27     spec:
        28       containers:
        29         - name: hello-world
        30           image: bartekrutkowski/mini-kube-demo:latest
        31           imagePullPolicy: Always
        32           ports:
        (... stripped ...)

## License

While I retain the author's rights for the application, the entire repository and all its artifacts are relased under the [BSD 3-Clause license](https://github.com/bartekrutkowski/mini-kube-demo/blob/main/LICENSE) and is free to use.
