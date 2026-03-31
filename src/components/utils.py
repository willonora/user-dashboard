import os
import logging
import json
from typing import List, Tuple

def get_config_file_path() -> str:
    return os.path.join(os.path.dirname(__file__), '../config/config.json')

def get_config() -> dict:
    config_file_path = get_config_file_path()
    try:
        with open(config_file_path, 'r') as config_file:
            return json.load(config_file)
    except json.JSONDecodeError as e:
        logging.getLogger('config').critical(f"Failed to load config: {e}")
        raise

def get_logger(name: str, level: int = logging.INFO) -> logging.Logger:
    logger = logging.getLogger(name)
    logger.setLevel(level)
    return logger

def get_user_data_path() -> str:
    return os.path.join(os.path.expanduser('~'), '.user-dashboard')

def get_user_data_file_path(user_id: str) -> str:
    user_data_path = get_user_data_path()
    return os.path.join(user_data_path, f'{user_id}.json')

def get_user_data(user_id: str) -> dict:
    user_data_file_path = get_user_data_file_path(user_id)
    if not os.path.exists(user_data_file_path):
        return {}
    try:
        with open(user_data_file_path, 'r') as user_data_file:
            return json.load(user_data_file)
    except json.JSONDecodeError as e:
        logging.getLogger('user_data').critical(f"Failed to load user data for {user_id}: {e}")
        return {}

def save_user_data(user_id: str, data: dict) -> None:
    user_data_file_path = get_user_data_file_path(user_id)
    try:
        with open(user_data_file_path, 'w') as user_data_file:
            json.dump(data, user_data_file)
    except Exception as e:
        logging.getLogger('user_data').critical(f"Failed to save user data for {user_id}: {e}")