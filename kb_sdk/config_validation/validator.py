"""
Validate the configuration settings for an app.

Configuration is set in a YAML file but converted to a python dict before
getting passed here.
"""

from cerberus import Validator

from kb_sdk.logger import logger


def validate(config):
    """ Validate the kbase.yaml config """
    schema = {
        'module': {
            'type': 'dict',
            'schema': module_schema
        },
        'narrative_methods': {
            'type': 'dict',
            'required': True,
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
            # TODO create schema
        }
    }
    validator = Validator(schema)
    validator.validate(config)
    if validator.errors:
        logger.error('-------------------------------------------')
        logger.error('Config file validation errors on kbase.yaml')
        logger.error('-------------------------------------------')
        _log_errors(validator.errors)
        exit(1)
    return validator


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


# Schemas

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
