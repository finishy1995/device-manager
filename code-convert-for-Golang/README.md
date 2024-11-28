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

## Environment settings

The Environment settings are stored in a `.env` file with the following format:

```.env
AZURE_OPENAI_KEY="<your_api_key>"
AZURE_OPENAI_ENDPOINT="https://<yourcompany>.openai.azure.com/"
DEPLOYMENT_NAME="<prefix>-gpt4-32k"
LLM_TEMPERATURE=0
LLM_TOP_P=0
NEWPATH="new_code"
```

A sample `.env` file is included at the project's root as `sample.env`. Duplicate this file, save it as `.env`, and adjust the variables as per your setup.

## How to Run

```bash
python3 -m flask run
```

**Note:** If you use virtual environment, Make sure you are using .venv/bin/python3 to run above command by check with `which python3` command , otherwise you will ocur "No module named xxx" error
