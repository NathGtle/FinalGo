#  Consignes pour le projet Go à réaliser le 4e jour

## Le projet devra :

* être compilable avec `go build` avec Go version >= 1.16
* commenté

## Fonctionnalités générales

* Mettre en place un routeur HTTP avec `github.com/gin-gonic/gin`

* Servir le front situé dans `../public` avec `github.com/gin-contrib/static`

* Ajouter une route POST `/login` décodant le corps de la requête
  au format JSON pour obtenir un couple identifiant/mot de passe,  
  ex:
  `{ "user_id" : "superfrog", "password" : "gopherade" }`

  Et retourner en JSON un token alphanumérique (idéalement) aléatoire, 
  ex : 
  `{ "token" : "09nioczicnzocxh" }`

  Note : ce token devrait être ensuite passé par les clients lors :
  - en QueryString de l'URL de connexion WebSocket
  - de leur requêtes XHR via un request header nommé 'Authorization'

* Ajouter une route GET `/agents` retournant en format JSON un tableau  
  des positions de tous les agents, ex :
  '[ { "id : "", "pos_x" : "23", "pos_y" : "300" } ]'  
  Note : ceci ne pourra être entièrement réalisable qu'ultérieurement  
  quand aura mis en place la gestion des connexions websocket.

Si possible, pour l'ensemble des routes REST la présence du token  
d'authentification devra être vérifiée via une fonction middleware.


## Mise en place de la gestion des WebSockets avec Melody

* Importer la lib [Melody](https://github.com/olahol/melody) v1 depuis le site `gopkg.in`

Dans `main()` :

* Créer une instance de `melody.Melody` avec `melody.New()`

* Ajouter 2 routes `GET` pour les connexions des rôle "monitor" et "agent"  
  et y placer un appel à :
	
  ```
  melodyRouter.HandleRequest(c.Writer, c.Request)
  ```

  Les chemins de ces routes sont ceux utilisés par le frontend dans  
  les fichiers `index.html` et `agent.html`

* Sur l'instance `*melody.Melody` ajouter un appel à `Melody.HandleConnect`  
  - Remarquer que l'on peut distinguer les connexions issues du moniteur  
  et des agents en inspectant `melody.Melody.Request.URL.Path`.

  En tirant parti de l'information précédente :

  - mémoriser la session du moniteur dans une variable utile plus tard.

  - envoyer au moniteur la position initiale de l'agent, et un ID de  
  session en format JSON, par ex: 
  `{ "ID" : "0x2347878", "x" : posX, "y" : posY }`

  TOFIX
  Recommandation : cet ID pourrait être l'adresse de la melody.Session  
  transformée en représentation textuelle.

  https://github.com/olahol/melody/blob/7bd65910e5ab/session.go#L189

* Sur l'instance `*melody.Melody` ajouter un appel à `*Melody.HandleMessage`  
  avec le [paramètre attendu](https://pkg.go.dev/github.com/olahol/melody#Melody.HandleMessage).

* Dans la fonction passée à `Melody.HandleMessage`, déclenchée par les  
  envois de mises à jour de position par les clients, il faut retransmettre  
  celles-ci au moniteur, avec en plus, l'ID de l'agent.

## Suivi des sessions WebSockets

TODO fonctions add/remove/get
TODO test unitaire


## Suggestions d'améliorations

Ces consignes demandaient d'assigner aux agents un ID égal à l'adresse  
de leur session *Melody.

Pour des raisons de conception et de sécurité ce n'est pas souhaitable.
On pourrait mettre en place une autre map permettant de retrouver une  
session Agent à partir de son ID.