"""
Execute `kb-sdk init`
(this is called from ./cli.py)
"""

from kb_sdk.initializer.initializer import initializer


def execute(args):
    name = args['<name>']
    directory = args['<directory>'] or name
    print('Initializing a new SDK project with name', name, '...')
    print('Name:', name)
    print('Directory:', directory)
    print('Checking uniqueness of name...')
    print('Checking for your developer token...')
    initializer(name, directory)
    return 'init'
