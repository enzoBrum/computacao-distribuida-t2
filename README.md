Para compilar  o projeto, use:
```bash
make compile-server
```

O projeto é organizado em uma arquitetura cliente-servidor, onde o cliente é uma aplicação de cadastro de alunos e o
servidor é um nó de um Cluster Raft que implementa uma máquina de estados replicada que serve como um espaço de tuplas.

Para rodar o servidor, use o comando:
```bash
./server_executable -nodeid x
```

onde `x` é um número no intervalo [1,5].
É necessário iniciar ao menos 3 servidores para que o sistema funcione corretamente.

Caso deseje iniciar todos os servidores ao mesmo tempo, pode use:
```bash
./server_executable -nodeid 1 &
./server_executable -nodeid 2 &
./server_executable -nodeid 3 &
./server_executable -nodeid 4 &
./server_executable -nodeid 5 &
```


Para iniciar o cliente, use:
```bash
python client.py
```
