from bs4 import BeautifulSoup
import requests
from parse import search_company, get_company_data 

def handler(event, context):
    limit = 5
    maxLimit = 25
    headers = {'User-Agent': 'Mozilla/5.0'}
    # Get name by q
    name = ''
    if 'queryStringParameters' in event and 'q' in event['queryStringParameters']:
        name = event['queryStringParameters']['q']
    # Get company id by id
    id = ''
    if 'queryStringParameters' in event and 'id' in event['queryStringParameters']:
            id = event['queryStringParameters']['id']
    if name == '' and id == '':
        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'isBase64Encoded': False,
            'body': {
                "error":True,
                "message":"set q param or id"
            }
        }
    if name != '' and id != '':
        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'isBase64Encoded': False,
            'body': {
                "error":True,
                "message":"set only q param or id"
            }
        }
    # Get limit
    if 'queryStringParameters' in event and 'limit' in event['queryStringParameters']:
        limit = event['queryStringParameters']['limit']
        # Convert limit to int or send error
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
    # If id set we find by id
    if id != '':
        company = get_company_data(id)
        return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json'
        },
        'isBase64Encoded': False,
        'body': {
            "company":company,
        }
    }
    # Else we find companies by name
    companies = search_company(name, int(limit))
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