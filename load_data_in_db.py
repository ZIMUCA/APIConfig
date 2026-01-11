import requests
import uuid

URL = "http://localhost:8081/agendas"

data = [
    ("13295", "M1 Groupe 1 langue"),
    ("13345", "M1 Groupe 2 langue"),
    ("13397", "M1 Groupe 3 langue"),
    ("7224",  "M1 Groupe 1 option"),
    ("7225",  "M1 Groupe 2 option"),
    ("62962", "M1 Groupe 3 option"),
    ("62090", "M1 Groupe option"),
    ("56529", "M1 - Tutorat L2"),
]

headers = {
    "Content-Type": "application/json"
}

for agenda_id, name in data:
    payload = {
        "id": str(uuid.uuid4()),
        "agenda_id": agenda_id,
        "name": name
    }

    response = requests.post(URL, json=payload, headers=headers)

URL = "http://localhost:8081/alerts"

dataAlerts = [
    ("13295", "Maxime.VIMPERE@etu.uca.fr"),
    ("13345", "Maxime.VIMPERE@etu.uca.fr"),
    ("13397", "Maxime.VIMPERE@etu.uca.fr"),
    ("7224",  "Maxime.VIMPERE@etu.uca.fr"),
    ("7225",  "Maxime.VIMPERE@etu.uca.fr"),
    ("62962", "Maxime.VIMPERE@etu.uca.fr"),
    ("62090", "Maxime.VIMPERE@etu.uca.fr"),
    ("56529", "Maxime.VIMPERE@etu.uca.fr"),
]


for agenda_id, mail in dataAlerts:
    payload = {
        "id": str(uuid.uuid4()),
        "agenda_id": agenda_id,
        "mail": mail
    }

    response = requests.post(URL, json=payload, headers=headers)