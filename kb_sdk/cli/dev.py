"""
Execute `kb-sdk dev`
(this is called from ./cli.py)

Run a flask-based simple development server
"""

import kb_sdk.dev_server.server as dev_server
import yaml
import os


def execute(args):
    # TODO make sure we are in app directory
    # TODO load the kbase.yaml configuration
    # TODO print the module name 
    if not os.path.isfile('./kbase.yaml'):
        raise ValueError(
            "Whoops, it doesn't look like we are in an SDK app directory"
        )
    config = {}
    with open('./kbase.yaml', 'r') as stream:
        try:
            config = yaml.load(stream)
        except yaml.YAMLError as err:
            print('Error loading config for app:', err)
    print('config', config)
    dev_server.app.config['ENV'] = 'development'
    dev_server.app.run()
    return 'dev'
