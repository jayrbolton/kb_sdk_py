"""
Validate the configuration settings for an app.

Configuration is set in a YAML file but converted to a python dict before
getting passed here.
"""

import inspect
import importlib.util as import_util
import os
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
    _validate_main_methods(config)
    return validator


def _validate_main_methods(config):
    """ Validate that the methods found in kbase.yaml are present in main.py """
    module_path = os.path.join(os.getcwd(), 'src', 'main.py')
    spec = import_util.spec_from_file_location('main', module_path)
    main = import_util.module_from_spec(spec)
    spec.loader.exec_module(main)
    functions = inspect.getmembers(main.Main, predicate=inspect.isfunction)
    narratives = config.get('narrative_methods') or {}
    directs = config.get('direct_methods') or {}
    found_methods = {}
    for name, func in functions:
        if name == '__init__':
            args = inspect.getargspec(func)
            if args.args[0] is not 'self':
                logger.error('The first parameter to Main.' + name + ' should be "self"')
                exit(1)
            if args.args[1] is not 'context':
                logger.error('The first parameter to Main.' + name + ' should be "context"')
                exit(1)
        elif not narratives.get(name) and not directs.get(name):
            logger.error('Method "Main.' + name + '" is not found in kbase.yaml')
            exit(1)
        else:
            found_methods[name] = True
            args = inspect.getargspec(func)
            if args.args[0] is not 'self':
                logger.error('The first parameter to Main.' + name + ' should be "self"')
                exit(1)
            if args.args[1] is not 'params':
                logger.error('The second paramter to Main.' + name + ' should be "params"')
                exit(1)
    # Find all methods in config that are missing in Main
    for method in narratives:
        if not found_methods.get(method):
            logger.error('Narrative method "' + method + '" is registered in kbase.yaml but not found in Main')
            exit(1)
    for method in directs:
        if not found_methods.get(method):
            logger.error('Direct method "' + method + '" is registered in kbase.yaml but not found in Main')
            exit(1)


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
