Use esses dois comandos.
```bash
make build # builda os contâineres
make server # levanta uma instância do espaço de tuplas
make client # levanta um client.
```

P.S: BFT-SNART parece ser um **pouco** mais tranquilo que jgroups, mas em compensação, é um saco de rodar já que tem que configurar tudo 
manualmente. Como o jgroups-raft faz praticamente a mesma coisa que ela, da pra continuar com ele

## TODO

- [x] espaço de tuplas
    - [x] obter tupla
    - [x] obter e remover tupla
    - [x] adicionar tupla
    - [x] obter todas as tuplas (pra debugar)
- [x] Apenas fazer bind no localhost (-Djgroups.bind_addr=127.0.0.1 -Djava.net.preferIPv4Stack=true)
- [x] Comunicação do cliente com o espaço de tuplas
