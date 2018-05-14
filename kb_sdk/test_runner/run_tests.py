"""
Run the test suite for an app using unittest

This also boots the app's docker container
"""

import os
import time
import shutil
import subprocess
import docker

from kb_sdk.logger import logger


def run_tests(config, env, options):
    """
    :param config: A dict of configuration data from kbase.yaml
    :param env: Environment variables from os.environ and dotfile
    :param options: Command-line arguments (see cli.py)
    """
    start_time = time.time()
    # Check for the `docker` executable
    if shutil.which('docker') is None:
        logger.error('The `docker` command was not found. Install docker before running tests.')
        logger.error('  Instructions here: https://docs.docker.com/install/')
        exit(1)
    _build_docker_image(config, options)
    _run_unittest(config, env, options)
    end_time = time.time() - start_time
    logger.debug('Ran tests in ' + str(end_time) + ' seconds')


def _build_docker_image(config, options):
    """
    Build the docker image, if necessary. We build it if any of these conditions are met:
    - The image does not exist yet
    - The Dockerfile has been modified
    - The --build flag has been passed to kb-sdk-py
    """
    client = docker.from_env()
    image_name = config['docker_image_name']
    image = None
    try:
        image = client.images.get(image_name)
    except docker.errors.ImageNotFound as err:
        logger.debug('Docker image not found')
    modified = _is_dockerfile_modified()
    logger.debug('The Dockerfile has ' + ('not ' if not modified else '') + 'been modified')
    if modified or not image or options['build']:
        logger.debug('Building docker image with name ' + image_name)
        # It's easier to log a subprocess command rather than using the docker-py build function
        args = ['docker', 'build', '.', '--tag', image_name]
        proc = subprocess.Popen(args, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        for line in proc.stdout:
            logger.debug(line.decode('utf-8').rstrip())
        proc.wait()


def _is_dockerfile_modified():
    """
    Check if the Dockerfile has been modified since the previous run
    We serialize the modification time to build/.dockerfile_modified
    :return: boolean whether it has been modified or not
    """
    dockerfile_modified = os.path.getmtime('Dockerfile')
    modified = False
    modified_file_path = 'build/.dockerfile_modified'
    if os.path.isfile(modified_file_path):
        with open(modified_file_path, 'r') as f:
            contents = f.read()
            try:
                prev_modified = float(contents)
            except ValueError as err:
                prev_modified = 0
            modified = dockerfile_modified > prev_modified
    with open(modified_file_path, 'w') as f:
        f.write(str(dockerfile_modified))
    return modified


def _run_unittest(config, env, options):
    logger.debug('Calling unit tests')
    image_name = config['docker_image_name']
    # test_dir = config['paths']['test']
    # test_command = '"python -m unittest discover test"'
    test_command = 'python -m unittest discover test'
    if options.get('single_test'):
        test_command += ' ' + options['single_test']
    args = ['docker', 'run', image_name, 'bash', '-c', test_command]
    logger.debug('Running: ' + ' '.join(args))
    proc = subprocess.Popen(args, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    for line in proc.stdout:
        logger.info(line.decode('utf-8').rstrip())
    proc.wait()
