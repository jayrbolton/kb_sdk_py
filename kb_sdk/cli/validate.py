"""
Execute `kb-sdk test <module.Class.method>`

This is called from ./cli.py
"""

from kb_sdk.config_validation.validator import validate
from kb_sdk.logger import logger


def execute(args, config, env):
    logger.debug('Validating configuration')
    validate(config)  # Will throw and exit if invalid
    logger.info('✓ Configuration is valid.')
    if not env.get('KBASE_USERNAME'):
        logger.error('Set your KBASE_USERNAME environment variable')
        exit(1)
    if not env.get('KBASE_TOKEN'):
        logger.error('Set your KBASE_TOKEN environment variable')
        exit(1)
    logger.info('✓ Env vars are set.')
