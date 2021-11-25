# coderlytics.io file-server

The idea of this project is to have a small server which serves defined files via webservice.

## Getting started

1. Download the latest release (currently only available for windows)
2. Copy the example configuration from the "config/" directory
3. Start the file server "file-server.exe -config config.yml"

## Configuration

Example configuration:

```yml
file-server:
  port: 5140 # Port on which the fle server is listening
  files: # List of files to serve
    - name: Example file # Human readable name of this configuration. Has no effect on the way of how the file gets served
      file: /var/www/test.json # Path to the file which should get served
      endpoint: /myfolder/test.json # Endpoint from which the file can be loaded from the server
      token: 4790d6b24dbd11ec81d30242ac130003 #S ecurity token used for authentication (optional)

logging:
  level: info # Log level. Everything above the defined level will be printed in the console. Possible values are trace, debug, info, warn, error, fatal, panic
```

The security token has to be set in the requester header. e.g.
```http
GET http://localhost:5140/myfolder/test.json
Authorization: 4790d6b24dbd11ec81d30242ac130003
```