# <tuple-space-address>/add
# <tuple-space-address>/get
# <tuple-space-address>/read
# <tuple-space-address>/all
# <tuple-space-address>/
import json

import requests

def run():    

    server_adress = choose_server()
    
    while True:
        print("1 - ADICIONAR")
        print("2 - LER")
        print("3 - REMOVER")
        print("4 - LER TODAS")
        print("5 - ALTERAR ENDEREÇO DO SERVIDOR")
        print("6 - SAIR")

        choose = input()

        match choose:
            case '1':
                response = requests.post(url=f"{server_adress}/add", data=json.dumps(ask_input(True)))
            case '2':
                response = requests.post(url=f"{server_adress}/read", data=json.dumps(ask_input(False)))
                data = response_status(response)
                print("="*10)
                if data:
                    print(f"Aluno recebido: {response.json()}")
                else:
                    print("Aluno não encontrado")
            case '3':
                response = requests.get(url=f"{server_adress}/get", data=json.dumps(ask_input(False)))
                data = response_status(response)
                print("="*10)
                if data:
                    print(f"Aluno removido: {response.json()}")
                else:
                    print("Aluno não encontrado")
            case '4':
                response = requests.get(url=f"{server_adress}/all")
                data = response_status(response)
                print("="*10)
                if data:
                    print(f"Lista de alunos: {response.json()}")
                else:
                    print("Não há tuplas no servidor")
            case '5':
                server_adress = choose_server()
            case '6':
                break
            case _:
                print("Opção inválida.")

        print("="*10)

def response_status(response: int) -> str | None:
    if response.status_code == 200:
        return response.json()

    print(f"Falha na requisição. Status code: {response.status_code}")
    return None

def choose_server() -> str:
    while True:
        server_adress = input("ENDEREÇO DO SERVIDOR: ")
        if not server_adress.startswith("http://"):
            server_adress = f"http://{server_adress}"
        ping = requests.get(url=f"{server_adress}/").status_code
        
        if ping == 200:
            break
        print('Endereço inválido.')

    return server_adress

def ask_input(is_add: bool) -> list:
    if is_add:
        subject = input_with_empty_check("Nome da disciplina: ")
        enrollment = input_with_empty_check("Matrícula do(a) estudante: ")
        name = input_with_empty_check("Nome do(a) estudante: ")
        attendance = input_with_empty_check("Frequência: ")
        average = input_with_empty_check("Média do estudante: ")
    else:
        subject = input("Nome da disciplina: ")
        enrollment = input("Matrícula do(a) estudante: ")
        name = input("Nome do(a) estudante: ")
        attendance = input("Frequência: ")
        average = input("Média do estudante: ")
    data_list = [subject, enrollment, name, attendance, average]
    for i in data_list:
        if i == "" or '*' in i:
            i = '*'
    return data_list
    
    """
DISCIPLINA
MATRíCULA
NOME
FREQUÊNCIA
MÉDIA
"""
def input_with_empty_check(msg: str) -> str:
    result = ''
    while result == '':
        input_string = input(msg)
        if len(input_string) > 0 and '*' not in input_string:
            result = input_string
        else:
            print("A entrada não pode ser string vazia ou conter *")
            
    return result

if __name__ == '__main__':
    run()
