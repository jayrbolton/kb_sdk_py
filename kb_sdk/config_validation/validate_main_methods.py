"""
Validate that the methods registered in kbase.yaml match those found in src/main.py
Also validate that the parameters for every method match a certain form
"""

import importlib.util as import_util
import inspect

from kb_sdk.logger import logger


def validate_main_methods(config):
    """
    Validate that the methods found in kbase.yaml are present in main.py
    Will log errors and exit if any validations fail
    :param config: the kbase.yaml config parsed into a dictionary
    """
    logger.debug('Validating methods in Main')
    module_path = config['paths']['main.py']
    spec = import_util.spec_from_file_location('main', module_path)
    main = import_util.module_from_spec(spec)
    spec.loader.exec_module(main)
    functions = inspect.getmembers(main.Main, predicate=inspect.isfunction)
    # Provide fallback dictionaries for the methods
    # We can't use config.get('narrative_methods', {}) as that key can be actually set to None
    narrative_methods = config.get('narrative_methods') or {}
    direct_methods = config.get('direct_methods') or {}
    # Keep track of all method names we find in main.py
    found_methods = {}
    for name, func in functions:
        args = inspect.getargspec(func).args
        if name is '__init__':
            if len(args) is not 2 or args[0] is not 'self' or args[1] is not 'context':
                logger.error('In src/main.py, __init__ constuctor:')
                _log_param_error('(self, context)', args)
                exit(1)
        elif name not in narrative_methods and name not in direct_methods:
            logger.error('Method in Main called "' + name + '" is not found in kbase.yaml')
            exit(1)
        else:
            found_methods[name] = True
            if len(args) is not 2 or args[0] is not 'self' or args[1] is not 'params':
                logger.error('In src/main.py, ' + name + ' method:')
                _log_param_error('(self, params)', args)
                exit(1)
    # Find all methods in config that are missing in Main
    for name in narrative_methods:
        if not found_methods.get(name):
            logger.error('Narrative method "' + name + '" is registered in kbase.yaml but not found in Main')
            exit(1)
    for name in direct_methods:
        if not found_methods.get(name):
            logger.error('Direct method "' + name + '" is registered in kbase.yaml but not found in Main')
            exit(1)


def _log_param_error(expected, args):
    """ Log an error where there is a mismatch between expected and actual parameters """
    logger.error('  Parameters should be ' + expected + '. Current params are ' + '(' + ', '.join(args) + ')')
