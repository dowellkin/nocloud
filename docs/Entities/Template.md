# Instance (service) Template

NoCloud services are inspired by Docker Compose files, and have simillar top-level schema
Which means each service may contain more than one Instance by default(don't worry if, it's only one instance, we'll manage to not show it nested in the Frontend)

First things first you should define a name:

```yaml
name: Website
```

Now we can proceed with services, let's make it just one default service deployable to IONe only for now.

So we add `instances` section and `default` instance called "example.com"

```yaml
name: Website
instances:
    default:
        name: example.com
```

```yaml
name: Website
instances:
    default:
        name: example.com
        compute: # Instance part
            ione:
                cpu: 1
                ram: 2
                service_type: paas # or iaas
                # host: auto | host_id 
        network:
            ione:
                - auto # One from settings gonna be taken
                # or
                - dummy # Meaning deploy driver doesn't need 
        storage:
            ione:
                type: default
                drive: ssd # or hdd
                size: 10
        extra:
            ione:
                ansible:
                    playbook: 14
                    vars:
                        email: "user@example.com"
```
