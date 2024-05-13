import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

with open("requirements.txt", "r") as f:
    requirements = f.readlines()

setuptools.setup(
    name="HT",
    version="0.0.1",
    author="Krunoslav Lehman Pavasovic, Umut Simsekli",
    author_email="krunolp@gmail.com",
    description="Heavy Tailed Experiments",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url='github.com/krunolp/heavy_tails',
    package_dir={'heavy_tails': 'heavy_tails'},
    packages=setuptools.find_packages(),
    install_requires=[
        'numpy',
        'pandas',
        'scipy',
        'sklearn',
        'matplotlib',
        'jax',
        'jaxlib',
    ],
)
