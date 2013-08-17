GoPost
=======
GoPost is a real time communication system designed to send messages between users and groups. Implements an API REST in order to be used directly from the client side using the [BOSH protocol](http://en.wikipedia.org/wiki/BOSH "BOSH - Wikipedia").

All the system can run as autonomous or together with another system that can be used to authenticate users, validate actions (messages who needs a backend validation), etc.

Some cases of use:
 * Real time video games that need a real time communication system between users
 * Collaborative text editing suites
 * Chat systems
 * News systems

GoPost is designed to scale horizontally, and be as fast and light as possible.

REST API
--------------
**/subject/ POST** : This method register a new subject on the system, and authenticate it if necessary, should to be called in order to establish a session on the system

Parameters:
 * id (optional): The subject identifier which is guaranteed to be unique, if this parameter is not specified, the system will generate a UUID for the user and return it
 * name: The nick name of the user
 * auth_token (optional): This parameter if mandatory when the authentication system is enabled, and is the token to be sent to the authentication server

Response:
 * id: If the id was not specified will return the auto generated identifier
 * security_token: A token that should to be included in all the next queries to the system

**/subject/ DELETE** : This method closes the connection with the system and removes all the information of the subject, should to be called when the subject closes the session. If the system don't detect activity for the subject it will execute this method automatically

Parameters:
 * id: The subject identifier 
 * security_token: The security token returned by the authentication request

**/groups/ GET** : Retuns a list with all the aviable groups,
**/groups/ POST**
**/groups/ PUT**
**/groups/ DELETE**

**/msg/ POST**
**/msg/ GET**
