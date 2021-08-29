from bs4 import BeautifulSoup
import requests
import re
import unicodedata

main_url = 'https://www.1cont.ru'
okved_path = "/contragent/by-okved/"

# def get_company_data(url):
    
#     html_doc = requests.get(main_url+search_path+url)
#     f = html_doc.text
#     soup = BeautifulSoup(f, 'lxml')

#     name = soup.find("div", class_="company_name").find("a").text
#     logo = soup.find("a", class_="logo").find("img")["src"]
#     description = soup.find("div", class_="description").find_all("p")
#     # print(name)

#     # print("Адреса:")
#     addresses = soup.find_all("div", class_="address")
#     addresses_array = []
#     for address in addresses:
#         addresses_array.append(address.text)
#         # print(address.text)

#     # print("Контакты:")
#     contacts = soup.find_all("div", class_="contact")
#     contacts_array = []
#     for contact in contacts:
#         contact_info = {}
#         contact_type = contact.find("div", class_="type")
#         contact_info["type"] = contact_type.text
#         contact_value = contact.find("div", class_="value")
#         if contact_value.find("a"):
#             contact_info["link"]=contact_value.find("a")["href"]
#             contact_info["value"]=contact_value.find("a").text
#         else:
#             contact_info["link"] = ""
#             contact_info["value"]=contact_value.text
#         contacts_array.append(contact_info)
#         # print(contact_type.text, contact_value)

#     site = soup.find("div", class_="company_site").find("a")["href"]
#     # print("Сайт "+site.text)

#     # print("Скилы:")
#     skills = soup.find_all("a", class_="skill")
#     skills_array = []
#     for skill in skills:
#         skills_array.append(skill.text)
#         # print(skill.text)

#     # get emploee
#     # print("Get employees")
#     employees = soup.find_all("a", class_="company_public_member")
#     employees_array = []
#     for employee in employees:
#         employee_url = main_url+employee["href"]
#         avatar = employee.find("div", class_="avatar").find("img")["src"]
#         username = employee.find("div", class_="username").text
#         position = employee.find("div", class_="position").find("div", class_="text").text
#         employee_info = {
#             "employee_url":employee_url,
#             "avatar":avatar,
#             "username":username,
#             "position":position
#         }
#         employees_array.append(employee_info)

#     return {
#         "name": name,
#         "logo":logo,
#         "addresses": addresses_array,
#         "site": site,
#         "skills": skills_array,
#         "contacts":contacts_array,
#         "employees":employees_array,
#         "description":description
#     }

def deleteSpaces(text):
    text = re.sub("\\n +", '', text)
    text = re.sub("\\n +", '', text)
    return text

def search_companies(id,limit):
    # limit = 50
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

# companies = search_companies('remont-prochih-bytovyh-izdeliy-i-predmetov-lichnogo-poljhzovaniya_95_29_9',5)
# print(companies)