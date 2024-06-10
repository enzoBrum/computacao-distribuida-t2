Use esses dois comandos.
```bash
make build # builda os contâineres
make server # levanta uma instância do espaço de tuplas
make client # levanta um client.
```

P.S: BFT-SNART parece ser um **pouco** mais tranquilo que jgroups, mas em compensação, é um saco de rodar já que tem que configurar tudo 
manualmente. Como o jgroups-raft faz praticamente a mesma coisa que ela, da pra continuar com ele

## TODO

- [x] Como rodar jgroups no docker?
- [ ] espaço de tuplas
    - [ ] obter tupla
    - [ ] obter e remover tupla
    - [ ] adicionar tupla
    - [ ] obter todas as tuplas (pra debugar)
    - [ ] usar uma Trie ao invés de uma ArrayList? (O pior caso ainda é n^2, mas o resto tende a ser melhor.)
- [ ] Comunicação do cliente com o espaço de tuplas
- [ ] script de testes
    - [ ] adicionar servidor ao espaço de tuplas
    - [ ] remover servidor ao espaço de tuplas
    - [ ] simular falha por colapso
