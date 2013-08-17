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
 * id (optional): The subject identifier which is guaranteed to be unique, if this parameter is not specified, the system will generate a GUID for the user and return it
 * name: The nick name of the user
 * auth_token (optional): This parameter if mandatory when the authentication system is enabled, and is the token to be sent to the authentication server

Response:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string
 * id: If the id was not specified will return the auto generated identifier
 * security_token: A token that should to be included in all the next queries to the system

**/subject/ DELETE** : This method closes the connection with the system and removes all the information of the subject, should to be called when the subject closes the session. If the system don't detect activity for the subject it will execute this method automatically

Parameters:
 * id: The subject identifier 
 * security_token: The security token returned by the authentication request

Response:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/ GET** : Retuns a list with all the aviable groups

Parameters:
 * security_token: The security token returned by the authentication request

Response:
 * list of dictionaries:
 ** id : The group identifier
 ** name : The group name

**/groups/ POST** : Creates a new group and returns the group identifier

Parameters:
 * id (optional): The group identifier which is guaranteed to be unique, if this parameter is not specified, the system will generate a GUID for the user and return it
 * name: The group name
 * max_subjects (optional): The nax number of subjects that can observe this group, if this param is not specified, the number of subjects will be unlimited
 * security_token: The security token returned by the authentication request

Returns:
 * id : The group identifier
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/ PUT** : Updates a group configuration

Parametes:
 * id: The group unique identifier
 * name: The group name
 * max_subjects (optional): The nax number of subjects that can observe this group, if this param is not specified, the number of subjects will be unlimited
 * security_token: The security token returned by the authentication request

Returns:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/ DELETE** : Removes an existing group

Parametes:
 * id: The unique identifier of the group to be removed

Returns:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/msg/ POST**
**/msg/ GET**
