from setuptools import setup, find_packages

setup(
    name="kb-sdk-py",
    python_requires='>=3',
    version="0.0.1",
    py_modules="kb_sdk",
    packages=find_packages(),
    install_requires=[
        'docopt>=0.6.2',
        'pyyaml>=3.12',
        'flask>=1.0.2',
        'cerberus>=1.2',
        'python-dotenv>=0.8.2',
        'coloredlogs>=9.3.1'
    ],
    entry_points={
        'console_scripts': [
            'kb-sdk-py = kb_sdk.cli.cli:main'
        ]
    }
)
