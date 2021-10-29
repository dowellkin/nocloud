Instance:

```json
{
 "uuid": "UUID | string",
 "title": "Instance title | string",
 "state": { init, create, deploy, boot, runn, shut, off, deleted },
 "deploy_id": "Instance ID at back-end(KVM/Docker/etc)",
 "permissions": [ // One to Many
    Permission
 ],
 "host": Host, // One to One
 "resources": { // Instance Resources description per each type
    "type": Object<any>{} // like "aws": { "tier": "tierType" } or/and "kvm": { "cpu": "2", "ram": "4", ... }
 },
 "attributes": { // Other attributes description per each type
    "type": Object<any>{} // like "docker": { "container_name": "not_default_name" }
 }
}
```

Resource Permission:

```json
{
    "account": "UUID of User account or Group | string",
    "resource": "UUID of Resource(Instance, Network, Datastore, etc)",
    "access": { none, read, write, admin }
}
```

Host:

```json
{
 "uuid": "UUID | string",
 "title": "Instance title | string",
 "permissions": [ // One to Many
    Permission
 ],
 "type": "End Host service type(example/KVM/Docker)"
}
```
