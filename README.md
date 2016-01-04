# heatit
A command line tool that simplifies HEAT templates authoring and processing.

Build status: [![Build Status](https://travis-ci.org/pavlo/heatit.svg?branch=master)](https://travis-ci.org/pavlo/heatit)

## The problem 'heatit' solves

Heat Orchestration Template (HOT) is a template format used by [Openstack](https://www.openstack.org) Orchestration engine to launch cloud infrastructure composed of servers, networks, users, security groups and others. Usually we end up with a huge template file which is not very comfortable to maintain as a whole, it leads to constant copy/paste cycles and issues related it. 

`heatit` is a tool that allows you to compile a HEAT template using reusable pieces called `assets`. Consider the following project directory layout:

```
    ├── assets
    │   ├── flavors.yaml
    │   ├── ssh-keys
    │   │   ├── id_dsa
    │   │   └── id_dsa.pub
    │   ├── systemd
    │   │   ├── docker.service
    │   │   ├── drop-ins
    │   │   │   ├── etcd2-private-networking.txt
    │   │   │   └── etcd2-timeout.txt
    │   │   └── etcd-env-generator.service
    │   └── userdata
    │       ├── coreos.txt
    │       └── debian.txt
    ├── heat.yaml
    └── params.yaml
        
```

So, there's `heat.yaml` file in the root of the project directory. Pay attention to `@insert` and `@param` directives inline:

```yaml
heat_template_version: 2014-10-16
parameters:
  flavor:
    @insert: "file:assets/flavors.yaml"  <= inserts the value of assets/flavors.yaml here
resources:
    servers:
      type: OS::Heat::ResourceGroup
      properties:
        count: "@param:servers_count"   <= inserts the value of `servers_count` parameter
        resource_def:
          type: "OS::Nova::Server"
          properties:
            name: coreos-%index%
              ...
              user_data_format: RAW
              config_drive: "true"
              user_data: |
                @insert: "file:assets/userdata/coreos.txt" <= inserts assets/userdata/coreos.txt here
outputs:
  ...
```

`heatit` will process the `heat.yaml` in the root of the project and:

 1. Insert the content of `assets/flavors.yaml` and `assets/userdata/coreos.txt` in place of corresponding `@insert` directives.
 2. Insert parameter values (read from `params.yaml` file in the project root directory) in place of `@param` directive.

Using the similar technique, `heatit` can be used to produce a fully featured HEAT template with SSH keys, systemd/fleet units, networks, security groups while keeping everything modular in a highly reusable manner.