"""
Execute `kb-sdk init`
(this is called from ./cli.py)
"""

from kb_sdk.initializer.initializer import initializer
from kb_sdk.logger import logger


def execute(args, env):
    name = args['<name>']
    directory = args['<directory>'] or name
    logger.info('Initializing "' + name + '" into ' + directory)
    logger.info('Checking uniqueness of name...')
    logger.info('Checking for your developer token...')
    initializer(name, directory, env)
    return 'init'
