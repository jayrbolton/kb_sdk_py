"""
A registry of file templates for generating an SDK module

Uses the Jinja2 templates found in ./templates
"""

import os
from jinja2 import PackageLoader, Environment


# Initialize the Jinja2 environment loader
_env = Environment(
    loader=PackageLoader('kb_sdk', 'initializer/templates'),
    keep_trailing_newline=True  # Required to have flake8 pass
)

# 'destination' is a path in the generated app where each file goes
templates = [
    {
        'template': _env.get_template('kbase.yaml'),
        'destination': 'kbase.yaml'
    },
    {
        'template': _env.get_template('main.py.txt'),
        'destination': os.path.join('src', 'main.py')
    },
    {
        'template': _env.get_template('Dockerfile'),
        'destination': 'Dockerfile'
    },
    {
        'template': _env.get_template('kbase.yaml'),
        'destination': 'kbase.yaml'
    },
    {
        'template': _env.get_template('gitignore.txt'),
        'destination': '.gitignore'
    },
    {
        'template': _env.get_template('LICENSE.txt'),
        'destination': 'LICENSE.txt'
    },
    {
        'template': _env.get_template('README.md'),
        'destination': 'README.md'
    },
    {
        'template': _env.get_template('test_main.py.txt'),
        'destination': os.path.join('test', 'test_main.py')
    }
]
