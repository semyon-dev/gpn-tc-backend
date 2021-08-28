from bs4 import BeautifulSoup
import requests
from parse import search_news 

def handler(event, context):
    # defaultLimit = 5
    limit = 5
    maxLimit = 25
    headers = {'User-Agent': 'Mozilla/5.0'}
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
        
    articles = search_news(name, int(limit))
    
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json'
        },
        'isBase64Encoded': False,
        'body': {
            "companies":articles,
        }
    }