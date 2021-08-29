from bs4 import BeautifulSoup
import requests
import re

main_url = 'https://www.rbc.ru'
search_path = "/search/?project=rbcnews&query="

def deleteSpaces(text):
    text = re.sub("\\n +", '', text)
    text = re.sub("\\n +", '', text)
    return text

def search_news(name,limit):
    html_doc = requests.get(main_url+search_path+name)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    news = soup.find_all("div", class_="search-item")
    news_array = []
    if len(news) > 0:
        total = 0
        for new in news:
            if total == limit:
                break
            new = new.find("div")
            articleName = new.find("span", class_="search-item__title").text
            text = new.find("span", class_="search-item__text").text
            text = deleteSpaces(text)
            link = new.find('a', class_="search-item__link")["href"]
            categoryAndDate = new.find('span',class_="search-item__category").text
            categoryAndDate = deleteSpaces(categoryAndDate)
            img = new.find('img', class_='search-item__image')
            if img:
                img = img["src"]
            else:
                img = ""
            news_array.append({
                "name":articleName,
                "text":text,
                "link":link,
                "category":categoryAndDate,
                "img":img,
            })
            total+=1
    else:
        # companies not found
        return []
    # return result
    return news_array