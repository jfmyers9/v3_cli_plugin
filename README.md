CF V3 Cli Plugin
================

This is a CF cli plugin to interract with the new v3 CloudController Api.

## Installation
```
./bin/install
```

## Commands
| command | usage | description|
|:-----------------|:-------------------|:---------------------------------------------|
| `create-v3-app`| `cf create-v3-app <app-name>`| Creates a V3 application in the targeted space.|
| `delete-v3-app`| `cf delete-v3-app [-f] <app-name>`| Deletes a V3 application in the targeted space.|
| `v3-app`| `cf v3-app <app-name>`| Retrieves a V3 application in the targeted space.|
| `v3-apps`| `cf v3-apps`| Retrieves the V3 applications in the targeted space.|
| `procfile`| `cf procfile <app-name> <path-to-procfile>`| Creates processes defined in procfile on the application in the targeted space.|
| `remove-process`| `cf remove-process <app-name> <process-type>`| Removes the process of specified type from the application in the targeted space.|
| `processes`| `cf processes <app-name>`| Lists processes on the application in the targeted space.|
