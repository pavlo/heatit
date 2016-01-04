# heatit
A command line tool that simplifies HEAT templates authoring and processing.

[![Build Status](https://travis-ci.org/pavlo/heatit.svg?branch=master)](https://travis-ci.org/pavlo/heatit)

```sh
$ heatit process --source=heat.yaml --params=params.yaml --destination=result.yaml
```

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

## Directives

When you launch `heatit`, it reads the source file line by line and seeks for specific marks in each line. Those marks are called directives and have special syntax. Upon encountering a directive, `heatit` processes it. The action it actually does depends on directive type that are described in details below.

`heatit` currently supports two directives - `@insert` and `@param`. The first one is used to insert contents of assets (files and potentially URLs) into the target file while the second one allows to replace pieces with pre-defined values, called parameters.

`heatit` first processes all `@insert` directives recursively and compiles the result. After that the result is seeked for `@param` directives and they get replaced with parameters given in the `--params=foo.yaml` argument.

Detailed documentation for the two directives can be found below.

### The @insert directive

#### Purpose 

This directive prescribes `heatit` to insert specific content in place of the current line in the source file. The basic syntax of this directive is like follows:
 
`@insert: file:/file/whose/content/to/insert/here.txt`

So, the directive begins with `@insert:` mark, followed by `file:` suffix that prescribes `heatit` to read the contents of the following file `/file/whose/content/to/insert/here.txt` and replace the current line with its content.

##### Example of placing @insert into a YAML file

The `@insert` directive can be placed in an `YAML` without breaking the format so you keep the ability to use a YAML editor and validator:

```yaml

heat_template_version: 2014-10-16
parameters:
  flavor:
    @insert: "file:assets/rackspace/flavors.yaml"
resources:
  @insert: "file:assets/rackspace/cloud-files-contaier.yaml"
      
```

The `YAML` above has two `@insert` directives, the first one inserts a list of flavours from `assets/rackspace/flavors.yaml` file:

```yaml

type: string
default: 1GB Standard Instance
constraints:
- allowed_values:
  - 512MB Standard Instance
  - 2 GB Performance
  - 4 GB Performance
  - 8 GB Performance

```

and `assets/rackspace/cloud-files-contaier.yaml` under `resources` section:

```yaml

cloud-files-container:
  type: OS::Swift::Container
  properties:
    name: { get_param: "OS::stack_name" }

```

so, after `heatit` processing, the resulting YAML file will look like follows:

```yaml

heat_template_version: 2014-10-16
parameters:
  flavor:
    type: string
    default: 1GB Standard Instance
    constraints:
    - allowed_values:
      - 512MB Standard Instance
      - 2 GB Performance
      - 4 GB Performance
      - 8 GB Performance
resources:
  cloud-files-container:
    type: OS::Swift::Container
    properties:
      name: { get_param: "OS::stack_name" }

```

##### Example of placing @insert into a NON YAML file

You are not limited using YAML files only. You may want to create plain text assets to configure `systemd` units, `userdata`, firewall rules etc. The `@insert` directive however looks basically the same -
 
```txt

[Unit]
Requires=network.target
After=network.target
[Service]
ExecStart=/opt/bin/etcd-env-generator.sh @param:network-interface @param:coreos-token
@insert: file:assets/fleet/chicago

```

So if `assets/fleet/chicago` consists of this:

```txt
[X-Fleet]
MachineMetadata=location=chicago
Conflicts=monitor*
```

Then the result `heatit` compiles would be so:

```txt
[Unit]
Requires=network.target
After=network.target
[Service]
ExecStart=/opt/bin/etcd-env-generator.sh @param:network-interface @param:coreos-token
[X-Fleet]
MachineMetadata=location=chicago
Conflicts=monitor*
```

Again, the assets can have other directives so the work gets done in a recursive manner.