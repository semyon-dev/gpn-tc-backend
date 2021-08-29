from bs4 import BeautifulSoup
import requests
from parse import search_companies


def handler(event, context):
    limit = 5
    maxLimit = 50
    # Get name by q
    name = ''
    if 'queryStringParameters' in event and 'q' in event['queryStringParameters']:
        name = event['queryStringParameters']['q']
    if name == '':
        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'isBase64Encoded': False,
            'body': {
                "error":True,
                "message":"set q param"
            }
        }
    # Get limit
    if 'queryStringParameters' in event and 'limit' in event['queryStringParameters']:
        limit = event['queryStringParameters']['limit']
    # Convert limit to int
    try:
        limit = int(limit)
    except ValueError:
        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'isBase64Encoded': False,
            'body': {
                "error":True,
                "message":"limit must be int"
            }
        }
    # Check limit on max limit
    if limit > maxLimit:
        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'isBase64Encoded': False,
            'body': {
                "error":True,
                "message":"limit is more than max"
            }
        }
        
    companies = search_companies(name, limit)
    
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json'
        },
        'isBase64Encoded': False,
        'body': {
            "companies":companies,
        }
    }