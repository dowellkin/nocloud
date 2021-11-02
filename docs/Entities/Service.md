# Service Config

```json
{
  "title": "example",
  "computed": {
    "host": {
      "TEMPLATE_ID": {
        "first-host-id": 0,
        "another-host-id": 3
      }
    }
  },
  "instances": [{
    "title": "example-ione-only-instance",
    "configs": [{
      "type": "ione",
      "config": {
        "params": {
          "CPU": 1,
          "VCPU": 1,
          "MEMORY": 1,
          "DRIVE_TYPE": "HDD",
          "DRIVE_SIZE": 10,
          "TEMPLATE_ID": "computed(host)"
        }
      }
    }]
  }]
}
```
