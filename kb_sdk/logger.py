"""
Logger
======

A customized logger with levels, timestamping, formatting

Levels include: CRITICAL, ERROR, WARNING, INFO, DEBUG, NOTSET

See: https://docs.python.org/3/library/logging.html

Usage:
    from kb_sdk.logger import logger

    logger.info("My log message")
    logger.debug("My debugging info")
"""

import logging
from logging.handlers import RotatingFileHandler
import os
import coloredlogs
from dotenv import load_dotenv

load_dotenv(dotenv_path='./.env')

# Log level defaults to INFO and can be set with the env var "LOG_LEVEL"
level = os.getenv('LOG_LEVEL', logging.INFO)

# Example: INFO     : Running `kb-sdk dev`
formatter = '%(levelname)-9s: %(message)s'

# Instantiate the logger with the handler and level
logger = logging.getLogger('kb-sdk')
coloredlogs.install(logger=logger, fmt=formatter, level=level)
# Fetch the handler we just created
main_handler = logger.handlers[-1]


# Create the file logger
if not os.path.isdir('log'):
    if os.path.isfile('log'):
        logger.error('log/ should be a directory; it is a file')
        exit(1)
    os.makedirs('log')
file_formatter = logging.Formatter('%(asctime)s - %(levelname)-9s: %(message)s')
log_path = os.path.join('log', 'debug.log')
# maxBytes is 1MB
file_handler = RotatingFileHandler(log_path, maxBytes=1000000, backupCount=1)
file_handler.setLevel(logging.DEBUG)
file_handler.setFormatter(file_formatter)
logger.addHandler(file_handler)
