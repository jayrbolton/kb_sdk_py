# KBase Software Development Kit

## Usage

_Install with Conda_

```sh
$ conda install kb-sdk-py -c kbase
```

_Run the CLI_

```sh
$ kb-sdk-py --help
```

_Start a new project_

```sh
$ kb-sdk-py init project_name [directory]
```

_Run the development server_

```sh
$ cd project_name
$ kb-sdk-py dev
```

## Development

Set up the environment and install dependencies with:

```sh
$ python3 -m venv env
$ source env/bin/activate
$ python3 -m pip install --editable .
```

Test that things are set up by running `kb-sdk-py`.

### Project anatomy

* `/kb_sdk`: Root package
* `/kb_sdk/cli`: Command line handler
* `/kb_sdk/dev_server`: Development server with flask
* `/kb_sdk/initializer`: Module initializer