from bs4 import BeautifulSoup
import requests
import re
import unicodedata

main_url = 'https://www.1cont.ru'
okved_path = "/contragent/by-okved/"

def deleteSpaces(text):
    text = re.sub("\\n +", '', text)
    text = re.sub("\\n +", '', text)
    return text

def search_companies(id,limit):
    html_doc = requests.get(main_url+okved_path+id)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    companies = soup.find_all("div", class_="tr tbody-tr")
    companies_array = []
    if len(companies) > 0:
        total = 0
        for company in companies:
            if total == limit:
                break
            all_info = company.find_all("div", class_="td")
            block_index = 0
            data_info = {}
            data_info["other"] = []
            for block_info in all_info:
                if block_index == 0:
                    name = block_info.find("a")
                    link = ""
                    if name:
                        link = name["href"]
                        name = name.text
                        data_info["name"]=name
                        data_info["link"]=link
                else:
                    block_name = unicodedata.normalize("NFKD", block_info.find("div", class_="td__caption").text) 
                    block_value = block_info.find("div", class_="td__text").text
                    data_info["other"].append({
                        "name":block_name,
                        "value":block_value,
                    })
                block_index += 1
            
            companies_array.append(data_info)
            total+=1
    else:
        # companies not found
        return []
    # return result
    return companies_array
