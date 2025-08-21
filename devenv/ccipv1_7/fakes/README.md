# CCIPv2 Fakes

This server contains CCIPv2 specific fakes

To run it locally and develop (you need a fine-grained GH approved token here)
```
just build $tag
just run
```

Test it
```
curl -v "http://localhost:9111/static-fake"
curl -v "http://localhost:9111/dynamic-fake"
```
Publish it
```
just publish $tag
```