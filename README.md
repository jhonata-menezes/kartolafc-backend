# Documentação: KartolaFC API
### O sistema do KartolaFc encontra-se disponivel em: https://api.kartolafc.com.br

### Todas rotas disponiveis podem ser visualizadas pelo arquivo [routes.go](https://github.com/jhonata-menezes/kartolafc-backend/blob/master/routes.go)

Esta é uma breve documentação sobre os endpoints disponiveis


##### URL Base: https://api.kartolafc.com.br

* Status da rodada: GET https://api.kartolafc.com.br/mercado/status

* Pesquisa por nome do time ou nome do cartoleiro: GET /times/{nome} ex.: https://api.kartolafc.com.br/times/jhonata

* Consulta a escalação do time por id: GET /time/id/{id} ex.: https://api.kartolafc.com.br/time/id/14254256

* Consulta o time em rodada passada: GET /time/id/{id}/{rodada} ex.: https://api.kartolafc.com.br/time/id/14254256/2

* Consulta dos atletas disponiveis para escalação: GET https://api.kartolafc.com.br/atletas/mercado

* Destaques para a rodada (mais escalados): GET https://api.kartolafc.com.br/mercado/destaques

* Pesquisa de ligas: GET /ligas/{nome_da_Liga} ex.: https://api.kartolafc.com.br/ligas/cartola

* Detalhes especificos de uma liga: GET /liga/{slug}/{page} ex.: https://api.kartolafc.com.br/liga/cartolamizade/1

* Parciais do cartolafc: GET https://api.kartolafc.com.br/atletas/pontuados

* Jogos da rodada, caso não informe a partida é retornado informações da rodada atual: GET /partidas/{partida} ex.: https://api.kartolafc.com.br/partidas/2

* Histórico de pontuação de um atleta especifico: GET /atletas/historico/{id} ex.: https://api.kartolafc.com.br/atletas/historico/37281

* Realiza o login no portal da globo e retorna o token, necessario passar no corpo da requisição
```javascript
 {"email": "", "senha": ""}
 ```
POST /login/cartolafc

Exemplo de resposta:
```javascript
{
    "id":"Authenticated",
    "userMessage":"Usuário autenticado com sucesso",
    "glbId":"1c31f51ae0c28fab34390d51d4eecf904326e7a7a354761364e616b306472613335564c70523430617930716d35306938784e42664b643076716c504858556a4c504a716"
}

```

* Informações sobre o time, necessário passar o token de autenticação: GET /time/info
Header necessário: token: {valor}