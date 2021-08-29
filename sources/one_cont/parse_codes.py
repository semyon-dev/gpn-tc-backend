from bs4 import BeautifulSoup
import requests
import re
import time
import json

main_url = 'https://www.1cont.ru'

codes_path = "/contragent/by-okved"

headers = {'User-Agent': 'Mozilla/5.0'}

def deleteSpaces(text):
    text = re.sub("\\n +", '', text)
    text = re.sub("\\n +", '', text)
    text = re.sub("  +", '', text)
    return text

def search_codes_by_code(page):
    html_doc = requests.get(main_url+page, headers=headers)
    print("Статус код: ", html_doc.status_code)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    okveds = soup.find_all("div", class_="code-item")
    okveds_array = []
    if len(okveds) > 0:
        for okved in okveds:
            num = okved.find("span", class_="code").text
            # num = okved.text[:2]
            okved_info = okved.find("span",class_="name")
            link =""
            name=""
            if okved_info:
                if okved_info.find("a", class_="contragent-link"):
                    name = deleteSpaces(okved_info.find("a", class_="contragent-link").text)
                    link = okved_info.find("a", class_="contragent-link")["href"]
            if link == "" and name == "":
                continue
            okveds_array.append({
                "num":num,
                "link":link,
                "name":name,
            })
    else:   
        # companies not found
        return []
    # return result
    return okveds_array

# Search from all okveds with recursive and delay 5s
def search_codes_main_2(page):
    html_doc = requests.get(main_url+page, headers=headers)
    print("Статус код: ", html_doc.status_code)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    okveds = soup.find_all("li")
    okveds_array = []
    if len(okveds) > 0:
        for okved in okveds:
            # num = okved.find("div", class_="okved-company-line__num").text
            num = okved.text[:2]
            okved_info = okved.find("a",class_="contragent-link")
            link =""
            name=""
            if okved_info:
                name = deleteSpaces(okved_info.text)
                link = okved_info["href"]
            if link == "" and name == "":
                continue
            okveds_array.append({
                "num":num,
                "link":link,
                "name":name,
            })

            time.sleep(5)
            okveds_by_code = search_codes_by_code(link)
            okveds_array.extend(okveds_by_code)

    else:   
        # companies not found
        return []
    # return result
    return okveds_array

# Search from first page with okved
def search_codes_main(page):
    html_doc = requests.get(main_url+page, headers=headers)
    # print("Статус код: ", html_doc.status_code)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    okveds = soup.find_all("li")
    okveds_array = []
    if len(okveds) > 0:
        for okved in okveds:
            # num = okved.find("div", class_="okved-company-line__num").text
            num = okved.text[:2]
            okved_info = okved.find("a",class_="contragent-link")
            link =""
            name=""
            if okved_info:
                name = deleteSpaces(okved_info.text)
                link = okved_info["href"]
            if link == "" and name == "":
                continue
            okveds_array.append({
                "num":num,
                "link":link,
                "name":name,
            })
    else:   
        # companies not found
        return []
    # return result
    return okveds_array

okveds = search_codes_main_2(codes_path)
with open("base_codes.json", 'w', encoding='utf-8') as f:
    json.dump(okveds, f, ensure_ascii=False, indent=4)