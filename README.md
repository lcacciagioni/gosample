# GO Sample | Cloudfoundry
This is a simple app created to work and show the capabilities of golang when deployed to [Cloudfoundry](https://www.cloudfoundry.org/).Using manifest file /home/lcacciagioni/workspace/go/gopath/src/github.com/lcacciagioni/gosample/manifest.yml


## First push
All you have to do is to clone this repo and then move to it and run:
```bash
$ cf push

Updating app gosample in org pcfdev-org / space pcfdev-space as user...
OK

Uploading gosample...
Uploading app files from: $GOPATH/src/github.com/lcacciagioni/gosample
Uploading 22K, 25 files

                             
Done uploading
OK

Stopping app gosample in org pcfdev-org / space pcfdev-space as user...
OK

Starting app gosample in org pcfdev-org / space pcfdev-space as user...
Downloading dotnet-core_buildpack...
Downloading java_buildpack...
Downloading ruby_buildpack...
Downloading nodejs_buildpack...
Downloading go_buildpack...
Downloaded nodejs_buildpack
Downloading python_buildpack...
Downloaded ruby_buildpack
Downloading php_buildpack...
Downloaded java_buildpack
Downloading staticfile_buildpack...
Downloaded dotnet-core_buildpack
Downloading binary_buildpack...
Downloaded go_buildpack
Downloaded php_buildpack
Downloaded python_buildpack
Downloaded binary_buildpack
Downloaded staticfile_buildpack
Creating container
Successfully created container
Downloading app package...
Downloaded app package (21.5K)
Downloading build artifacts cache...
Downloaded build artifacts cache (66.6M)
Staging...
-------> Buildpack version 1.7.15
file:///tmp/buildpacks/4351f5cc91c6d5bb8d11a7418ea9cbff/dependencies/https___buildpacks.cloudfoundry.org_concourse-binaries_godep_godep-v75-linux-x64.tgz
file:///tmp/buildpacks/4351f5cc91c6d5bb8d11a7418ea9cbff/dependencies/https___buildpacks.cloudfoundry.org_concourse-binaries_glide_glide-v0.12.3-linux-x64.tgz
-----> Checking Godeps/Godeps.json file.
-----> Using go1.7.3
[1;33m !!    Installing package '.' (default)[0m
-----> Running: go install -v -tags cloudfoundry . 
github.com/lcacciagioni/gosample/vendor/github.com/mitchellh/mapstructure
github.com/lcacciagioni/gosample/vendor/github.com/cloudfoundry-community/go-cfenv
github.com/lcacciagioni/gosample
Exit status 0
Staging complete
Uploading droplet, build artifacts cache...
Uploading build artifacts cache...
Uploading droplet...
Uploaded build artifacts cache (66.5M)
Uploaded droplet (2M)
Uploading complete
Destroying container
Successfully destroyed container

2 of 2 instances running

App started


OK

App gosample was started using this command `gosample`

Showing health and status for app gosample in org pcfdev-org / space pcfdev-space as user...
OK

requested state: started
instances: 2/2
usage: 32M x 2 instances
urls: gosample.local.pcfdev.io
last uploaded: Wed Jan 25 14:06:34 UTC 2017
stack: cflinuxfs2
buildpack: Go

     state     since                    cpu    memory        disk          details
#0   running   2017-01-25 03:07:05 PM   0.0%   2.4M of 32M   7.2M of 64M
#1   running   2017-01-25 03:07:05 PM   0.0%   2.7M of 32M   7.2M of 64M
```
It will use the default manifest with almost no requirements as you can see. Now if you do: `curl gosample.local.pcfdev.io` you will gonna be able to see something like:
```
Hello, World! from GO
Super I'm running in CloudFoundry and this are my variables:
ID: ba6a61df-5134-4fbc-728b-ab7966b80d4d
Index: 1
Name: gosample
Host: 0.0.0.0
Port: 8080
Version: c50d21c5-41df-4ef9-94ea-9a4f5a4b39e2
Home: /home/vcap/app
MemoryLimit: 32m
WorkingDir: /home/vcap/app
TempDir: /home/vcap/tmp
User: vcap

MYSQL: false

RABBITMQ: false

REDIS: false
```
## Enabling Services
In the next few lines I'll try to show you how to enable the different services to see some of the exposed variables using also `curl`.
### RabbitMQ
Execute this commands to have rabbitmq env vars in place.
```bash
$ cf cs p-rabbitmq standard rabbitmq
$ cf bs gosample rabbitmq
$ cf restage gosample
```
After doing this a new section will appear when you do `curl gosample.local.pcfdev.io` with something similar to this:
```
RABBITMQ: true

AMQP: 
        host:  rabbitmq.local.pcfdev.io
        port:  5672
        user:  48c41519-20bb-433b-b477-a6a228c8f463
        pass:  vh2plek65pskf61c67evnu9bgs
        vhost:  0ff2794f-f17d-40ef-8c3e-081c76c1fd9b
        ssl:  false

MGMT: 
        host:  rabbitmq.local.pcfdev.io
        port:  15672
        user:  48c41519-20bb-433b-b477-a6a228c8f463
        pass:  vh2plek65pskf61c67evnu9bgs
        ssl:  false
```

