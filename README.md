# Gotlet
Simple command line Templating tool using go template engine.

# Overview

Gotlet is a simple and lightweight single binary go application that helps you with your templating needs. 
___


## Features

- Simple and powerful templating capabilities based on the go template engine.
- A lightweigh single binary application that can be built and/or easily shipped to a wide variety of platforms, and works with no hassle.
- You can substitude the variables imported from files and/or environment variables.
- You can limit the environment variable imports based on a prefix, so only the ones you need for a specific bulid are imported.
___

### Sample gotlet directory

```
~/sample-gotlet-directory
  ├── templates
  │       ├── kubernetes-deployment.yaml
  │       ├── kubernetes-namespace.yaml
  │       └── Dockerfile 
  └── variables     
          ├── 
          ├── kubernetes--templateest.yaml
          ├── kubernetes-acceptance.yaml
          ├── kubernetes--varsfileroduction.yaml
          ├── Dockerfile--templateest.yaml
          ├── Dockerfile-acceptance.yaml
          └── Dockerfile--varsfileroduction.yaml
```
___

# Installation 

**Installation from Binary**:

On macOS 
```bash
curl -L github.com/niima/gotlet/releases/latest/download/gotlet-macos -o /usr/local/bin/gotlet
chmod +x ./gotlet
mv ./gotlet /usr/local/bin/gotlet
```

On Linux:
```bash
curl -L github.com/niima/gotlet/releases/latest/download/gotlet-amd64-linux -o ./gotlet
chmod +x ./gotlet
sudo mv ./gotlet /usr/local/bin/gotlet
```
**Installation from Source**: 

Build requirements:
```bash
- golang 1.18+
```

On macOS/Windows run : 
```bash
go build .
```

On Linux :
```bash
GOOS=linux GOARCH=amd64 go build -o gotlet-amd64-linux main.go
```
___

### **Gotlet workflow**:

1. Create a Template file using the go template engine format.
2. Provide Variables to populate the template with (Provided via a file or OS Environment Variables).
3. Enter the command and let the application do its magic
4. Observe the rendered template in STDOUT or a file.

Let's discuss each part of the workflow in detail in the usage section:
___
# Usage and Practices

Parts of the workflow: 
- Template File
- Variables
- Rendered files
___

## Template file

  > The template file is a file in which you want some fields substituded with variables, in this case, the application works based on the go template engine, which we will have a small series of samples below. You can find the full go templating engine reference [here]((https://pkg.go.dev/text/template))
  ___
  ### **Template engine reference**

  To use a variable, either from environment variables or the specified use the following syntax which is standard go template syntax.

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
  > Read more about go templating engine in [Official Go Documents, Template Engine Reference](https://pkg.go.dev/text/template)
___
  ## **Variables**

  You can use include your variables in two different ways:
  
  - OS Environment variables, with or without a prefix
    - in case you're using environment variables, please prefix them and provide the application with that prefix with `--varsfile PREFIX_NAME_ENV_NAME` arguments
      
      From a security standpoint, you should even separate the sensitive environment variables between different application Environments.
      Sample environment variables with a prefix:
      ```
      TEST_APP1_DEV_ENV_1=dev1
      TEST_APP1_DEV_ENV_2=dev2
      TEST_APP2_DEV_ENV_1=dev1
      TEST_APP2_DEV_ENV_2=dev2
      TEST_APP1_PROD_ENV_1=prod1
      TEST_APP1_PROD_ENV_2=prod2
      TEST_APP1_PROD_ENV_1=prod1
      TEST_APP1_PROD_ENV_2=prod2
      ```  
    - Do the separation in a manner in which you can separate every relevant group as much as possible, so that each env/app only has access to what it definitely needs only.

    - Sample command for rendering a template with a **prefixed series of environment variables**: 
      ```
      gotlet --template ./examples/template.yaml --envprefix TEST_APP1_DEV 
      ```

  - A simple `variables.yaml` file with the format below; Create a separate file each separate application group and environment.
      `test-app1-dev-variables.yaml` for test-app1-dev 
      ```yaml
      variables:
        TEST_APP1_DEV_ENV_1=dev1
        TEST_APP1_DEV_EN_2=dev2
      ```
    - Sample command for rendering a template with a **a variables.yaml file**: 
      ```
      gotlet --template ./examples/template.yaml --varsfile TEST_APP1_DEV 
      ```

  ## **Rendered Files**

  As explained with the small examples above, the application renders the template with the provided variables for you. 
  Gotltet is able to export the results to a file or STDOUT. 

  - To have gotlet export the results to a file, use the following command arg 
      ```bash 
      gotlet --template ./example/template.yaml --varsfile ./example/variables.yaml --output ./example/output.yaml
      ```
  - And for having the output in STDOUT, you don't need to do anything, just remove the --output flag. 

> For further CLI Usage, see the section below
___
# Command Line Flags Usage
- `-h` or `--help` Prints the usage request. 
- `--template` **(required)** Path to the the template file.
- `--varsfile` **(optional)** Path to the variables file (optional, if not specified you have only environment variables to include).
- `--envprefix` **(technically optional, logically, please provide a value)** Prefix for filtering which env variables to include. **(recommended if you're using environment variables for populating templates)** 
- `--output` **(optional)** Output file path, if a file name is not provided, the results will be printed in stdout.

## Example Command Line usage:
```
gotlet --template deployment.yaml --varsfile variables.yaml -v 
```
___

# Example Usage

With the basic explanations out of the way, let's use the application for real.
In this example, we're going to:

- Populate our template using the variables found in a `variable.yaml` file.
- Have the application output the results in stdout

### **Steps**:
1. Create a Template file using the go template engine format.
2. Provide Variables to populate the template with (Provided via a file or OS Environment Variables).
3. Enter the command and let the application do its magic
4. Observe the rendered template in STDOUT or a file.

> **Note**: The files we're going to use in the example below can be found in [the examples directory in this repo](./examples/)

> Note that variables root element is required and MUST be `variables`
___
#### Step **1** : creating a `variables.yaml` file.
you can define the variables under `variables` obviously
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
You can also define os application 
___
#### Step **2** : creating a `template.yaml` file. 
This is a sample kubernetes deployment which was turned into a go template.
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
___

#### Step **3** :
Now run the following command to render the template:
```bash 
./gotlet --template ./examples/template.yaml --varsfile ./examples/variables.yaml 
```
___

#### Step **4** :
Voila, you should get the following output from the application, if you want, you can just apply it using kubectl with piping the output of gotlet into `| kubectl apply -f -`

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
___

## TODO:
- [X] Changing the Exit codes to the correct ones so that the application exits with the appropriate error code.
- [ ] Only import environment variables if flag is passed to the application, or a switch that turns it off, it shouldn't necessarily happen all the time.
- [ ] Github Actions pipeline to automate the build.
- [ ] Linux packages.
- [ ] Increase test coverage.
- [ ] Better YAML validation.
- [ ] Add long flag names for better readability, alias short flag names to them for better usability.
- [ ] CLI input validation and more comprehensive error handling.
- [ ] Adding subcommands for yaml/template format validation.
- [ ] Change the flag package to cobra for better subcommand support (e.g. the per-file validation).
- [ ] Add a mechanism that excludes sensitive environment variables such as AWS/GCP, etc keys from imports by default.
- [ ] Simple and efficient docker images with GCP SDK / AWS CLI installed.
- [ ] Also add a flag that prints the output to STDOUT.
- [ ] Convert the readme commandline usage section into a table.