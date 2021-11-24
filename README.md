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
  port: 5140
  files:
    - name: Example file
      file: /var/www/test.json
      endpoint: /myfolder/test.json

logging:
  level: info
```

Configuraiton options

| **Configuration**    | **Description**                                                                                     | **Example**             | **Possible values**                               | **Default value** |
|------------------|-------------------------------------------------------------------------------------------------|---------------------|-----------------------------------------------|---------------|
| file-server      | Configuration of the file server module                                                         |                     |                                               |               |
|     port         | Port on which the file server is listening                                                      | 5140                |                                               |               |
|     files        | List of files to serve                                                                          |                     |                                               |               |
|         name     | Human readable name of this configuration. Has no effect on the way of how the file gets served | Example file        |                                               |               |
|         file     | Path to the file which should get served                                                        | /var/www/test.json  |                                               |               |
|         endpoint | Endpoint from which the file can be loaded from the server                                      | /myfolder/test.json |                                               |               |
| logging          | Configuration of the logging module                                                             |                     |                                               |               |
|     level        | Log level. Everything above the defined level will be printed in the console                    | info                | trace, debug, info, warn, error, fatal, panic | error         |
|                  |                                                                                                 |                     |                                               |               |

