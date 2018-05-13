# KBase Software Development Kit

## Usage

_Install_ (TODO)

_Run the CLI_

```sh
$ kb-sdk-py --help
```

_Start a new project_

```sh
$ kb-sdk-py init project_name [directory]
```

_Run the status server_

```sh
$ cd project_name
$ kb-sdk-py status
```

_Run tests_

```sh
$ kb-sdk-py test
```

### Logging

Logs have levels: CRITICAL, ERROR, WARNING, INFO, or DEBUG. You can set the level for commands you run by using the `LOG_LEVEL` environment variable. For example:

```sh
# Will only show error messages
$ LOG_LEVEL=error kb-sdk-py validate
```

All logs for the `kb-sdk-py` are saved into `/log/debug.log` at the DEBUG level (this file should be git-ignored). The log file has a size limit of 1MB and will start rotating through a couple backups (`debug.log.1` and `debug.log2.`) when it gets full.

### dotenv

You can set project-scoped environment variables by adding them to a `.env` file in your app repo. This file should be git-ignored. The format of the `.env` file is as follows:

```
KBASE_USERNAME=my_name
KBASE_TOKEN=my_dev_auth_token
LOG_LEVEL=debug
```

## Development

Set up the environment and install dependencies with:

```sh
$ python3 -m venv env
$ source env/bin/activate
$ pip install --editable .
```

Test that things are set up by running `kb-sdk-py`.

### Project anatomy

* `/kb_sdk`: Root package
* `/kb_sdk/cli`: Command line handler
* `/kb_sdk/dev_server`: Development server with flask
* `/kb_sdk/initializer`: Module initializer
* `/kb_sdk/config_validation`: Validate config found in kbase.yaml
* `/kb_sdk/param_validation`: Validate parameters passed to Main
