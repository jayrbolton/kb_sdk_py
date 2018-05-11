"""
Module Initializer
====================

Generate a new SDK project, with initial directories and files.

This is run with the `kb-sdk init` CLI command
"""

import os
from datetime import datetime

from kb_sdk.initializer.file_templates import templates
from kb_sdk.logger import logger


def initializer(module_name, directory, env):
    if os.path.exists(directory):
        # TODO special error loggger
        logger.error('The path ' + os.path.abspath(directory) + ' already exists. ')
        exit(1)
    # Create initial directories
    os.mkdir(directory)
    for dirname in ['src', 'build', 'assets', 'test', 'log']:
        os.mkdir(os.path.join(directory, dirname))
    # Write some blank __init__.py files
    with open(os.path.join(directory, 'src', '__init__.py'), 'w') as f:
        f.write('')
    with open(os.path.join(directory, 'test', '__init__.py'), 'w') as f:
        f.write('')
    # Common data used in any of the templates
    template_data = {
        'module_name': module_name,
        'username': env.get('KBASE_USERNAME', ''),
        'year': datetime.now().year
    }
    for tmpl in templates:
        with open(os.path.join(directory, tmpl['destination']), 'w') as f:
            f.write(tmpl['template'].render(**template_data))
