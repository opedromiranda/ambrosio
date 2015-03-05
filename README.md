# Ambrosio

A bot written in Go. Inspired by Github's [Hubot](https://hubot.github.com) and created to get to know Go better.

### Version
0.0.1

### Usage

Add the behaviour to Ambrosio:

```golang
import (
    "github.com/opedromiranda/ambrosio"
)

func main() {
	steve := ambrosio.NewAmbrosio("Steve") // give a name to the bot

    // add behaviours

	steve.Listen(3000)
}

```

#### Interaction
```
localhost:3000/ask?action="<string>"

```
License
----

MIT
