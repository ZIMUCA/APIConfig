import requests
import uuid

URL = "http://localhost:8080/agendas"

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