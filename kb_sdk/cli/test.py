"""
Execute `kb-sdk test <module.Class.method>`

This is called from ./cli.py
"""

import subprocess
import os

import kb_sdk.cli.validate as validate
from kb_sdk.logger import logger


def execute(args, config, env):
    if not args.get('--skip-validation'):
        validate.execute(args, config, env)
    else:
        logger.debug('Skipping validation')
    logger.debug('Calling unit tests')
    module_option = args['<module.Class.method>']
    custom_env = os.environ.copy()
    custom_env['PYTHONPATH'] = 'src'
    test_dir = os.path.join(os.getcwd(), 'test')
    args = [
        'python', '-m', 'unittest', 'discover',
        test_dir
    ]
    if module_option:
        args.append(module_option)
    logger.debug('Running: ' + " ".join(args))
    proc = subprocess.Popen(args)
    proc.wait()
