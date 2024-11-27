# Code Convert for Golang

## Installation Guide

**Prerequisites** : Ensure [Python 3](https://www.python.org/downloads/) and [PIP](https://pip.pypa.io/en/stable/) is installed on your system.

It's recommended to use a virtual environment. Follow these commands to set up the environment and install the necessary packages:

```bash
python3 -m venv .venv
source .venv/bin/activate
pip install --upgrade pip
pip install -r requirements.txt
```

If virtual environments are not considered an option we need to run the following commands instead:

```bash
pip install --upgrade pip
pip install -r requirements.txt
```

**Note:** Without virtual environements dependancy conflicts could occur.

## How to Run

```bash
python3 -m flask run
```

**Note:** Make sure you are using .venv/bin/python3 to run above command by check with `which python3` command , otherwise you will ocur "No module named xxx" error
