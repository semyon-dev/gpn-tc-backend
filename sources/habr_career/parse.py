from bs4 import BeautifulSoup
import requests

main_url = 'https://career.habr.com'

def get_company_data(url):
    
    html_doc = requests.get(main_url+url)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')

    name = soup.find("div", class_="company_name").find("a").text
    logo = soup.find("a", class_="logo").find("img")["src"]
    description = soup.find("div", class_="description").find_all("p")
    # print(name)

    # print("Адреса:")
    addresses = soup.find_all("div", class_="address")
    addresses_array = []
    for address in addresses:
        addresses_array.append(address.text)
        # print(address.text)

    # print("Контакты:")
    contacts = soup.find_all("div", class_="contact")
    contacts_array = []
    for contact in contacts:
        contact_info = {}
        contact_type = contact.find("div", class_="type")
        contact_info["type"] = contact_type.text
        contact_value = contact.find("div", class_="value")
        if contact_value.find("a"):
            contact_info["link"]=contact_value.find("a")["href"]
            contact_info["value"]=contact_value.find("a").text
        else:
            contact_info["link"] = ""
            contact_info["value"]=contact_value.text
        contacts_array.append(contact_info)
        # print(contact_type.text, contact_value)

    site = soup.find("div", class_="company_site").find("a")["href"]
    # print("Сайт "+site.text)

    # print("Скилы:")
    skills = soup.find_all("a", class_="skill")
    skills_array = []
    for skill in skills:
        skills_array.append(skill.text)
        # print(skill.text)

    # get emploee
    # print("Get employees")
    employees = soup.find_all("a", class_="company_public_member")
    employees_array = []
    for employee in employees:
        employee_url = main_url+employee["href"]
        avatar = employee.find("div", class_="avatar").find("img")["src"]
        username = employee.find("div", class_="username").text
        position = employee.find("div", class_="position").find("div", class_="text").text
        employee_info = {
            "employee_url":employee_url,
            "avatar":avatar,
            "username":username,
            "position":position
        }
        employees_array.append(employee_info)

    return {
        "name": name,
        "logo":logo,
        "addresses": addresses_array,
        "site": site,
        "skills": skills_array,
        "contacts":contacts_array,
        "employees":employees_array,
        "description":description
    }



def search_company(name, limit):
    html_doc = requests.get(main_url+"/companies?q="+name)
    f = html_doc.text
    soup = BeautifulSoup(f, 'lxml')
    companies = soup.find_all("div", class_="companies-item")
    companies_array = []
    if len(companies) > 0:
        total = 0
        for company in companies:
            if total == limit:
                break
            name = company.find('div', class_="companies-item-name").find("a", class_="title")
            # print(name["href"])
            # print(name.text)
            # TODO: check if not found
            company_info = get_company_data(name["href"])
            # print(company_info)
            companies_array.append(company_info)
            total+=1
    else:
        # companies not found
        return []
    # return result
    return companies_array

# companies = search_company('сбер',5)
# print(companies)