"""
Execute `kb-sdk test <module.Class.method>`
(this is called from ./cli.py)
"""


def execute(args):
    print('Calling unit tests...')
    module_option = args['<module.Class.method>']
    if module_option:
        print('Calling specific module, class, and method...')
    return 'test'
