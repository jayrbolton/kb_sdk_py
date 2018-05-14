"""
Validate various aspects of an app, including configuration, flake8 linting, and some static code analysis

Configuration is set in a YAML file but converted to a python dict before getting passed here.
"""

import subprocess
from cerberus import Validator

from kb_sdk.config_validation.validate_main_methods import validate_main_methods
from kb_sdk.logger import logger


def validate(config):
    """ Validate the kbase.yaml config """
    validator = Validator(main_schema)
    validator.validate(config)
    if validator.errors:
        logger.error('-------------------------------------------')
        logger.error('Config file validation errors on kbase.yaml')
        logger.error('-------------------------------------------')
        _log_errors(validator.errors)
        exit(1)
    validate_main_methods(config)
    _run_flake8()
    logger.info('Congrats! Everything looks valid.')


def _run_flake8():
    """ Run the Flake8 validator on the app's codebase. Only shows warnings. """
    logger.debug('Running flake8 validation')
    args = ['python', '-m', 'flake8']
    proc = subprocess.Popen(args, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    for line in proc.stdout:
        logger.warning('flake8: ' + line.decode('utf-8').rstrip())
    proc.wait()


def _log_errors(errors, indent=0):
    """ Print a nested dictionary of errors from Cerberus """
    spaces = indent * " "
    # Create a bulleted list of each cerberus error message
    for key, messages in errors.items():
        for msg in messages:
            if isinstance(msg, dict):
                logger.error(spaces + key + ":")
                _log_errors(msg, indent + 2)
            else:
                logger.error(spaces + key + ": " + msg)


# Cerberus Schemas

# Schema for an input to a direct or narrative method
method_input_schema = {
    'type': {
        'required': True,
        'type': 'string',
        'minlength': 1
    },
    'label': {
        'required': True,
        'type': 'string',
        'minlength': 1
    }
}

# Schema for a narrative or direct method
method_schema = {
    'input': {
        'type': 'dict',
        'required': True,
        'allow_unknown': True,
        'valueschema': {
            'type': 'dict',
            'schema': method_input_schema
        }
    }
}

module_schema = {
    'name': {
        'required': True,
        'type': 'string',
        'minlength': 1
        },
    'description': {
        'required': True,
        'type': 'string',
        'minlength': 1
        },
    'version': {
        'required': True,
        'type': 'string',
        'regex': '^([0-9]+)\.([0-9]+)\.([0-9]+)$',
        'minlength': 1
        },
    'authors': {
        'type': 'list',
        'minlength': 1,
        'schema': {
            'type': 'string'
            }
        }
}

# Top-level schema for kbase.yaml
main_schema = {
    'module': {
        'type': 'dict',
        'schema': module_schema
    },
    'narrative_methods': {
        'type': 'dict',
        'required': False,
        'nullable': True,
        'allow_unknown': True,
        'valueschema': {
            'type': 'dict',
            'schema': method_schema
        }
    },
    'direct_methods': {
        'type': 'dict',
        'required': False,
        'nullable': True,
        'allow_unknown': True,
        'valueschema': {
            'type': 'dict',
            'schema': method_schema
        }
    },
    # Settings from cli.py
    'paths': {'type': 'dict', 'allow_unknown': True},
    'docker_image_name': {'type': 'string'}
}
