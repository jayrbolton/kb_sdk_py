"""
Module Initializer
====================

Generate a new SDK project, with initial directories and files.

This is run on the `kb-sdk init` CLI command
"""

import os
import string


def initializer(name, directory):
    if os.path.exists(directory):
        # TODO special error loggger
        print('-------------------------------------------------------')
        print('Oops! The path', os.path.abspath(name), 'already exists.')
        print("Please choose a directory that doesn't exist")
        print('-------------------------------------------------------')
        exit(1)
    # Create initial directories
    os.mkdir(directory)
    for dirname in ['src', 'build', 'assets', 'test']:
        os.mkdir(os.path.join(directory, dirname))
    # Write the default YAML configuration
    with open(os.path.join(directory, 'kbase.yaml'), 'w') as f:
        f.write(render_template(yaml_config_template, name=name))
    # Write the default KIDL spec
    with open(os.path.join(directory, 'kidl.spec'), 'w') as f:
        f.write(render_template(kidl_template, name=name))
    # Write the placeholder module file
    with open(os.path.join(directory, 'src', name + '.py'), 'w') as f:
        f.write(render_template(module_template, camel_name=name.title()))
    # Write the gitignore file
    with open(os.path.join(directory, '.gitignore'), 'w') as f:
        f.write(gitignore_template)
    return


def render_template(template, **kwargs):
    """ Render a string template given some parameters """
    s = string.Template(template)
    return s.substitute(**kwargs)


# String templates
# ================

module_template = """
class $camel_name(object):

    def __init__(context):
        self.ctx = context
        pass

    def my_method(params):
        return {
          'message': "My report message"
        }
"""

dockerfile_template = """
"""

kidl_template = """
module $name {

    typedef structure {
         string workspace_name;
    } Parameters;

    typedef structure {
        string report_name;
        string report_ref;
    } Results;

    funcdef method_name(Parameters params)
        returns (Results results) authentication required;
};
"""

yaml_config_template = """
module:
    name: $name
    description: Module description
    version: 0.0.1
    authors: []
"""

gitignore_template = """build
*.pyc
.env
"""
