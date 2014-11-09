##  Show me the code

Let's look at some code. 5 key things I want to point in the code:

- The main func is the app entry point, There is no web server to tie into. This app is completely stand alone, no unicorn/uwsgi or apache required
- Unit tests live in the same package as the code but aren't compiled into the final exe
- Error handling is explicit, no try catch block to swallow exceptions
- Go allows and uses pointers but handles the dangerous and hard parts. Like pointer arthrmatic, and memory managment are taken care of.
- The use of channels and go routines to handle async processes, including the sync.WaitGroup to provide go routine managment.
