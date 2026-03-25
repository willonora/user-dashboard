# utils.py

import os
import logging
from typing import List, Tuple

def get_config_file_path() -> str:
    """Returns the path to the project's configuration file."""
    return os.path.join(os.path.dirname(__file__), '../config/config.json')

def get_config() -> dict:
    """Returns the project's configuration as a dictionary."""
    config_file_path = get_config_file_path()
    with open(config_file_path, 'r') as config_file:
        return json.load(config_file)

def get_logger(name: str, level: int = logging.INFO) -> logging.Logger:
    """Returns a logger instance with the specified name and level."""
    logger = logging.getLogger(name)
    logger.setLevel(level)
    return logger

def get_user_data_path() -> str:
    """Returns the path to the user's data directory."""
    return os.path.join(os.path.expanduser('~'), '.user-dashboard')

def get_user_data_file_path(user_id: str) -> str:
    """Returns the path to the user's data file."""
    user_data_path = get_user_data_path()
    return os.path.join(user_data_path, f'{user_id}.json')

def get_user_data(user_id: str) -> dict:
    """Returns the user's data as a dictionary."""
    user_data_file_path = get_user_data_file_path(user_id)
    if not os.path.exists(user_data_file_path):
        return {}
    with open(user_data_file_path, 'r') as user_data_file:
        return json.load(user_data_file)

def save_user_data(user_id: str, data: dict) -> None:
    """Saves the user's data to the user's data file."""
    user_data_file_path = get_user_data_file_path(user_id)
    with open(user_data_file_path, 'w') as user_data_file:
        json.dump(data, user_data_file)

import json