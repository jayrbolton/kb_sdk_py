"""
Execute `kb-sdk test <module.Class.method>`

This is called from ./cli.py
"""

import kb_sdk.cli.validate as validate
from kb_sdk.logger import logger
from kb_sdk.test_runner.run import run_tests


def execute(args, config, env):
    if not args.get('--skip-validation'):
        validate.execute(args, config, env)
    else:
        logger.debug('Skipping validation')
    logger.debug('Calling unit tests')
    module_option = args['<module.Class.method>']
    run_tests(config, env, module_option)
