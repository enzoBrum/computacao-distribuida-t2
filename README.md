Use esses dois comandos.
```bash
docker build -t distribuida-tuple-space .
docker run --rm -it distribuida-tuple-space
```

## TODO

### Tuple Spaces
- [x] Como rodar jgroups no docker?
- [ ] Eleição do Líder
- [ ] Envio de mensagens seguindo a ordem FIFO
- [ ] Recuperação de erros em caso de falha do líder (Outros casos não são importantes)
- [ ] Adicionar tupla
- [ ] Obter tupla
- [ ] Obter tupla async?
- [ ] Ler tupla sem remover

### Client
- [ ] Definir o que vai ser
- [ ] Comunicação com o espaço de tuplas