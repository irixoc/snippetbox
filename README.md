# Snippet Box App

A web app written in go from Alex Edward's Amazing Book "Let's Go"


What happens when you register routes without creating/declaring a servemux
Behind the scenes, these functions register their routes with something called the DefaultServeMux. There’s nothing special about this — it’s just regular servemux like we’ve already been using, but which is initialized by default and stored in a net/http global variable. Here’s the relevant line from the Go source code:

Additional Information
Servemux Features and Quirks
In Go’s servemux, longer URL patterns always take precedence over shorter ones. 
- So, if a servemux contains multiple patterns which match a request, it will always dispatch the request to the handler corresponding to the longest pattern. This has the nice side-effect that you can register patterns in any order and it won’t change how the servemux behaves.

Request URL paths are automatically sanitized. 
- If the request path contains any . or .. elements or repeated slashes, the user will automatically be redirected to an equivalent clean URL. For example, if a user makes a request to /foo/bar/..//baz they will automatically be sent a 301 Permanent Redirect to /foo/baz instead.

If a subtree path has been registered and a request is received for that subtree path without a trailing slash, then the user will automatically be sent a 301 Permanent Redirect to the subtree path with the slash added. 
- For example, if you have registered the subtree path /foo/, then any request to /foo will be redirected to /foo/.


# Restful Routing in GO
It’s important to acknowledge that the routing functionality provided by Go’s servemux is pretty lightweight. 
It doesn’t support:
- routing based on the request method, 
- it doesn’t support semantic URLs with variables in them, and 
- it doesn’t support regexp-based patterns.