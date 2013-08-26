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
**/subject/ GET** : Check the status for a list of users


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

**/subject/{subject_id} DELETE** : This method closes the connection with the system and removes all the information of the subject, should to be called when the subject closes the session. If the system don't detect activity for the subject it will execute this method automatically

Parameters:
 * security_token: The security token returned by the authentication request

Response:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/ GET** : Retuns a list with all the aviable groups

Parameters:
 * security_token: The security token returned by the authentication request

Response:
 * list of dictionaries:
  * id : The group identifier
  * name : The group name
  * security_token: The security token returned by the authentication request

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

**/groups/{group_id}/ PUT** : Updates a group configuration

Parametes:
 * name: The group name
 * max_subjects (optional): The nax number of subjects that can observe this group, if this param is not specified, the number of subjects will be unlimited
 * security_token: The security token returned by the authentication request

Returns:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/{group_id} DELETE** : Removes an existing group with the given group identifier

Parameters:
 * security_token: The security token returned by the authentication request

Returns:
 * success : [ok|ko]
 * reason : In case of success ko, returns the reason as a string

**/groups/{group_id}/ GET** : Retuns a list with all the subjects suscribed to the given group

Returns:
 * subjects : The list of all the identifiers for the suscribed subjects to the given group

**/msgs/ POST** : Send a new message to a list of users or groups, the messages can include a special validation and parametersbo

Parameters:
 * users : List of user IDs that will receive the messages
 * groups : List of groups wich users will receive the messages
 * message : The text of the message to be sent
 * type : The type of the message, for special messages which requires validation or any backend action, specify the type here according to the pre configured types
 * args : JSON with all the arguments to be used on the validation of the message type is necessary, this parrameters will be always sent to target subjects
 * security_token: The security token returned by the authentication request

**/msgs/ GET** : Returns the list of messages for the current user (the system takes the user from the security_token). This request is not answered until one or more messages for the subject are available, or the max time defined in the config is expired, the client should to wait until the server responses and send a new request inmediatly after the server responses. See: [BOSH protocol](http://en.wikipedia.org/wiki/BOSH "BOSH - Wikipedia")

Parameters:
 * security_token: The security token returned by the authentication request

Returns:
 * List of messages for the current subject, with a dictionary with the next params for each one:
  * txt : The text of the message as was sent by the sender
  * type : The message type
  * args : The object with all the arguments

