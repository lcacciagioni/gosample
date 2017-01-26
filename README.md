# GO Sample | Cloudfoundry
This is a simple app created to work and show the capabilities of golang when deployed to [Cloudfoundry](https://www.cloudfoundry.org/).

> Intended (but not limited) to be used with [pcf-dev](http://pivotal.io/pcf-dev). This project also assumes that you have a functional version of golang.

## First push
All you have to do is to clone this repo and then move to it and run:
```bash
$ go get -u github.com/lcacciagioni/gosample
$ cd $GOPATH/src/github.com/lcacciagioni/gosample
$ cf push
Using manifest file $GOPATH/src/github.com/lcacciagioni/gosample/manifest.yml
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
Installing package '.' (default)
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
It will use the default manifest with almost no requirements as you can see. Now if you do: `$ curl gosample.local.pcfdev.io` you will gonna be able to see something like:
```
Hello, World! from GO

Complete Environment: 

 [CF_INSTANCE_ADDR=10.0.2.15:60200 TMPDIR=/home/vcap/tmp USER=vcap VCAP_APPLICATION={"application_id":"e78e2dd4-8e0a-4ab9-9cd0-38b63ea72dc5","application_name":"gosample","application_uris":["gosample.local.pcfdev.io"],"application_version":"aac0a29b-6bfb-48f3-bd95-9f72c7aca9f2","cf_api":"http://api.local.pcfdev.io","host":"0.0.0.0","instance_id":"3a59e422-e25d-48c5-7b1a-3dc1279690e6","instance_index":1,"limits":{"disk":64,"fds":16384,"mem":32},"name":"gosample","port":8080,"space_id":"0f9dede9-6eab-496f-a39c-8b419e1d2f9f","space_name":"pcfdev-space","uris":["gosample.local.pcfdev.io"],"version":"aac0a29b-6bfb-48f3-bd95-9f72c7aca9f2"} CF_INSTANCE_GUID=3a59e422-e25d-48c5-7b1a-3dc1279690e6 PATH=/usr/local/bin:/usr/bin:/bin:/home/vcap/app/bin PWD=/home/vcap/app LANG=en_US.UTF-8 CF_INSTANCE_PORT=60200 VCAP_SERVICES={} CF_INSTANCE_IP=10.0.2.15 CF_INSTANCE_INDEX=1 HOME=/home/vcap/app SHLVL=1 CF_INSTANCE_PORTS=[{"external":60200,"internal":8080},{"external":60201,"internal":2222}] INSTANCE_INDEX=1 PORT=8080 INSTANCE_GUID=3a59e422-e25d-48c5-7b1a-3dc1279690e6 MEMORY_LIMIT=32m _=/home/vcap/app/bin/gosample]

Super I'm running in CloudFoundry and this are my variables:
ID: 3a59e422-e25d-48c5-7b1a-3dc1279690e6
Index: 1
Name: gosample
Host: 0.0.0.0
Port: 8080
Version: aac0a29b-6bfb-48f3-bd95-9f72c7aca9f2
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
After doing this a new section will appear when you do `$ curl gosample.local.pcfdev.io` with something similar to this:
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
### MySql
Execute this commands to have mysql env vars in place.
```bash
$ cf cs p-mysql 1gb mysql
$ cf bs gosample mysql
$ cf restage gosample
```
And then the MySql section will appear with something similar to this:
```
MYSQL: true

DB Host:  mysql-broker.local.pcfdev.io
DB Port:  3306
DB Name:  cf_7f27f3fe_5f38_4f3f_be18_3c50dad5458a
DB User:  eRhRJwVDmH3Ah4Us
DB Pass:  Yc9Py9Llux4uTs38
```
### Redis
Execute this commands to have redis env vars in place.
```bash
$ cf cs p-redis shared-vm redis
$ cf bs gosample redis
$ cf restage gosample
```
And then the Redis section will appear with something similar to this:
```
REDIS: true

Host:  redis.local.pcfdev.io
Pass:  8e8d56ff-b0d9-4f59-b7de-790012711aef
```

## VERSIONS

```
go: go version go1.7.4 linux/amd64
glide: glide version v0.12.3
pcf-dev: PCF Dev version 0.23.0 (CLI: 474b3ba, OVA: 0.436.0)
cf cli: cf version 6.23.1+a70deb3.2017-01-13
```

### Available marketplace
```
Getting services from marketplace in org pcfdev-org / space pcfdev-space as user...
OK

service        plans             description
local-volume   free-local-disk   Local service docs: https://github.com/cloudfoundry-incubator/local-volume-release/
p-mysql        512mb, 1gb        MySQL databases on demand
p-rabbitmq     standard          RabbitMQ is a robust and scalable high-performance multi-protocol messaging broker.
p-redis        shared-vm         Redis service to provide a key-value store

TIP:  Use 'cf marketplace -s SERVICE' to view descriptions of individual plans of a given service.
```

## TODO
- [x] Create a Docker Image with same code (Prove docker workflow)
- [ ] Do something using Rabbit
- [ ] Do something using MySql
- [ ] Do Something using Redis