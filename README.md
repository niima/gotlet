# gotlet
Simple Templating command line tool using go template engine

## Instalation
```
chmod +x ./gotlet
mv ./gotlet /usr/local/bin/gotlet
```

## Example usage:

```
./gotlet -t deployment.yaml -d data.yaml -v
```

## Flags:
- `-t` Specify the template file path (required)
- `-d` Specify the data file path containing the variables in Yaml (optional, if not specified you have only environment variables to include)
- `-p` A prefix for filtering which env variables to include
- `-o` Output file path (default: result.yaml)
- `-v` Print out the result in stdout 

## Template Engine Reference: 
To use a variable, eighter from environment variables or the specified use the following syntax which is standard go template syntax.

``` 
statictext {{ .variable_name }} static text
```
Nested variables:
```
statictext {{ .variable_name.sub_var_name }} static text
```
Environment variables:
```
statictext {{ .USER }} static text
```
Iterate in a key-value dictionary
```
env: 
{{range $key, $value := .environment_variables}}
    - name: {{ $key }}
      value: {{ $value }} 
{{end}}
```

Read more at [official go documents](https://pkg.go.dev/text/template)

## Sample
This is a basic variable file:
> Note that variables root element is required and MUST be `variables`
```yaml
variables:
  service_name: nginx
  version: 2
  component: front-end
  port: 80
  frontend_max_replicas: 3
  frontend_image: "nginx:latest"
  environment_variables:
    environment: production
    api_url: "https://api.url.com"
    project_name: website
```
This is a sample template file which is a kubernetes deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.service_name}}-{{.component}}-{{.version}}
  labels:
    app: {{.service_name}}
    component: {{.component}}
    version: {{.version}}
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 2
      maxSurge: {{ .frontend_max_replicas }}
    spec:
      containers:
        - name: {{.service_name}}
          image: {{.frontend_image}}
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: {{.ops_port}}
          env: {{range $key, $value := .environment_variables}}
            - name: {{ $key }}
              value: {{ $value }} {{end}} 

```

running this command would generate this file:
```
./gotlet -t ./examples/template.yaml -d ./examples/variables.yaml -o export.yaml
```
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-front-end-2
  labels:
    app: nginx
    component: front-end
    version: 2
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 2
      maxSurge: 3
    spec:
      containers:
        - name: nginx
          image: nginx:latest
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: <no value>
          env: 
            - name: api_url
              value: https://api.url.com 
            - name: environment
              value: production 
            - name: project_name
              value: website  

```

## Build from source
To run on your local machine (you need go 1.18+ installed)
```
go build .
```

To run on a linux server or a linux pipeline runner:
```
GOOS=linux GOARCH=amd64 go build -o gotlet-amd64-linux main.go
```

### TODO:
- Increase test coverage
- Better YAML validation