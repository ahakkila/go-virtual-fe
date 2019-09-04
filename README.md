# Virtual host frontend in Go language
An example virtual host frontend in Go language proxying requests to defined 
backend servers using httputil.SingleHostReverseProxy

## Purpose
I couldn't find anything simple implementing this so I did it just for fun. 

For production use there should be better error handling and if you stumbled 
here looking for out-of-the-box solution, you'd probably be better off with
[GorillaMux](http://www.gorillatoolkit.org/pkg/mux) or 
[Vulcand](https://github.com/vulcand) or similar project. I couldn't figure out 
how to get various hosts with port numbers to route properly with Gorilla 
Toolkit Mux, though.

## Credits
This owes a lot to Stackoverflow question
http://stackoverflow.com/questions/14170799/how-to-get-virtualhost-functionality-in-go

and linked example
https://gist.github.com/camoles/523dac8cc0fe40d52f66

