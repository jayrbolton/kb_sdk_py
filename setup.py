from setuptools import setup, find_packages

setup(
    name='kb-sdk-py',
    description='KBase Software Development Kit',
    python_requires='>=3',
    version='0.0.1a5',
    packages=find_packages(),
    install_requires=[
        'docopt>=0.6.2',
        'pyyaml>=3.12',
        'flask>=1.0.2',
        'cerberus>=1.2',
        'python-dotenv>=0.8.2',
        'coloredlogs>=9.3.1',
        'docker>=3.3.0'
    ],
    package_data={
        'kb_sdk.initializer': ['templates/*.jinja2']
    },
    extras_require={
        'dev': [
            'flake8>=3.5.0'
        ]
    },
    entry_points={
        'console_scripts': [
            'kb-sdk-py = kb_sdk.cli.cli:main'
        ]
    },
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers'
    ]
)
