from setuptools import setup, find_packages

setup(
    name='GitNoob',
    version='0.1.0',
    description='A library to simplify Git commands for new developers.',
    author='Aman Mehta',
    author_email='amehta20@nyit.edu',
    url='https://github.com/yourusername/GitNoob',  # Update with your repository link
    packages=find_packages(where='src'),
    package_dir={'': 'src'},
    entry_points={
        'console_scripts': [
            'lazypush=gitnoob.lazypush:main',  # Link the `lazypush` command to the `main` function in `lazypush.py`
        ],
    },
    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)
