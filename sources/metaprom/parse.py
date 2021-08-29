from bs4 import BeautifulSoup
import requests
import re

main_url = 'https://metaprom.ru'

# Delete many spaces and \n symblos
def deleteSpaces(text):
    text = re.sub("\\n", '', text)
    text = re.sub("  +", '', text)
    # text = re.sub("\\n", '', text)
    return text

# This function decode email which protect cloudflare
def decodeEmail(e):
    de = ""
    k = int(e[:2], 16)
    for i in range(2, len(e)-1, 2):
        de += chr(int(e[i:i+2], 16)^k)
    return de

headers = {'User-Agent': 'Mozilla/5.0'}

def get_company_data(id):
    html_doc = requests.get(main_url+"/companies/"+id, headers=headers)
    html_doc.encoding = "cp1251"
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')

    tables = soup.find_all("tbody")
    index = 0 
    company_main_info = {}
    otraslNameTags = soup.find_all("li")
    otraslName = ""
    if len(otraslNameTags)>0:
        otraslName = otraslNameTags[0].find("a").text
    
    company_main_info["otraslName"] = otraslName
    for table in tables:
        if index == 0:
            datas = table.find_all("tr")
            company_info=[]
            for data in datas:
                name = deleteSpaces(data.find_all("td")[0].text)
                value = data.find_all("td")[1]
                if value.find("a"):
                    if name == "Регион":
                        company_info.append({
                            "name":name,
                            "value":value.find("a").text
                        })
                        continue
                    if value.find("span"):
                        # check if attribute find we decode email
                        if "data-cfemail" in value.find("span").attrs:
                            value = decodeEmail(value.find("span").attrs["data-cfemail"])
                    else:
                        value = value.find("a").text
                        
                else:
                    value = deleteSpaces(value.text)
                company_info.append({
                    "name":name,
                    "value":value
                })
            company_main_info["baseInfo"] = company_info
        index+=1
    return company_main_info

def getIDByLink(link:str)->str:
    data = link.split("/")
    id = ""
    if len(data) > 2:
        id = data[2]
    return id


def search_company(name, limit):
    html_doc = requests.get(main_url+"/search/?text="+name)
    html_doc.encoding = "cp1251"
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    companies = soup.find_all("div", class_="firm")
    companies_array = []
    if len(companies) > 0:
        total = 0
        for company in companies:
            if total == limit:
                break
            name = company.find('div', class_="firm_name").find("a")
            link = name["href"]
            descriptions = company.find_all("div", class_="smaller_txt")
            city=""
            type = ""
            allSees = ""
            if len(descriptions)>0:
                city = descriptions[0].text
            if len(descriptions)>1:
                type = descriptions[1].text
            if len(descriptions)>2:
                allSees = descriptions[2].text
            company_info = {
                "name":name.text,
                "id":getIDByLink(link),
                "city":city,
                "type":type,
                "allSees":allSees
            }
            companies_array.append(company_info)
            total+=1
    else:
        # companies not found
        return []
    # return result
    return companies_array
