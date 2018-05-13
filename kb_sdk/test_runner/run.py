"""
Run the test suite for an app using unittest

This also boots the app's docker container
"""

import subprocess
import os

from kb_sdk.logger import logger


def run_tests(config, env, module_option):
    """
    :param config: A dict of configuration data from kbase.yaml
    :param env: Environment variables from os.environ and dotfile
    :param module_option: The arg in `kb-sdk-py test <module.Class.method>`
        specifying a single module/class/method to run tests on
    """
    custom_env = env.copy()
    custom_env['PYTHONPATH'] = 'src'
    test_dir = os.path.join(os.getcwd(), 'test')
    args = [
        'python', '-m', 'unittest', 'discover',
        test_dir
    ]
    if module_option:
        args.append(module_option)
    logger.debug('Running: ' + " ".join(args))
    proc = subprocess.Popen(args, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    for line in proc.stdout:
        logger.info(line.decode('utf-8').rstrip())
    proc.wait()
