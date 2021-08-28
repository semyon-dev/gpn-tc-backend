from bs4 import BeautifulSoup
import requests
import re

main_url = 'https://metaprom.ru'

def deleteSpaces(text):
    text = re.sub("\\n", '', text)
    text = re.sub("  +", '', text)
    # text = re.sub("\\n", '', text)
    return text

def decodeEmail(e):
    de = ""
    k = int(e[:2], 16)
    for i in range(2, len(e)-1, 2):
        de += chr(int(e[i:i+2], 16)^k)
    return de
headers = {'User-Agent': 'Mozilla/5.0'}
def get_company_data(id):
    # print("URL: "+main_url+"/companies/"+id)
    html_doc = requests.get(main_url+"/companies/"+id, headers=headers)
    html_doc.encoding = "cp1251"
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')

    tables = soup.find_all("tbody")
    index = 0 
    company_main_info = {}
    otraslNameTags = soup.find_all("li")
    # print(otraslNameTags)
    otraslName = ""
    if len(otraslNameTags)>0:
        otraslName = otraslNameTags[0].find("a").text
    
    company_main_info["otraslName"] = otraslName
    for table in tables:
        # table with base data
        if index == 0:
            datas = table.find_all("tr")
            company_info=[]
            for data in datas:
                name = deleteSpaces(data.find_all("td")[0].text)
                value = data.find_all("td")[1]
                if value.find("a"):
                    if name == "Регион":
                        # value = value.find("a").text
                        company_info.append({
                            "name":name,
                            "value":value.find("a").text
                        })
                        continue
                    if value.find("span"):
                        # value = value.find("span")["data-cfemail"]
                        if "data-cfemail" in value.find("span").attrs:
                            # print("Attrs:",decodeEmail(value.find("span").attrs["data-cfemail"]))
                            value = decodeEmail(value.find("span").attrs["data-cfemail"])
                    # print(decodeEmail(value.find("span")["data-cfemail"]))
                    else:
                        value = value.find("a").text
                        
                else:
                    value = deleteSpaces(value.text)
                company_info.append({
                    "name":name,
                    "value":value
                })
            company_main_info["baseInfo"] = company_info
        # if index == 1:
        #     # Объявления
        #     pass
        # if index == 2:
        #     # поставщики
        #     otraslNameTags = soup.find_all("h2")
        #     otraslName = ""
        #     if len(otraslName)>2:
        #         otraslName = otraslNameTags[2].text
        #     company_main_info["otraslname"] = otraslName
        #     datas = table.find_all("tr")
        #     suppliers=[]
        #     for data in datas:
        #         name = data.find_all("td")[0]
        #         print(name)
        #         name = name.find("b").find("a")
        #         city = data.find_all("td")[1].find("span").text
        #         description = data.find_all("td")[2].find("span").text
        #         suppliers.append({
        #             "name":name.text,
        #             "link":name["href"],
        #             "city":city,
        #             "description":description
        #         })
        #     company_main_info["suppliers"] = suppliers
        index+=1
    return company_main_info


    # logo = soup.find("a", class_="logo").find("img")["src"]
    # description = soup.find("div", class_="description").find_all("p")
    # # print(name)   

    # # print("Адреса:")
    # addresses = soup.find_all("div", class_="address")
    # addresses_array = []
    # for address in addresses:
    #     addresses_array.append(address.text)
    #     # print(address.text)

    # # print("Контакты:")
    # contacts = soup.find_all("div", class_="contact")
    # contacts_array = []
    # for contact in contacts:
    #     contact_info = {}
    #     contact_type = contact.find("div", class_="type")
    #     contact_info["type"] = contact_type.text
    #     contact_value = contact.find("div", class_="value")
    #     if contact_value.find("a"):
    #         contact_info["link"]=contact_value.find("a")["href"]
    #         contact_info["value"]=contact_value.find("a").text
    #     else:
    #         contact_info["link"] = ""
    #         contact_info["value"]=contact_value.text
    #     contacts_array.append(contact_info)
    #     # print(contact_type.text, contact_value)

    # site = soup.find("div", class_="company_site").find("a")["href"]
    # # print("Сайт "+site.text)

    # # print("Скилы:")
    # skills = soup.find_all("a", class_="skill")
    # skills_array = []
    # for skill in skills:
    #     skills_array.append(skill.text)
    #     # print(skill.text)

    # # get emploee
    # # print("Get employees")
    # employees = soup.find_all("a", class_="company_public_member")
    # employees_array = []
    # for employee in employees:
    #     employee_url = main_url+employee["href"]
    #     avatar = employee.find("div", class_="avatar").find("img")["src"]
    #     username = employee.find("div", class_="username").text
    #     position = employee.find("div", class_="position").find("div", class_="text").text
    #     employee_info = {
    #         "employee_url":employee_url,
    #         "avatar":avatar,
    #         "username":username,
    #         "position":position
    #     }
    #     employees_array.append(employee_info)

    # return {
    #     "name": name,
    #     "logo":logo,
    #     "addresses": addresses_array,
    #     "site": site,
    #     "skills": skills_array,
    #     "contacts":contacts_array,
    #     "employees":employees_array,
    #     "description":description
    # }


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
            # company_info = get_company_data(name["href"])
            # print(company_info)
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

# companies = search_company('техгрант',5)
# print(companies)
# company_data = get_company_data('id57145-tehgrant-ooo')
# print(company_data)
# [{'name': 'Техгрант, ООО', 'id': 'id57145-tehgrant-ooo', 'city': 'Нижний Новгород', 'type': 'Производство СОЖ марок "Акватек" и "Техмол" для металлообработки. Импортозамещение СОЖ.', 'allSees': '8355'}, {'name': 'Аверс, ООО', 'id': 'id577623-avers-ooo', 'city': 'Нижний Новгород', 'type': 'Реализует СОЖ "Акватек" и "Техмол" производства ООО "ТЕХГРАНТ". ИМПОРТОЗАМЕЩЕНИЕ СОЖ. Масла, смазки.', 'allSees': '6671'}, {'name': 'Техгрант, ООО', 'id': 'id48138-tehgrant-ooo', 'city': 'Чебоксары', 'type': 'Производитель СОЖ', 'allSees': '1115'}]
