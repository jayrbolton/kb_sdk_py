from setuptools import setup, find_packages

setup(
    name="kb-sdk-py",
    version="0.0.1",
    py_modules="kb_sdk",
    packages=find_packages(),
    install_requires=[
        'docopt',
        'pyyaml',
        'flask'
    ],
    entry_points={
        'console_scripts': [
            'kb-sdk-py = kb_sdk.cli.cli:main'
        ]
    }
)
